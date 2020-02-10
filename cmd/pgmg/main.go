package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/hanpama/pgmg/model"

	"github.com/hanpama/pgmg/renderer"

	"github.com/hanpama/pgmg/introspect"
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

	outPath, err := filepath.Abs(*out)
	if err != nil {
		log.Fatal(err)
	}

	modelSchema, err := getModelSchema(*database, *schema)
	if err != nil {
		log.Fatal(err)
	}

	stat, err := os.Stat(outPath)
	if err == nil && stat.IsDir() {
		err = writeTableFiles(modelSchema, outPath)
	} else {
		err = writeSchemaFile(modelSchema, outPath)
	}

	if err != nil {
		log.Fatal(err)
	}
}

func getModelSchema(dbURL string, schemaName string) (*model.Schema, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}
	ins, err := introspect.IntrospectSchema(db, schemaName)
	if err != nil {
		return nil, err
	}
	return model.NewSchemaFromIntrospection(ins), nil
}

func writeTableFiles(schema *model.Schema, outDir string) (err error) {
	packageName := filepath.Base(outDir)
	outFile := filepath.Join(outDir, "pgmg_common.go")

	var b []byte
	if b, err = renderer.RenderCommon(packageName); err != nil {
		return err
	}
	if err = ioutil.WriteFile(outFile, b, os.ModePerm); err != nil {
		return err
	}
	for _, t := range schema.Tables {
		outFile = filepath.Join(outDir, t.SQLName+".go")
		if b, err = renderer.RenderTable(packageName, t); err != nil {
			return err
		}
		if err = ioutil.WriteFile(outFile, b, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func writeSchemaFile(schema *model.Schema, outFile string) (err error) {
	packageName := filepath.Base(filepath.Dir(outFile))
	var b []byte
	if b, err = renderer.RenderSchema(packageName, schema); err != nil {
		return err
	}
	if err = ioutil.WriteFile(outFile, b, os.ModePerm); err != nil {
		return err
	}
	return nil
}
