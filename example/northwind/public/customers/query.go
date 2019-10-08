// Code generated by github.com/hanpama/pgmg. DO NOT EDIT.
package customers

func Insert(vss ...Values) Query {
	return Query{InsertSQL, []interface{}{string(mustMarshalJSON(vss))}}
}
func InsertReturning(vss ...Values) Query {
	return Query{InsertReturningSQL, []interface{}{string(mustMarshalJSON(vss))}}
}
func Select(k key) Query {
	return k.selectSQL()
}
func Update(k key, args ...attribute) Query {
	return k.updateSQL(args...)
}
func Delete(k key) Query {
	return k.deleteSQL()
}

const (
	InsertSQL = `
		INSERT INTO "public"."customers" (
			"customer_id",
			"company_name",
			"contact_name",
			"contact_title",
			"address",
			"city",
			"region",
			"postal_code",
			"country",
			"phone",
			"fax"
		)
		SELECT
			"customer_id",
			"company_name",
			"contact_name",
			"contact_title",
			"address",
			"city",
			"region",
			"postal_code",
			"country",
			"phone",
			"fax"
		FROM json_populate_recordset(null::"public"."customers", $1)`
	InsertReturningSQL = `
		INSERT INTO "public"."customers" (
			"customer_id",
			"company_name",
			"contact_name",
			"contact_title",
			"address",
			"city",
			"region",
			"postal_code",
			"country",
			"phone",
			"fax"
		)
		SELECT
			"customer_id",
			"company_name",
			"contact_name",
			"contact_title",
			"address",
			"city",
			"region",
			"postal_code",
			"country",
			"phone",
			"fax"
		FROM json_populate_recordset(null::"public"."customers", $1)
		RETURNING
			"customer_id",
			"company_name",
			"contact_name",
			"contact_title",
			"address",
			"city",
			"region",
			"postal_code",
			"country",
			"phone",
			"fax"`
)

func (k PkCustomers) selectSQL() Query {
	return Query{SelectPkCustomers, []interface{}{
		k.CustomerID,
	}}
}
func (k PkCustomers) updateSQL(args ...attribute) Query {
	return Query{UpdatePkCustomers, []interface{}{
		k.CustomerID,
		string(mustMarshalJSON(Values(args))),
	}}
}
func (k PkCustomers) deleteSQL() Query {
	return Query{DeletePkCustomers, []interface{}{
		k.CustomerID,
	}}
}

const (
	SelectPkCustomers = `
		SELECT * FROM "public"."customers" WHERE ("customer_id") = ($1) LIMIT 1
		`
	UpdatePkCustomers = `
		UPDATE "public"."customers" __ut__
		SET "customer_id" = COALESCE(__ch__."customer_id", __ut__."customer_id"),
			"company_name" = COALESCE(__ch__."company_name", __ut__."company_name"),
			"contact_name" = COALESCE(__ch__."contact_name", __ut__."contact_name"),
			"contact_title" = COALESCE(__ch__."contact_title", __ut__."contact_title"),
			"address" = COALESCE(__ch__."address", __ut__."address"),
			"city" = COALESCE(__ch__."city", __ut__."city"),
			"region" = COALESCE(__ch__."region", __ut__."region"),
			"postal_code" = COALESCE(__ch__."postal_code", __ut__."postal_code"),
			"country" = COALESCE(__ch__."country", __ut__."country"),
			"phone" = COALESCE(__ch__."phone", __ut__."phone"),
			"fax" = COALESCE(__ch__."fax", __ut__."fax")
		FROM (SELECT * FROM json_populate_record(null::"public"."customers", $2)) __ch__
		WHERE (__ut__."customer_id") = ($1)`
	DeletePkCustomers = `
		DELETE FROM "public"."customers"
		WHERE ("customer_id") = ($1)`
)

type key interface {
	selectSQL() Query
	updateSQL(args ...attribute) Query
	deleteSQL() Query
}

type Query struct {
	sql  string
	args []interface{}
}

func (q Query) SQL() string         { return q.sql }
func (q Query) Args() []interface{} { return q.args }
