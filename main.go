/// 데이터베이스를 introspect 해서 테이블 스키마를 Go 구조체로 된 모델로 만듭니다.

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

const help = `usage: pgmg <connection_string> <schema_name> <outdir>

See https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING
for more information about connection string parameters.

example: pgmg 'user=postgres dbname=pgmg sslmode=disable' public example/northwind/public
`

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stdin, help)
		fmt.Fprint(os.Stderr, "Invalid number of arguments")
		return
	}
	dbURL := os.Args[1]
	schema := os.Args[2]
	outDir := os.Args[3]

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()
	tables, err := internal.Introspect(tx, schema)
	if err != nil {
		panic(err)
	}

	ensureMkdir(outDir)
	for _, table := range tables {
		println(table.Name)
		tableDir := path.Join(outDir, table.Name)
		ensureMkdir(tableDir)
		tableModelPath := path.Join(tableDir, "model.go")

		b, err := internal.RenderModel(&table)
		if err != nil {
			println(string(b))
			panic(err)
		}
		err = ioutil.WriteFile(tableModelPath, b, os.ModePerm)
		if err != nil {
			panic(err)
		}

		queryPath := path.Join(tableDir, "query.go")

		b, err = internal.RenderQuery(&table)
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

func ensureMkdir(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}
