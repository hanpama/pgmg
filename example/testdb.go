package example

import (
	"context"
	"database/sql"
	"io/ioutil"
	"path/filepath"
	"testing"
)

type TestDB struct {
	b *sql.DB
}

func (db *TestDB) QueryAndReceive(ctx context.Context, r func(int) []interface{}, sql string, args ...interface{}) (rowsReceived int, err error) {
	rows, err := db.b.QueryContext(ctx, sql, args...)
	if err != nil {
		return rowsReceived, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(r(rowsReceived)...)
		if err != nil {
			return rowsReceived, err
		}
		rowsReceived++
	}
	return rowsReceived, err
}

func (db *TestDB) ExecAndCount(ctx context.Context, sql string, args ...interface{}) (int64, error) {
	res, err := db.b.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func NewTestDB(ctx context.Context, t *testing.T, migdir string) *TestDB {
	db, err := sql.Open("postgres", "user=postgres dbname=pgmg sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	tdb := &TestDB{db}

	// DROP existing schema if exists
	var dbmig []byte
	if dbmig, err = ioutil.ReadFile(filepath.Join(migdir, "down.sql")); err != nil {
		t.Fatal(err)
	}
	if _, err = tdb.ExecAndCount(ctx, string(dbmig)); err != nil {
		t.Fatal(err)
	}
	// Prepare empty tables
	if dbmig, err = ioutil.ReadFile(filepath.Join(migdir, "up.sql")); err != nil {
		t.Fatal(err)
	}
	if _, err = tdb.ExecAndCount(ctx, string(dbmig)); err != nil {
		t.Fatal(err)
	}
	return tdb
}
