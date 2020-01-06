package example_test

import (
	"context"
	"database/sql"
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

	source := []*example.SemesterRow{
		example.NewSemesterRow(
			example.SemesterRowID(nil),
			example.SemesterRowYear(2018),
			example.SemesterRowSeason("spring"),
		),
		example.NewSemesterRow(
			example.SemesterRowID(nil),
			example.SemesterRowYear(2018),
			example.SemesterRowSeason("fall"),
		),
		example.NewSemesterRow(
			example.SemesterRowID(nil),
			example.SemesterRowYear(2019),
			example.SemesterRowSeason("spring"),
		),
		example.NewSemesterRow(
			example.SemesterRowID(nil),
			example.SemesterRowYear(2019),
			example.SemesterRowSeason("fall"),
		),
	}
	if saved, err := example.SaveAndReturnBySemesterPkey(ctx, tdb, source...); err != nil {
		t.Fatal(err)
	} else {
		for i, expected := range []*example.SemesterRow{
			&example.SemesterRow{ID: source[0].ID, Season: "spring", Year: 2018},
			&example.SemesterRow{ID: source[1].ID, Season: "fall", Year: 2018},
			&example.SemesterRow{ID: source[2].ID, Season: "spring", Year: 2019},
			&example.SemesterRow{ID: source[3].ID, Season: "fall", Year: 2019},
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
		for i, expected := range []*example.SemesterRow{
			&example.SemesterRow{ID: source[0].ID, Season: "spring", Year: 2018},
			&example.SemesterRow{ID: source[1].ID, Season: "fall", Year: 2018},
			nil, // should be nil
			&example.SemesterRow{ID: source[1].ID, Season: "fall", Year: 2018},
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
		for i, expected := range []*example.SemesterRow{nil, nil, source[2], source[3]} {
			if expected == nil && gots[i] != nil {
				t.Fatalf("Expected to be nil on index %d but got %+v", i, gots[i])
			}
			if expected != nil && !reflect.DeepEqual(gots[i], expected) {
				t.Fatalf("Expected to get proper values for index %d but got %+v", i, gots[i])
			}
		}
	}
}

func TestSaveReturningError(t *testing.T) {
	ctx := context.Background()
	tdb, err := newTestDB(ctx, "user=postgres dbname=pgmg sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	_, err = example.SaveAndReturnByProductPkey(ctx, tdb,
		&example.ProductRow{Price: 1},
		&example.ProductRow{},
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
