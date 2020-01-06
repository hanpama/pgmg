package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
)

const help = `usage: pgmg -database <connection_string> -schema <schema_name> -out <outfile>

See https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING
for more information about connection string parameters.

Tables generation:
	go run github.com/hanpama/pgmg \
		-database 'user=postgres dbname=pgmg sslmode=disable' \
		-schema public \
		-out public/schema.go
`

func main() {
	if len(os.Args) == 0 {
		fmt.Fprintf(os.Stdin, help)
		return
	}
	database := flag.String("database", "", "connection string")
	schema := flag.String("schema", "public", "target schema")
	out := flag.String("out", "pgmg_gen.go", "output file path")
	flag.Parse()

	err := run(*database, *schema, *out)
	if err != nil {
		log.Fatal(err)
	}
}

func run(dbURL string, schema string, outFile string) (err error) {
	outFile, err = filepath.Abs(outFile)
	if err != nil {
		return err
	}

	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		return err
	}
	tables, err := IntrospectSchema(db, schema)
	if err != nil {
		return err
	}
	println(outFile)

	outDir := filepath.Dir(outFile)

	if err = ensureMkdir(outDir); err != nil {
		return err
	}

	b, err := RenderTableModel(filepath.Base(outDir), tables)
	if err != nil {
		println(string(b))
		return err
	}
	err = ioutil.WriteFile(outFile, b, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func ensureMkdir(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return os.Mkdir(path, os.ModePerm)
	} else if err != nil {
		return err
	}
	return nil
}
