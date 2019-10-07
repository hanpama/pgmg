// Code generated by github.com/hanpama/pgmg. DO NOT EDIT.
package suppliers

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
		INSERT INTO "suppliers"
		SELECT * FROM json_populate_recordset(null::"suppliers", $1) `
	InsertReturningSQL = `
		INSERT INTO "suppliers"
		SELECT * FROM json_populate_recordset(null::"suppliers", $1)
		RETURNING *`
)

func (k PkSuppliers) selectSQL() Query {
	return Query{SelectPkSuppliers, []interface{}{
		k.SupplierID,
	}}
}
func (k PkSuppliers) updateSQL(args ...attribute) Query {
	return Query{UpdatePkSuppliers, []interface{}{
		k.SupplierID,
		string(mustMarshalJSON(Values(args))),
	}}
}
func (k PkSuppliers) deleteSQL() Query {
	return Query{DeletePkSuppliers, []interface{}{
		k.SupplierID,
	}}
}

const (
	SelectPkSuppliers = `
		SELECT * FROM "suppliers" WHERE ("suppliers"."supplier_id") = ($1) LIMIT 1
		`
	UpdatePkSuppliers = `
		UPDATE "suppliers"
		SET "supplier_id" = COALESCE(_ch."supplier_id", "suppliers"."supplier_id"),
			"company_name" = COALESCE(_ch."company_name", "suppliers"."company_name"),
			"contact_name" = COALESCE(_ch."contact_name", "suppliers"."contact_name"),
			"contact_title" = COALESCE(_ch."contact_title", "suppliers"."contact_title"),
			"address" = COALESCE(_ch."address", "suppliers"."address"),
			"city" = COALESCE(_ch."city", "suppliers"."city"),
			"region" = COALESCE(_ch."region", "suppliers"."region"),
			"postal_code" = COALESCE(_ch."postal_code", "suppliers"."postal_code"),
			"country" = COALESCE(_ch."country", "suppliers"."country"),
			"phone" = COALESCE(_ch."phone", "suppliers"."phone"),
			"fax" = COALESCE(_ch."fax", "suppliers"."fax"),
			"homepage" = COALESCE(_ch."homepage", "suppliers"."homepage")
		FROM (SELECT * FROM json_populate_record(null::"suppliers", $2)) _ch
		WHERE ("suppliers"."supplier_id") = ($1)`
	DeletePkSuppliers = `
		DELETE FROM "suppliers"
		WHERE ("suppliers"."supplier_id") = ($1)`
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
