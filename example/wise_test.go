package example_test

import (
	"database/sql"
	"testing"

	"github.com/hanpama/pgmg/example/wise/tables/professor"
)

func TestInsertWithDefault(t *testing.T) {
	withTx(func(tx *sql.Tx) {
		q := professor.Insert(professor.InputValues(
			professor.FamilyName("Jin"),
			professor.GivenName("Ping"),
		))
		_, err := tx.Exec(q.SQL(), q.Args()...)
		if err != nil {
			panic(err)
		}

		q = professor.InsertReturning(professor.InputValues(
			professor.FamilyName("Ming"),
			professor.GivenName("Weiwei"),
		))
		var record professor.Record
		err = tx.QueryRow(q.SQL(), q.Args()...).Scan(record.Receive()...)
		if err != nil {
			panic(err)
		}
		assertDeepEqual(t, record.FamilyName, "Ming")
		assertDeepEqual(t, record.GivenName, "Weiwei")

		var count int
		err = tx.QueryRow(`SELECT count(*) FROM wise.professor`).Scan(&count)
		assertDeepEqual(t, count, 2)
	})
}
