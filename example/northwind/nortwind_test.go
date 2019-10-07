package northwind_test

import (
	"database/sql"
	"testing"

	"github.com/hanpama/pgmg/example/northwind/public/customers"
	"github.com/hanpama/pgmg/example/northwind/public/employee_territories"
	"github.com/hanpama/pgmg/example/northwind/public/us_states"
	"github.com/hanpama/pgmg/example/northwind/queries/domestic_order"
)

func TestModelScanning(t *testing.T) {
	withTx(func(tx *sql.Tx) {
		var recordset customers.Recordset
		rows, err := tx.Query(`SELECT * FROM customers LIMIT 3`)
		defer rows.Close()
		if err != nil {
			panic(err)
		}
		for rows.Next() {
			if err = rows.Scan(recordset.ReceiveNext()...); err != nil {
				t.Fatal(err)
			}
		}
		testJSONSnapshot(t, "TestModelScanning", recordset)
	})
}

func TestSelectBySimpleKey(t *testing.T) {
	withTx(func(tx *sql.Tx) {
		var record customers.Record
		query := customers.Select(customers.PkCustomers{CustomerID: "ANTON"})

		err := tx.QueryRow(query.SQL(), query.Args()...).Scan(record.Receive()...)
		if err != nil {
			t.Fatal(err)
		}
		testJSONSnapshot(t, "TestSelectBySimpleKey", record)
	})
}

func TestSelectByCompositeKey(t *testing.T) {
	withTx(func(tx *sql.Tx) {
		var record employee_territories.Record

		query := employee_territories.Select(employee_territories.PkEmployeeTerritories{
			EmployeeID:  3,
			TerritoryID: "31406",
		})

		err := tx.QueryRow(query.SQL(), query.Args()...).Scan(record.Receive()...)
		if err != nil {
			t.Fatal(err)
		}
		testJSONSnapshot(t, "TestSelectByCompositeKey", record)
	})
}

func TestInsert(t *testing.T) {
	withTx(func(tx *sql.Tx) {
		query := customers.Insert(
			customers.InputValues(
				customers.CustomerID("BC1"),
				customers.CompanyName("Big Company"),
				customers.ContactName("Big Customer 1"),
				customers.ContactTitle("CEO"),
				customers.Address("614-12, Some Place"),
				customers.City("City"),
				customers.Region("Some Region"),
				customers.PostalCode("512-213"),
				customers.Country("Country"),
				customers.Phone("+0041234412"),
				customers.Fax("12312123"),
			),
			customers.InputValues(
				customers.CustomerID("BC2"),
				customers.CompanyName("Big Company"),
				customers.ContactName("Big Customer 2"),
				customers.ContactTitle("CFO"),
				customers.Address("614-12, Some Place"),
				customers.City("City"),
				customers.Region("Some Region"),
				customers.PostalCode("512-213"),
				customers.Country("Country"),
				customers.Phone("+0041234412"),
				customers.Fax("12312123"),
			),
		)
		_, err := tx.Exec(query.SQL(), query.Args()...)
		if err != nil {
			panic(err)
		}

		var recordset customers.Recordset
		sel := customers.Select(customers.PkCustomers{CustomerID: "BC1"})
		err = tx.QueryRow(sel.SQL(), sel.Args()...).Scan(recordset.ReceiveNext()...)
		if err != nil {
			panic(err)
		}

		sel = customers.Select(customers.PkCustomers{CustomerID: "BC2"})
		err = tx.QueryRow(sel.SQL(), sel.Args()...).Scan(recordset.ReceiveNext()...)
		if err != nil {
			panic(err)
		}

		testJSONSnapshot(t, "TestInsert", recordset)
	})
}

func TestUpdateBySimpleKey(t *testing.T) {
	withTx(func(tx *sql.Tx) {
		key := customers.PkCustomers{CustomerID: "ANTON"}
		var record customers.Record
		query := customers.Update(key,
			customers.ContactName("Test"),
			customers.Fax("5432"),
		)
		_, err := tx.Exec(query.SQL(), query.Args()...)
		if err != nil {
			t.Fatal(err)
		}

		query = customers.Select(key)
		err = tx.QueryRow(query.SQL(), query.Args()...).Scan(record.Receive()...)
		if err != nil {
			panic(err)
		}
		assertDeepEqual(t, *record.ContactName, "Test")
		assertDeepEqual(t, *record.Fax, "5432")
		testJSONSnapshot(t, "TestUpdateBySimpleKey", record)
	})
}

func TestDeleteBySimpleKey(t *testing.T) {
	withTx(func(tx *sql.Tx) {
		key := us_states.PkUsstates{StateID: 21} // Maryland
		query := us_states.Delete(key)
		_, err := tx.Exec(query.SQL(), query.Args()...)
		if err != nil {
			t.Fatal(err)
		}
		var count int
		err = tx.QueryRow(`SELECT count(*) FROM us_states WHERE state_id = 21`).Scan(&count)
		if err != nil {
			t.Fatal(err)
		}
		assertDeepEqual(t, count, 0)
	})
}

func TestQuery(t *testing.T) {
	withTx(func(tx *sql.Tx) {
		country := "Italy"
		var num int64 = 10
		q := domestic_order.Query(&country, &num)
		var recordset domestic_order.Recordset

		rows, err := tx.Query(q.SQL(), q.Args()...)
		if err != nil {
			t.Fatal(err)
		}
		for rows.Next() {
			err = rows.Scan(recordset.ReceiveNext()...)
			if err != nil {
				t.Fatal(err)
			}
		}
		testJSONSnapshot(t, "TestQuery", recordset)
	})
}
