package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/hanpama/pgmg/internal"
	_ "github.com/lib/pq"
)

const help = `usage: pgmg table <connection_string> <schema_name> <outdir>

See https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING
for more information about connection string parameters.

Tables generation:
  go run github.com/hanpama/pgmg table 'user=postgres dbname=pgmg sslmode=disable' public public
`

func main() {
	if len(os.Args) == 0 {
		fmt.Fprintf(os.Stdin, help)
		return
	}

	mode := os.Args[1]
	if mode == "table" {
		if len(os.Args) != 5 {
			fmt.Fprintf(os.Stdin, help)
			fmt.Fprint(os.Stderr, "Invalid number of arguments")
			return
		}
		dbURL := os.Args[2]
		schema := os.Args[3]
		outDir := os.Args[4]
		runTableMode(dbURL, schema, outDir)

	} else {
		fmt.Fprint(os.Stderr, "Mode should be 'table'")
		return
	}
}

func runTableMode(dbURL string, schema string, outDir string) {

	tx := createTx(dbURL)
	defer tx.Rollback()

	tables, err := internal.IntrospectSchema(tx, schema)
	if err != nil {
		panic(err)
	}

	ensureMkdir(outDir)

	for _, table := range tables {
		println(table.Name)
		tableModelPath := path.Join(outDir, table.Name+".go")

		b, err := internal.RenderTableModel(path.Base(outDir), &table)
		if err != nil {
			println(string(b))
			panic(err)
		}
		err = ioutil.WriteFile(tableModelPath, b, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

func createTx(dbURL string) *sql.Tx {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	return tx
}

func ensureMkdir(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}
