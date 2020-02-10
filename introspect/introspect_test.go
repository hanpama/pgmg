package introspect_test

import (
	"database/sql"
	"encoding/json"
	"testing"

	"github.com/hanpama/pgmg/introspect"
	"github.com/hanpama/pgmg/testutil"

	_ "github.com/lib/pq"
)

var inspFp = "../example/introspection.json"

func TestPGMG(t *testing.T) {
	db, err := sql.Open(
		"postgres",
		"user=postgres dbname=pgmg sslmode=disable",
	)
	if err != nil {
		t.Fatal(err)
	}
	tables, err := introspect.IntrospectSchema(db, "wise")
	if err != nil {
		t.Fatal(err)
	}

	gotJSON, err := json.MarshalIndent(tables, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	if err = testutil.TestSnapshot(gotJSON, inspFp); err != nil {
		t.Fatal(err)
	}
}
