package example_test

import (
	"context"
	"database/sql"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/hanpama/pgmg/example"
	_ "github.com/lib/pq"
)

func TestExample(t *testing.T) {
	ctx := context.Background()
	tdb, err := newTestDB(ctx, "user=postgres dbname=pgmg sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	// DROP existing schema if exists
	var dbmig []byte
	if dbmig, err = ioutil.ReadFile("down.sql"); err != nil {
		t.Fatal(err)
	}
	if _, err = tdb.Exec(ctx, string(dbmig)); err != nil {
		t.Fatal(err)
	}
	// Prepare empty tables
	if dbmig, err = ioutil.ReadFile("up.sql"); err != nil {
		t.Fatal(err)
	}
	if _, err = tdb.Exec(ctx, string(dbmig)); err != nil {
		t.Fatal(err)
	}

	source := []*example.Semester{
		example.NewSemester(
			example.SemesterID(nil),
			example.SemesterYear(2018),
			example.SemesterSeason("spring"),
		),
		example.NewSemester(
			example.SemesterID(nil),
			example.SemesterYear(2018),
			example.SemesterSeason("fall"),
		),
		example.NewSemester(
			example.SemesterID(nil),
			example.SemesterYear(2019),
			example.SemesterSeason("spring"),
		),
		example.NewSemester(
			example.SemesterID(nil),
			example.SemesterYear(2019),
			example.SemesterSeason("fall"),
		),
	}
	if saved, err := example.SaveAndReturnBySemesterPkey(ctx, tdb, source...); err != nil {
		t.Fatal(err)
	} else {
		for i, expected := range []*example.Semester{
			&example.Semester{ID: source[0].ID, Season: "spring", Year: 2018},
			&example.Semester{ID: source[1].ID, Season: "fall", Year: 2018},
			&example.Semester{ID: source[2].ID, Season: "spring", Year: 2019},
			&example.Semester{ID: source[3].ID, Season: "fall", Year: 2019},
		} {
			if *saved[i] != *expected {
				t.Fatalf("Expected to equal on index %d, but got %+v", i, saved[i])
			}
			if source[i] != saved[i] {
				t.Fatalf("Expected to have same reference but it changed")
			}
		}
	}

	if gots, err := example.GetBySemesterPkey(ctx, tdb,
		example.SemesterPkey{ID: *source[0].ID},
		example.SemesterPkey{ID: *source[1].ID},
		example.SemesterPkey{ID: 0}, // which does not exist
		example.SemesterPkey{ID: *source[1].ID},
	); err != nil {
		t.Fatal(err)
	} else {
		for i, expected := range []*example.Semester{
			&example.Semester{ID: source[0].ID, Season: "spring", Year: 2018},
			&example.Semester{ID: source[1].ID, Season: "fall", Year: 2018},
			nil, // should be nil
			&example.Semester{ID: source[1].ID, Season: "fall", Year: 2018},
		} {
			if expected == nil && gots[i] != nil {
				t.Fatalf("Expected to be nil on index %d but got %+v", i, gots[i])
			}
			if expected != nil && !reflect.DeepEqual(gots[i], expected) {
				t.Fatalf("Expected to get proper values for index %d but got %+v", i, gots[i])
			}
		}
	}

	// delete ...
	if deleted, err := example.DeleteBySemesterPkey(ctx, tdb,
		example.SemesterPkey{ID: *source[0].ID},
		example.SemesterPkey{ID: *source[1].ID},
		example.SemesterPkey{ID: *source[1].ID}, // dup
		example.SemesterPkey{ID: 0},             // which does not exist
	); err != nil {
		t.Fatal(err)
	} else {
		if deleted != 2 {
			t.Fatalf("Expected to delete 2 rows but got %d", deleted)
		}
	}
	// and update
	source[2].Season = "summer"
	if err = example.SaveBySemesterPkey(ctx, tdb, source[2]); err != nil {
		t.Fatal(err)
	}

	if gots, err := example.GetBySemesterPkey(ctx, tdb,
		example.SemesterPkey{ID: *source[0].ID},
		example.SemesterPkey{ID: *source[1].ID},
		example.SemesterPkey{ID: *source[2].ID},
		example.SemesterPkey{ID: *source[3].ID},
	); err != nil {
		t.Fatal(err)
	} else {
		for i, expected := range []*example.Semester{nil, nil, source[2], source[3]} {
			if expected == nil && gots[i] != nil {
				t.Fatalf("Expected to be nil on index %d but got %+v", i, gots[i])
			}
			if expected != nil && !reflect.DeepEqual(gots[i], expected) {
				t.Fatalf("Expected to get proper values for index %d but got %+v", i, gots[i])
			}
		}
	}

	var count int
	if count, err = example.CountSemesterRows(ctx, tdb, example.SemesterCondition{}); err != nil {
		t.Fatal(err)
	} else if count != 2 {
		t.Fatalf("Unexpected row count %d", count)
	}

	if count, err = example.CountSemesterRows(ctx, tdb, example.SemesterCondition{
		Season: &source[3].Season,
	}); err != nil {
		t.Fatal(err)
	} else if count != 1 {
		t.Fatalf("Unexpected row count %d", count)
	}
}

func TestSaveReturningError(t *testing.T) {
	ctx := context.Background()
	tdb, err := newTestDB(ctx, "user=postgres dbname=pgmg sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	_, err = example.SaveAndReturnByProductPkey(ctx, tdb,
		&example.Product{Price: 1},
		&example.Product{},
	)
	if err == nil {
		t.Fatal("Expected to get error by constraint ")
	}

}

type testDB struct {
	b *sql.DB
}

var _ example.PGMGDatabase = (*testDB)(nil)

func (db *testDB) QueryScan(ctx context.Context, r func(int) []interface{}, sql string, args ...interface{}) (rowsReceived int, err error) {
	rows, err := db.b.QueryContext(ctx, sql, args...)
	if err != nil {
		return rowsReceived, err
	}
	for rows.Next() {
		err = rows.Scan(r(rowsReceived)...)
		if err != nil {
			return rowsReceived, err
		}
		rowsReceived++
	}
	return rowsReceived, err
}

func (db *testDB) Exec(ctx context.Context, sql string, args ...interface{}) (int64, error) {
	res, err := db.b.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func newTestDB(ctx context.Context, url string) (*testDB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &testDB{db}, nil
}
