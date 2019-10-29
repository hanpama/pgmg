package example_test

import (
	"database/sql"
	"encoding/json"
	"testing"

	"github.com/hanpama/pgmg/example/wise"
)

func TestE2E(t *testing.T) {
	withTx(func(tx *sql.Tx) {
		// values marshaling
		b, err := json.Marshal(wise.Professor.Input(
			wise.Professor.FamilyName.New("Ming"),
			wise.Professor.GivenName.New("Zhiwei"),
		))
		if err != nil {
			t.Error(err)
			return
		}

		// insert one
		_, err = tx.Exec(
			wise.Professor.InsertOneJSON("$1").Statement(),
			b,
		)
		if err != nil {
			t.Error(err)
			return
		}

		// insert one returning
		var row wise.ProfessorRow
		err = tx.QueryRow(
			wise.Professor.InsertOneJSON("$1").Returning().Statement(),
			b,
		).Scan(row.Receive()...)
		if err != nil {
			t.Error(err)
			return
		}

		insertedID := row.ID
		if insertedID == 0 {
			t.Error("Inserted id should not be zero")
		}

		row = wise.ProfessorRow{}
		prof := wise.Professor.As("prof")

		// select
		err = tx.QueryRow(
			prof.Select().Where(prof.ID.Eq("$1")).Statement(),
			insertedID,
		).Scan(row.Receive()...)
		if err != nil {
			t.Error(err)
			return
		}

		if row.ID != insertedID {
			t.Errorf("Retrieve id should be %v but got %v", insertedID, row.ID)
		}

		// update by changeset
		b, err = json.Marshal(wise.ProfessorValues{
			wise.Professor.FamilyName.New("Shu"),
		})
		if err != nil {
			t.Error(err)
			return
		}

		res, err := tx.Exec(
			prof.UpdateJSON("$1").Where(prof.ID.Eq("$2")).Statement(),
			b, insertedID,
		)
		if err != nil {
			t.Error(err)
			return
		}
		num, err := res.RowsAffected()
		if err != nil {
			t.Error(err)
			return
		}
		if num != 1 {
			t.Errorf("Number of rows afftected should be 1 but got %v", num)
		}

		// deleting
		res, err = tx.Exec(
			prof.DeleteWhere(prof.ID.Eq("$1")).Statement(),
			insertedID,
		)
		if err != nil {
			t.Error(err)
			return
		}
		num, err = res.RowsAffected()
		if err != nil {
			t.Error(err)
			return
		}
		if num != 1 {
			t.Errorf("Number of rows afftected should be 1 but got %v", num)
		}
	})
}
