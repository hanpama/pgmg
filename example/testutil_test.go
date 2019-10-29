package example_test

import (
	"database/sql"
	"reflect"
	"testing"

	_ "github.com/lib/pq"
)

type testCases []struct {
	name      string
	test      func() string
	expecting string
}

func withTx(fn func(tx *sql.Tx)) {
	db, err := sql.Open("postgres", "user=postgres dbname=pgmg sslmode=disable")
	if err != nil {
		panic(err)
	}
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	defer tx.Rollback()
	fn(tx)
}

func assertDeepEqual(t *testing.T, val interface{}, expected interface{}) {
	if !reflect.DeepEqual(val, expected) {
		t.Fatalf("Expected value to equal to %+v but got %+v", expected, val)
	}
}
