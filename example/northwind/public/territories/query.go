// Code generated by github.com/hanpama/pgmg. DO NOT EDIT.
package territories

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
		INSERT INTO "public"."territories" (
			"territory_id",
			"territory_description",
			"region_id"
		)
		SELECT
			"territory_id",
			"territory_description",
			"region_id"
		FROM json_populate_recordset(null::"public"."territories", $1)`
	InsertReturningSQL = `
		INSERT INTO "public"."territories" (
			"territory_id",
			"territory_description",
			"region_id"
		)
		SELECT
			"territory_id",
			"territory_description",
			"region_id"
		FROM json_populate_recordset(null::"public"."territories", $1)
		RETURNING
			"territory_id",
			"territory_description",
			"region_id"`
)

func (k PkTerritories) selectSQL() Query {
	return selectPkTerritoriesQuery{k}
}
func (k PkTerritories) updateSQL(args ...attribute) Query {
	return query{UpdatePkTerritories, []interface{}{
		k.TerritoryID,
		string(mustMarshalJSON(Values(args))),
	}}
}
func (k PkTerritories) deleteSQL() Query {
	return query{DeletePkTerritories, []interface{}{
		k.TerritoryID,
	}}
}

const (
	SelectPkTerritories = `
		SELECT 
			"territory_id",
			"territory_description",
			"region_id"
		FROM "public"."territories" WHERE ("territory_id") = ($1) LIMIT 1`
	UpdatePkTerritories = `
		UPDATE "public"."territories" __ut__
		SET "territory_id" = COALESCE(__ch__."territory_id", __ut__."territory_id"),
			"territory_description" = COALESCE(__ch__."territory_description", __ut__."territory_description"),
			"region_id" = COALESCE(__ch__."region_id", __ut__."region_id")
		FROM (SELECT * FROM json_populate_record(null::"public"."territories", $2)) __ch__
		WHERE (__ut__."territory_id") = ($1)`
	DeletePkTerritories = `
		DELETE FROM "public"."territories"
		WHERE ("territory_id") = ($1)`
)

type selectPkTerritoriesQuery struct{ key PkTerritories }

func (q selectPkTerritoriesQuery) SQL() string         { return SelectPkTerritories }
func (q selectPkTerritoriesQuery) Args() []interface{} { return []interface{}{q.key.TerritoryID} }

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
