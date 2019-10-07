/// 데이터베이스를 introspect 해서 테이블 스키마를 Go 구조체로 된 모델로 만듭니다.

package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/hanpama/pgmg/internal"
	_ "github.com/lib/pq"
)

const help = `usage: pgmg table <connection_string> <schema_name> <outdir>
	OR pgmg query <connection_string> <sql_glob>

See https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING
for more information about connection string parameters.

example: pgmg 'user=postgres dbname=pgmg sslmode=disable' public example/northwind/public
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

	} else if mode == "query" {
		if len(os.Args) != 4 {
			fmt.Fprintf(os.Stdin, help)
			fmt.Fprint(os.Stderr, "Invalid number of arguments")
			return
		}
		dbURL := os.Args[2]
		glob := os.Args[3]
		runQueryMode(dbURL, glob)

	} else {
		fmt.Fprint(os.Stderr, "Mode should be 'table' or 'query'")
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
		tableDir := path.Join(outDir, table.Name)
		ensureMkdir(tableDir)
		tableModelPath := path.Join(tableDir, "model.go")

		b, err := internal.RenderTableModel(&table)
		if err != nil {
			println(string(b))
			panic(err)
		}
		err = ioutil.WriteFile(tableModelPath, b, os.ModePerm)
		if err != nil {
			panic(err)
		}

		queryPath := path.Join(tableDir, "query.go")

		b, err = internal.RenderTableQuery(&table)
		if err != nil {
			println(string(b))
			panic(err)
		}
		err = ioutil.WriteFile(queryPath, b, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

func runQueryMode(dbURL, glob string) {
	tx := createTx(dbURL)
	defer tx.Rollback()

	files, err := filepath.Glob(glob)
	if err != nil {
		panic(err)
	}

	for _, fp := range files {
		filename := filepath.Base(fp)
		println(fp)

		name := strings.Split(filename, ".")[0]

		b, err := ioutil.ReadFile(fp)
		if err != nil {
			panic(err)
		}
		query := string(b)
		qi, err := internal.IntrospectQuery(tx, fp, query)
		if err != nil {
			panic(err)
		}

		b, err = internal.RenderQuery(&qi)

		dir := filepath.Join(filepath.Dir(fp), name)

		ensureMkdir(dir)

		err = ioutil.WriteFile(filepath.Join(dir, name+".go"), b, os.ModePerm)
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
