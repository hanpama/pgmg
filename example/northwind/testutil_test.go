package northwind_test

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/kr/pretty"
	_ "github.com/lib/pq"
)

func testJSONSnapshot(t *testing.T, name string, res interface{}) {
	resB, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	path := filepath.Join("__snapshots__", name+".json")
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		ioutil.WriteFile(path, resB, os.ModePerm)
		t.Logf("Wrote snapshot: %s(%s)", name, path)
		return
	}
	snapshotB, err := ioutil.ReadFile(path)

	if string(resB) != string(snapshotB) {
		pretty.Log(string(resB), string(snapshotB))
		t.Fail()
	}
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
