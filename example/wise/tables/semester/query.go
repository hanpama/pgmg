// Code generated by github.com/hanpama/pgmg. DO NOT EDIT.
package semester

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
		INSERT INTO "wise"."semester" (
			"id",
			"year",
			"season"
		)
		SELECT
			COALESCE("id", nextval('wise.semester_id_seq'::regclass)),
			"year",
			"season"
		FROM json_populate_recordset(null::"wise"."semester", $1)`
	InsertReturningSQL = `
		INSERT INTO "wise"."semester" (
			"id",
			"year",
			"season"
		)
		SELECT
			COALESCE("id", nextval('wise.semester_id_seq'::regclass)),
			"year",
			"season"
		FROM json_populate_recordset(null::"wise"."semester", $1)
		RETURNING
			"id",
			"year",
			"season"`
)

func (k SemesterPkey) selectSQL() Query {
	return Query{SelectSemesterPkey, []interface{}{
		k.ID,
	}}
}
func (k SemesterPkey) updateSQL(args ...attribute) Query {
	return Query{UpdateSemesterPkey, []interface{}{
		k.ID,
		string(mustMarshalJSON(Values(args))),
	}}
}
func (k SemesterPkey) deleteSQL() Query {
	return Query{DeleteSemesterPkey, []interface{}{
		k.ID,
	}}
}

const (
	SelectSemesterPkey = `
		SELECT * FROM "wise"."semester" WHERE ("id") = ($1) LIMIT 1
		`
	UpdateSemesterPkey = `
		UPDATE "wise"."semester" __ut__
		SET "id" = COALESCE(__ch__."id", __ut__."id"),
			"year" = COALESCE(__ch__."year", __ut__."year"),
			"season" = COALESCE(__ch__."season", __ut__."season")
		FROM (SELECT * FROM json_populate_record(null::"wise"."semester", $2)) __ch__
		WHERE (__ut__."id") = ($1)`
	DeleteSemesterPkey = `
		DELETE FROM "wise"."semester"
		WHERE ("id") = ($1)`
)

func (k SemesterYearSeasonKey) selectSQL() Query {
	return Query{SelectSemesterYearSeasonKey, []interface{}{
		k.Year,
		k.Season,
	}}
}
func (k SemesterYearSeasonKey) updateSQL(args ...attribute) Query {
	return Query{UpdateSemesterYearSeasonKey, []interface{}{
		k.Year,
		k.Season,
		string(mustMarshalJSON(Values(args))),
	}}
}
func (k SemesterYearSeasonKey) deleteSQL() Query {
	return Query{DeleteSemesterYearSeasonKey, []interface{}{
		k.Year,
		k.Season,
	}}
}

const (
	SelectSemesterYearSeasonKey = `
		SELECT * FROM "wise"."semester" WHERE ("year", "season") = ($1, $2) LIMIT 1
		`
	UpdateSemesterYearSeasonKey = `
		UPDATE "wise"."semester" __ut__
		SET "id" = COALESCE(__ch__."id", __ut__."id"),
			"year" = COALESCE(__ch__."year", __ut__."year"),
			"season" = COALESCE(__ch__."season", __ut__."season")
		FROM (SELECT * FROM json_populate_record(null::"wise"."semester", $2)) __ch__
		WHERE (__ut__."year", __ut__."season") = ($1, $2)`
	DeleteSemesterYearSeasonKey = `
		DELETE FROM "wise"."semester"
		WHERE ("year", "season") = ($1, $2)`
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
