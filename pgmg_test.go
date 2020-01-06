package main_test

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	pgmg "github.com/hanpama/pgmg"
)

var inspFp = "example/introspection.json"
var schemaFp = "example/schema.go"

func TestPGMG(t *testing.T) {
	db, err := sql.Open(
		"postgres",
		"user=postgres dbname=pgmg sslmode=disable",
	)
	if err != nil {
		t.Fatal(err)
	}
	tables, err := pgmg.IntrospectSchema(db, "wise")
	if err != nil {
		t.Fatal(err)
	}

	gotJSON, err := json.MarshalIndent(tables, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	if err = testSnapshot(gotJSON, inspFp); err != nil {
		t.Fatal(err)
	}

	gotGoFile, err := pgmg.RenderTableModel("example", tables)
	if err != nil {
		t.Fatal(err)
	}

	if err = testSnapshot(gotGoFile, schemaFp); err != nil {
		t.Fatal(err)
	}
}

// testSnapshot tests the bytes equals the data written on fp file.
func testSnapshot(got []byte, fp string) (err error) {
	if _, err = os.Stat(fp); os.IsNotExist(err) {
		if err = ioutil.WriteFile(fp, got, os.ModePerm); err != nil {
			return err
		}
	} else if err != nil {
		return err
	} else {
		expected, err := ioutil.ReadFile(fp)
		if err != nil {
			return err
		}
		if string(got) != string(expected) {
			if err = ioutil.WriteFile(fp+".expected", expected, os.ModePerm); err != nil {
				panic(err)
			}
			if err = ioutil.WriteFile(fp, got, os.ModePerm); err != nil {
				panic(err)
			}
			return fmt.Errorf("Unexpected inspection result: %s", string(got))
		}
	}
	return nil
}
