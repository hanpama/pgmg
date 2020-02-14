package example

import (
	"context"
	"database/sql"
	"io/ioutil"
	"path/filepath"
)

type TestDB struct {
	DB    *sql.DB
	stmts map[string]*sql.Stmt
}

func (db *TestDB) QueryAndReceive(ctx context.Context, r func(int) []interface{}, sqlstmt string, args ...interface{}) (rowsReceived int, err error) {
	var rows *sql.Rows
	if stmt, ok := db.stmts[sqlstmt]; ok {
		rows, err = stmt.QueryContext(ctx, args...)
	} else {
		rows, err = db.DB.QueryContext(ctx, sqlstmt, args...)
	}
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
	res, err := db.DB.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (db *TestDB) Prepare(ctx context.Context, sql string) (err error) {
	db.stmts[sql], err = db.DB.PrepareContext(ctx, sql)
	return err
}

func NewTestDB(ctx context.Context, migdir string) (*TestDB, error) {
	db, err := sql.Open("postgres", "user=postgres dbname=pgmg sslmode=disable")
	if err != nil {
		return nil, err
	}
	tdb := &TestDB{DB: db, stmts: make(map[string]*sql.Stmt)}

	// DROP existing schema if exists
	var dbmig []byte
	if dbmig, err = ioutil.ReadFile(filepath.Join(migdir, "down.sql")); err != nil {
		return nil, err
	}
	if _, err = tdb.ExecAndCount(ctx, string(dbmig)); err != nil {
		return nil, err
	}
	// Prepare empty tables
	if dbmig, err = ioutil.ReadFile(filepath.Join(migdir, "up.sql")); err != nil {
		return nil, err
	}
	if _, err = tdb.ExecAndCount(ctx, string(dbmig)); err != nil {
		return nil, err
	}
	return tdb, nil
}
