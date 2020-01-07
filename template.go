package main

import (
	"text/template"
)

var Tmpl = template.Must(template.New("pgmg").Parse(
	`{{- define "table_model" -}}
// Code generated by github.com/hanpama/pgmg. DO NOT EDIT.
package {{ .PackageName }}

import (
	"context"
	"encoding/json"
	"fmt"
	{{ range $k, $v := .Dependencies -}}
	"{{$k}}"
	{{end}}
)


{{ range $i, $m := .Models }}
// {{$m.CapitalName}} represents a row for table "{{$m.SQLName}}"
type {{$m.CapitalName}} struct {
	{{- range $i, $p := $m.Properties}}
	{{$p.CapitalName}} {{$p.GoInsertType}} ` + "`json:" + `"{{$p.SQLName}}"` + "`" + `
	{{- end}}
}
{{ end -}}

{{ range $i, $m := .Models }}

// New{{$m.CapitalName}}Repository creates a new *{{$m.CapitalName}}Repository
func New{{$m.CapitalName}}Repository(db PGMGDatabase) *{{$m.CapitalName}}Repository {
	return &{{$m.CapitalName}}Repository{db}
}

// {{$m.CapitalName}}Repository gets, saves and deletes rows of table "{{$m.SQLName}}"
type {{$m.CapitalName}}Repository struct {
	db PGMGDatabase
}
{{ range $j, $k := $m.Keys }}
// GetBy{{$k.CapitalName}} gets matching rows for given {{$k.CapitalName}} keys from table "{{$m.SQLName}}"
func (rep *{{$m.CapitalName}}Repository) GetBy{{$k.CapitalName}}(ctx context.Context, keys ...{{$k.CapitalName}}) (rows []*{{$m.CapitalName}}, err error) {
	return GetBy{{$k.CapitalName}}(ctx, rep.db, keys...)
}

// SaveBy{{$k.CapitalName}} upserts the given rows for table "{{$m.SQLName}}" checking uniqueness by contstraint "{{$k.SQLName}}"
func (rep *{{$m.CapitalName}}Repository) SaveBy{{$k.CapitalName}}(ctx context.Context, rows ...*{{$m.CapitalName}}) error {
	return SaveBy{{$k.CapitalName}}(ctx, rep.db, rows...)
}

// SaveAndReturnBy{{$k.CapitalName}} upserts the given rows for table "{{$m.SQLName}}" checking uniqueness by contstraint "{{$k.SQLName}}"
// It returns the new values and scan them into given row references.
func (rep *{{$m.CapitalName}}Repository) SaveAndReturnBy{{$k.CapitalName}}(ctx context.Context, rows ...*{{$m.CapitalName}}) ([]*{{$m.CapitalName}}, error) {
	return SaveAndReturnBy{{$k.CapitalName}}(ctx, rep.db, rows...)
}

// DeleteBy{{$k.CapitalName}} deletes matching rows by {{$k.CapitalName}} keys from table "{{$m.SQLName}}"
func (rep *{{$m.CapitalName}}Repository) DeleteBy{{$k.CapitalName}}(ctx context.Context, keys ...{{$k.CapitalName}}) (int64, error) {
	return DeleteBy{{$k.CapitalName}}(ctx, rep.db, keys...)
}
{{- end }}
// {{$m.CapitalName}}Condition is used for quering table "{{$m.SQLName}}"
type {{$m.CapitalName}}Condition struct {
	{{- range $i, $p := $m.Properties}}
	{{$p.CapitalName}} {{$p.FilterType}} ` + "`json:" + `"{{$p.SQLName}}"` + "`" + `
	{{- end}}
}

// Count{{$m.CapitalName}}Rows counts the number of rows which match the condition
func (rep *{{$m.CapitalName}}Repository) Count{{$m.CapitalName}}Rows(ctx context.Context, cond {{$m.CapitalName}}Condition) (int, error) {
	return Count{{$m.CapitalName}}Rows(ctx, rep.db, cond)
}

{{ end }}

{{ range $i, $m := .Models }}
{{ range $i, $p := $m.Properties }}
// {{$m.CapitalName}}{{$p.CapitalName}} represents column "{{$p.SQLName}}" of table "{{$m.SQLName}}"
type {{$m.CapitalName}}{{$p.CapitalName}} {{$p.GoInsertType}}
{{- end }}

// New{{$m.CapitalName}} creates a new row for table "{{$m.SQLName}}" with all column values
func New{{$m.CapitalName}}(
	{{- range $i, $p := $m.Properties }}
	{{$p.LowerName}} {{$m.CapitalName}}{{$p.CapitalName}},
	{{- end }}
) *{{$m.CapitalName}} {
	return &{{$m.CapitalName}}{
		{{- range $i, $p := $m.Properties }}
		({{$p.GoInsertType}})({{$p.LowerName}}),
		{{- end }}
	}
}

func (r *{{$m.CapitalName}}) receive() []interface{} {
	return []interface{}{
		{{- range $i, $p := $m.Properties -}}
		{{if $i}}, {{end}}&r.{{$p.CapitalName}}
		{{- end -}}
	}
}
{{ end }}


{{- range $i, $m := .Models }}
{{ range $j, $k := $m.Keys }}

// {{$k.CapitalName}} represents key defined by UNIQUE constraint "{{$k.SQLName}}" for table "{{$m.SQLName}}"
type {{$k.CapitalName}} struct {
	{{- range $h, $p := $k.Properties }}
	{{$p.CapitalName}} {{$p.GoSelectType}} ` + "`json:" + `"{{$p.SQLName}}"` + "`" + `
	{{- end }}
}

func (r *{{$m.CapitalName}}) {{$k.CapitalName}}() {{$k.CapitalName}} {
	k := {{$k.CapitalName}}{}
	{{- range $h, $p := $k.Properties }}
	{{- if $p.Default}}
	if r.{{$p.CapitalName}} != nil {
		k.{{$p.CapitalName}} = *r.{{$p.CapitalName}}
	}
	{{- else}}
	k.{{$p.CapitalName}} = r.{{$p.CapitalName}}
	{{- end}}
	{{- end }}
	return k
}

var SQLGetBy{{$k.CapitalName}} = ` + "`" + `
	WITH __key AS (
		SELECT ROW_NUMBER() over () __keyindex,
			{{ range $h, $p := $k.Properties }}{{if $h }}, {{end}}{{$p.SQLName}}{{end}}
		FROM json_populate_recordset(null::"{{$m.Schema}}"."{{$m.SQLName}}", $1)
	)
	SELECT {{ range $h, $p := $m.Properties }}{{if $h }}, {{end}}{{$p.SQLName}}{{end}}
	FROM __key JOIN "{{$m.Schema}}"."{{$m.SQLName}}" AS __table USING ({{ range $h, $p := $k.Properties }}{{if $h }}, {{end}}{{$p.SQLName}}{{end}})
	ORDER BY __keyindex
` + "`" + `

// GetBy{{$k.CapitalName}} gets matching rows for given {{$k.CapitalName}} keys from table "{{$m.SQLName}}"
func GetBy{{$k.CapitalName}}(ctx context.Context, db PGMGDatabase, keys ...{{$k.CapitalName}}) (rows []*{{$m.CapitalName}}, err error) {
	var b []byte
	if b, err = json.Marshal(keys); err != nil {
		return nil, err
	}
	rows = make([]*{{$m.CapitalName}}, len(keys))
	if _, err = db.QueryScan(ctx, func(i int) []interface{} {
		rows[i] = &{{$m.CapitalName}}{}
		return rows[i].receive()
	}, SQLGetBy{{$k.CapitalName}}, string(b)); err != nil {
		return nil, err
	}
	for i := 0; i < len(keys); i++ {
		if rows[i] == nil {
			break
		} else if rows[i].{{$k.CapitalName}}() != keys[i] {
			copy(rows[i+1:], rows[i:])
			rows[i] = nil
		}
	}
	return rows, nil
}

var SQLSaveBy{{$k.CapitalName}} = ` + "`" + `
	WITH __values AS (
		SELECT
			{{ range $h, $p := $m.Properties }}{{if $h }},
			{{end -}}
			{{- if $p.Default}}COALESCE(__input.{{$p.SQLName}}, {{$p.Default}}) {{$p.SQLName}}
			{{- else}}__input.{{$p.SQLName}}
			{{- end}}
		{{- end}}
		FROM json_populate_recordset(null::"{{$m.Schema}}"."{{$m.SQLName}}", $1) __input
	)
	INSERT INTO "{{$m.Schema}}"."{{$m.SQLName}}" SELECT * FROM __values
	ON CONFLICT ({{ range $h, $p := $k.Properties }}{{if $h }}, {{end}}{{$p.SQLName}}{{end}}) DO UPDATE
		SET ({{ range $h, $p := $m.Properties }}{{if $h }}, {{end}}{{$p.SQLName}}{{end}}) = (
			SELECT {{ range $h, $p := $m.Properties }}{{if $h }}, {{end}}{{$p.SQLName}}{{end}} FROM __values
		)
` + "`" + `

// SaveBy{{$k.CapitalName}} upserts the given rows for table "{{$m.SQLName}}" checking uniqueness by contstraint "{{$k.SQLName}}"
func SaveBy{{$k.CapitalName}}(ctx context.Context, db PGMGDatabase, rows ...*{{$m.CapitalName}}) error {
	return execJSONSave(ctx, db, SQLSaveBy{{$k.CapitalName}}, rows, len(rows))
}

var SQLSaveAndReturnBy{{$k.CapitalName}} = SQLSaveBy{{$k.CapitalName}} + " RETURNING {{ range $h, $p := $m.Properties }}{{if $h }}, {{end}}{{$p.SQLName}}{{end}}"

// SaveAndReturnBy{{$k.CapitalName}} upserts the given rows for table "{{$m.SQLName}}" checking uniqueness by contstraint "{{$k.SQLName}}"
// It returns the new values and scan them into given row references.
func SaveAndReturnBy{{$k.CapitalName}}(ctx context.Context, db PGMGDatabase, rows ...*{{$m.CapitalName}}) ([]*{{$m.CapitalName}}, error) {
	return rows, execJSONSaveAndReturn(ctx, db, func(i int) []interface{} { return rows[i].receive() }, SQLSaveAndReturnBy{{$k.CapitalName}}, rows, len(rows))
}

var SQLDeleteBy{{$k.CapitalName}} = ` + "`" + `
WITH __key AS (SELECT * FROM json_populate_recordset(null::"{{$m.Schema}}"."{{$m.SQLName}}", $1))
DELETE FROM "{{$m.Schema}}"."{{$m.SQLName}}" AS __table
	USING __key
	WHERE {{ range $h, $p := $k.Properties }}{{if $h }}  AND {{end}}(__key.{{$p.SQLName}} = __table.{{$p.SQLName}})
	{{ end -}}
` + "`" + `

// DeleteBy{{$k.CapitalName}} deletes matching rows by {{$k.CapitalName}} keys from table "{{$m.SQLName}}"
func DeleteBy{{$k.CapitalName}}(ctx context.Context, db PGMGDatabase, keys ...{{$k.CapitalName}}) (int64, error) {
	b, err := json.Marshal(keys);
	if err != nil {
		return 0, err
	}
	return db.Exec(ctx, SQLDeleteBy{{$k.CapitalName}}, string(b))
}

{{end}}{{ end}}


{{ range $i, $m := .Models }}
// Count{{$m.CapitalName}}Rows counts the number of rows which match the condition
func Count{{$m.CapitalName}}Rows(ctx context.Context, db PGMGDatabase, cond {{$m.CapitalName}}Condition) (count int, err error) {
	var arg1 []byte
	if arg1, err = json.Marshal(cond); err != nil {
		return 0, err
	}
	_, err = db.QueryScan(ctx, func(int) []interface{} { return []interface{}{&count } }, ` + "`" + `
		SELECT count(*) FROM "{{$m.Schema}}"."{{$m.SQLName}}" AS __t
		WHERE TRUE
			{{- range $i, $p := $m.Properties }}
			AND (($1::json->>'{{$p.SQLName}}' IS NULL) OR CAST($1::json->>'{{$p.SQLName}}' AS {{$p.SQLType}}) = __t."{{$p.SQLName}}")
			{{- end }}
	` + "`" + `, arg1)
	return count, err
}
{{ end -}}

func execJSONSave(ctx context.Context, db PGMGDatabase, sql string, rows interface{}, ern int) (err error) {
	var arg1 []byte
	if arg1, err = json.Marshal(rows); err != nil {
		return err
	}
	if affected, err := db.Exec(ctx, sql, string(arg1)); err != nil {
		return err
	} else if affected != int64(ern) {
		return ErrUnexpectedRowNumberAffected
	}
	return nil
}

func execJSONSaveAndReturn(ctx context.Context, db PGMGDatabase, receive func(int) []interface{}, sql string, rows interface{}, ern int) (err error) {
	var arg1 []byte
	if arg1, err = json.Marshal(rows); err != nil {
		return err
	}
	if affected, err := db.QueryScan(ctx, receive, sql, string(arg1)); err != nil {
		return err
	} else if affected != ern {
		return ErrUnexpectedRowNumberAffected
	}
	return nil
}

// PGMGDatabase represents PostgresQL database
type PGMGDatabase interface {
	QueryScan(ctx context.Context, receiver func(int) []interface{}, sql string, args ...interface{}) (int, error)
	Exec(ctx context.Context, sql string, args ...interface{}) (int64, error)
}

var ErrUnexpectedRowNumberAffected = fmt.Errorf("unexpected row number affected")
var ErrInvalidConditions = fmt.Errorf("invalid conditions")

{{ end -}}
	`,
))
