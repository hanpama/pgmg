// Code generated by github.com/hanpama/pgmg. DO NOT EDIT.
package us_states

func Insert(vss ...Values) Query {
	return query{InsertSQL, []interface{}{string(mustMarshalJSON(vss))}}
}
func InsertReturning(vss ...Values) Query {
	return query{InsertReturningSQL, []interface{}{string(mustMarshalJSON(vss))}}
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
		INSERT INTO "public"."us_states" (
			"state_id",
			"state_name",
			"state_abbr",
			"state_region"
		)
		SELECT
			"state_id",
			"state_name",
			"state_abbr",
			"state_region"
		FROM json_populate_recordset(null::"public"."us_states", $1)`
	InsertReturningSQL = `
		INSERT INTO "public"."us_states" (
			"state_id",
			"state_name",
			"state_abbr",
			"state_region"
		)
		SELECT
			"state_id",
			"state_name",
			"state_abbr",
			"state_region"
		FROM json_populate_recordset(null::"public"."us_states", $1)
		RETURNING
			"state_id",
			"state_name",
			"state_abbr",
			"state_region"`
)

func (k PkUsstates) selectSQL() Query {
	return selectPkUsstatesQuery{k}
}
func (k PkUsstates) updateSQL(args ...attribute) Query {
	return query{UpdatePkUsstates, []interface{}{
		k.StateID,
		string(mustMarshalJSON(Values(args))),
	}}
}
func (k PkUsstates) deleteSQL() Query {
	return query{DeletePkUsstates, []interface{}{
		k.StateID,
	}}
}

const (
	SelectPkUsstates = `
		SELECT 
			"state_id",
			"state_name",
			"state_abbr",
			"state_region"
		FROM "public"."us_states" WHERE ("state_id") = ($1) LIMIT 1`
	UpdatePkUsstates = `
		UPDATE "public"."us_states" __ut__
		SET "state_id" = COALESCE(__ch__."state_id", __ut__."state_id"),
			"state_name" = COALESCE(__ch__."state_name", __ut__."state_name"),
			"state_abbr" = COALESCE(__ch__."state_abbr", __ut__."state_abbr"),
			"state_region" = COALESCE(__ch__."state_region", __ut__."state_region")
		FROM (SELECT * FROM json_populate_record(null::"public"."us_states", $2)) __ch__
		WHERE (__ut__."state_id") = ($1)`
	DeletePkUsstates = `
		DELETE FROM "public"."us_states"
		WHERE ("state_id") = ($1)`
)

type selectPkUsstatesQuery struct{ key PkUsstates }

func (q selectPkUsstatesQuery) SQL() string         { return SelectPkUsstates }
func (q selectPkUsstatesQuery) Args() []interface{} { return []interface{}{q.key.StateID} }

type key interface {
	selectSQL() Query
	updateSQL(args ...attribute) Query
	deleteSQL() Query
}

type Query interface {
	SQL() string
	Args() []interface{}
}

type query struct {
	sql  string
	args []interface{}
}

func (q query) SQL() string         { return q.sql }
func (q query) Args() []interface{} { return q.args }
