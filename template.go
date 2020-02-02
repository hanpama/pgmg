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

// {{$m.CapitalName}}Rows represents multiple rows for table "{{$m.SQLName}}"
type {{$m.CapitalName}}Rows []*{{$m.CapitalName}}

{{ end -}}

// NewPGMGRepository creates a new PGMGRepository
func NewPGMGRepository(db PGMGDatabase) *PGMGRepository {
	return &PGMGRepository{db}
}

// PGMGRepository provides methods which get, insert, save and delete rows in db
type PGMGRepository struct {
	db PGMGDatabase
}

{{ range $i, $m := .Models }}

func (rep *PGMGRepository) Insert{{$m.CapitalName}}Rows(ctx context.Context, rows ...*{{$m.CapitalName}}) error {
	return Insert{{$m.CapitalName}}Rows(ctx, rep.db, rows...)
}

func (rep *PGMGRepository) InsertAndReturn{{$m.CapitalName}}Rows(ctx context.Context, rows ...*{{$m.CapitalName}}) ({{$m.CapitalName}}Rows, error) {
	return InsertAndReturn{{$m.CapitalName}}Rows(ctx, rep.db, rows...)
}

{{ range $j, $k := $m.Keys }}
// GetBy{{$k.CapitalName}} gets matching rows for given {{$k.CapitalName}} keys from table "{{$m.SQLName}}"
func (rep *PGMGRepository) GetBy{{$k.CapitalName}}(ctx context.Context, keys ...{{$k.CapitalName}}) (rows {{$m.CapitalName}}Rows, err error) {
	return GetBy{{$k.CapitalName}}(ctx, rep.db, keys...)
}

// SaveBy{{$k.CapitalName}} upserts the given rows for table "{{$m.SQLName}}" checking uniqueness by contstraint "{{$k.SQLName}}"
func (rep *PGMGRepository) SaveBy{{$k.CapitalName}}(ctx context.Context, rows ...*{{$m.CapitalName}}) error {
	return SaveBy{{$k.CapitalName}}(ctx, rep.db, rows...)
}

// SaveAndReturnBy{{$k.CapitalName}} upserts the given rows for table "{{$m.SQLName}}" checking uniqueness by contstraint "{{$k.SQLName}}"
// It returns the new values and scan them into given row references.
func (rep *PGMGRepository) SaveAndReturnBy{{$k.CapitalName}}(ctx context.Context, rows ...*{{$m.CapitalName}}) ({{$m.CapitalName}}Rows, error) {
	return SaveAndReturnBy{{$k.CapitalName}}(ctx, rep.db, rows...)
}

// DeleteBy{{$k.CapitalName}} deletes matching rows by {{$k.CapitalName}} keys from table "{{$m.SQLName}}"
func (rep *PGMGRepository) DeleteBy{{$k.CapitalName}}(ctx context.Context, keys ...{{$k.CapitalName}}) (int64, error) {
	return DeleteBy{{$k.CapitalName}}(ctx, rep.db, keys...)
}
{{- end }}
// {{$m.CapitalName}}Condition is used for quering table "{{$m.SQLName}}"
type {{$m.CapitalName}}Condition struct {
	{{- range $i, $p := $m.Properties}}
	{{$p.CapitalName}} {{$p.FilterType}} ` + "`json:" + `"{{$p.SQLName}}"` + "`" + `
	{{- end}}
}

// Find{{$m.CapitalName}}Rows find the rows matching the condition from table "{{$m.SQLName}}"
func (rep *PGMGRepository) Find{{$m.CapitalName}}Rows(ctx context.Context, cond {{$m.CapitalName}}Condition) ({{$m.CapitalName}}Rows, error) {
	return Find{{$m.CapitalName}}Rows(ctx, rep.db, cond)
}
// Delete{{$m.CapitalName}}Rows delete the rows matching the condition from table "{{$m.SQLName}}"
func (rep *PGMGRepository) Delete{{$m.CapitalName}}Rows(ctx context.Context, cond {{$m.CapitalName}}Condition) (afftected int64, err error) {
	return Delete{{$m.CapitalName}}Rows(ctx, rep.db, cond)
}
// Count{{$m.CapitalName}}Rows counts the number of rows matching the condition from table "{{$m.SQLName}}"
func (rep *PGMGRepository) Count{{$m.CapitalName}}Rows(ctx context.Context, cond {{$m.CapitalName}}Condition) (int, error) {
	return Count{{$m.CapitalName}}Rows(ctx, rep.db, cond)
}

{{ end }}

{{ range $i, $m := .Models }}
{{ range $i, $p := $m.Properties }}
// {{$m.CapitalName}}{{$p.CapitalName}} represents value type of column "{{$p.SQLName}}" of table "{{$m.SQLName}}"
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

func (r *{{$m.CapitalName}}) ReceiveRow() []interface{} {
	return []interface{}{
		{{- range $i, $p := $m.Properties -}}
		{{if $i}}, {{end}}&r.{{$p.CapitalName}}
		{{- end -}}
	}
}

// ReceiveRows returns pointer slice to receive data for the row on index i
func (rs *{{$m.CapitalName}}Rows) ReceiveRows(i int) []interface{} {
	if cap(*rs) <= i {
		source := *rs
		*rs = make({{$m.CapitalName}}Rows, i+1)
		copy(*rs, source)
	}
	if (*rs)[i] == nil {
		(*rs)[i] = new({{$m.CapitalName}})
	}
	return (*rs)[i].ReceiveRow()
}

{{ end }}


{{- range $i, $m := .Models }}

var sqlInsert{{$m.CapitalName}}Rows = ` + "`" + `
	WITH __values AS (
		SELECT
			{{ range $h, $p := $m.Properties }}{{if $h }},
			{{end -}}
			{{- if $p.Default}}COALESCE(__input."{{$p.SQLName}}", {{$p.Default}}) "{{$p.SQLName}}"
			{{- else}}__input."{{$p.SQLName}}"
			{{- end}}
		{{- end}}
		FROM json_populate_recordset(null::"{{$m.Schema}}"."{{$m.SQLName}}", $1) __input
	)
	INSERT INTO "{{$m.Schema}}"."{{$m.SQLName}}" AS _t ({{ range $h, $p := $m.Properties }}{{if $h }}, {{end}}"{{$p.SQLName}}"{{end}})
	SELECT {{ range $h, $p := $m.Properties }}{{if $h }}, {{end}}"{{$p.SQLName}}"{{end}} FROM __values` + "`" + `

func Insert{{$m.CapitalName}}Rows(ctx context.Context, db PGMGDatabase, inputs ...*{{$m.CapitalName}}) (err error) {
	if err = execJSON(ctx, db, sqlInsert{{$m.CapitalName}}Rows, inputs, len(inputs)); err != nil {
		return fmt.Errorf("%w( Insert{{$m.CapitalName}}Rows, %w)", ErrPGMG, err)
	}
	return nil
}

var sqlReturning{{$m.CapitalName}}Rows = ` + "`" + `
	RETURNING {{ range $h, $p := $m.Properties }}{{if $h }}, {{end}}{{$p.SQLName}}{{end}}
` + "`" + `

var sqlInsertAndReturn{{$m.CapitalName}}Rows = sqlInsert{{$m.CapitalName}}Rows + sqlReturning{{$m.CapitalName}}Rows

func InsertAndReturn{{$m.CapitalName}}Rows(ctx context.Context, db PGMGDatabase, inputs ...*{{$m.CapitalName}}) (rows {{$m.CapitalName}}Rows, err error) {
	rows = inputs
	err = execJSONAndReturn(ctx, db, rows.ReceiveRows, sqlInsertAndReturn{{$m.CapitalName}}Rows, rows, len(rows))
	if err != nil {
		return rows, fmt.Errorf("%w(SQLInsertAndReturn{{$m.CapitalName}}Rows, %w)", ErrPGMG, err)
	}
	return rows, nil
}

var sqlFind{{$m.CapitalName}}Rows = ` + "`" + `
	SELECT {{range $i, $p := $m.Properties }}{{if $i}}, {{end}}__t.{{$p.SQLName}}{{end}}
	FROM "{{$m.Schema}}"."{{$m.SQLName}}" AS __t
	WHERE TRUE
		{{- range $i, $p := $m.Properties }}
		AND (($1::json->>'{{$p.SQLName}}' IS NULL) OR CAST($1::json->>'{{$p.SQLName}}' AS {{$p.SQLType}}) = __t."{{$p.SQLName}}")
		{{- end }}
` + "`" + `

// Find{{$m.CapitalName}}Rows find the rows matching the condition from table "{{$m.SQLName}}"
func Find{{$m.CapitalName}}Rows(ctx context.Context, db PGMGDatabase, cond {{$m.CapitalName}}Condition) (rows {{$m.CapitalName}}Rows, err error) {
	var arg1 []byte
	if arg1, err = json.Marshal(cond); err != nil {
		return nil, err
	}
	if _, err = db.QueryScan(ctx, rows.ReceiveRows, sqlFind{{$m.CapitalName}}Rows, arg1); err != nil {
		return nil, err
	}
	return rows, nil
}

var sqlDelete{{$m.CapitalName}}Rows = ` + "`" + `
	DELETE FROM "{{$m.Schema}}"."{{$m.SQLName}}" AS __t
	WHERE TRUE
		{{- range $i, $p := $m.Properties }}
		AND (($1::json->>'{{$p.SQLName}}' IS NULL) OR CAST($1::json->>'{{$p.SQLName}}' AS {{$p.SQLType}}) = __t."{{$p.SQLName}}")
		{{- end }}
` + "`" + `

// Delete{{$m.CapitalName}}Rows delete the rows matching the condition from table "{{$m.SQLName}}"
func Delete{{$m.CapitalName}}Rows(ctx context.Context, db PGMGDatabase, cond {{$m.CapitalName}}Condition) (afftected int64, err error) {
	var arg1 []byte
	if arg1, err = json.Marshal(cond); err != nil {
		return 0, err
	}
	return db.ExecCountingAffected(ctx, sqlDelete{{$m.CapitalName}}Rows, arg1)
}

var sqlCount{{$m.CapitalName}}Rows = ` + "`" + `
	SELECT count(*) FROM "{{$m.Schema}}"."{{$m.SQLName}}" AS __t
	WHERE TRUE
		{{- range $i, $p := $m.Properties }}
		AND (($1::json->>'{{$p.SQLName}}' IS NULL) OR CAST($1::json->>'{{$p.SQLName}}' AS {{$p.SQLType}}) = __t."{{$p.SQLName}}")
		{{- end }}
` + "`" + `

// Count{{$m.CapitalName}}Rows counts the number of rows matching the condition from table "{{$m.SQLName}}"
func Count{{$m.CapitalName}}Rows(ctx context.Context, db PGMGDatabase, cond {{$m.CapitalName}}Condition) (count int, err error) {
	var arg1 []byte
	if arg1, err = json.Marshal(cond); err != nil {
		return 0, err
	}
	_, err = db.QueryScan(ctx, func(int) []interface{} { return []interface{}{&count } }, sqlCount{{$m.CapitalName}}Rows, arg1)
	return count, err
}

{{ range $j, $k := $m.Keys }}

// {{$k.CapitalName}} represents the key defined by UNIQUE constraint "{{$k.SQLName}}" for table "{{$m.SQLName}}"
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

func (rs {{$m.CapitalName}}Rows) {{$k.CapitalName}}Slice() (keys []{{$k.CapitalName}}) {
	keys = make([]{{$k.CapitalName}}, len(rs))
	for i, r := range rs {
		keys[i] = r.{{$k.CapitalName}}()
	}
	return keys
}

var sqlGetBy{{$k.CapitalName}} = ` + "`" + `
	WITH __key AS (
		SELECT ROW_NUMBER() over () __keyindex,
			{{ range $h, $p := $k.Properties }}{{if $h }}, {{end}}{{$p.SQLName}}{{end}}
		FROM json_populate_recordset(null::"{{$m.Schema}}"."{{$m.SQLName}}", $1)
	)
	SELECT {{ range $h, $p := $m.Properties }}{{if $h }}, {{end}}"{{$p.SQLName}}"{{end}}
	FROM __key JOIN "{{$m.Schema}}"."{{$m.SQLName}}" AS __table USING ({{ range $h, $p := $k.Properties }}{{if $h }}, {{end}}"{{$p.SQLName}}"{{end}})
	ORDER BY __keyindex
` + "`" + `

// GetBy{{$k.CapitalName}} gets matching rows for given {{$k.CapitalName}} keys from table "{{$m.SQLName}}"
func GetBy{{$k.CapitalName}}(ctx context.Context, db PGMGDatabase, keys ...{{$k.CapitalName}}) (rows {{$m.CapitalName}}Rows, err error) {
	var b []byte
	if b, err = json.Marshal(keys); err != nil {
		return nil, fmt.Errorf("%w(GetBy{{$k.CapitalName}}, %w)", ErrPGMG, err)
	}
	rows = make({{$m.CapitalName}}Rows, len(keys))
	if _, err = db.QueryScan(ctx, rows.ReceiveRows, sqlGetBy{{$k.CapitalName}}, b); err != nil {
		return nil, fmt.Errorf("%w(GetBy{{$k.CapitalName}}, %w)", ErrPGMG, err)
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

var sqlSaveBy{{$k.CapitalName}} = sqlInsert{{$m.CapitalName}}Rows + ` + "`" + `
	ON CONFLICT ({{ range $h, $p := $k.Properties }}{{if $h }}, {{end}}"{{$p.SQLName}}"{{end}}) DO UPDATE
		SET ({{ range $h, $p := $m.Properties }}{{if $h }}, {{end}}"{{$p.SQLName}}"{{end}}) = (
			SELECT {{ range $h, $p := $m.Properties }}{{if $h }}, {{end}}"{{$p.SQLName}}"{{end}} FROM __values
			WHERE {{ range $h, $p := $k.Properties }}{{if $h }}
				AND {{end}}__values."{{$p.SQLName}}" = _t."{{$p.SQLName}}"{{end}}
		)
` + "`" + `

// SaveBy{{$k.CapitalName}} upserts the given rows for table "{{$m.SQLName}}" checking uniqueness by contstraint "{{$k.SQLName}}"
func SaveBy{{$k.CapitalName}}(ctx context.Context, db PGMGDatabase, rows ...*{{$m.CapitalName}}) (err error) {
	if err = execJSON(ctx, db, sqlSaveBy{{$k.CapitalName}}, rows, len(rows)); err != nil {
		return fmt.Errorf("%w(SaveBy{{$k.CapitalName}}, %w)", ErrPGMG, err)
	}
	return nil
}

var sqlSaveAndReturnBy{{$k.CapitalName}} = sqlSaveBy{{$k.CapitalName}} + sqlReturning{{$m.CapitalName}}Rows

// SaveAndReturnBy{{$k.CapitalName}} upserts the given rows for table "{{$m.SQLName}}" checking uniqueness by contstraint "{{$k.SQLName}}"
// It returns the new values and scan them into given row references.
func SaveAndReturnBy{{$k.CapitalName}}(ctx context.Context, db PGMGDatabase, inputs ...*{{$m.CapitalName}}) (rows {{$m.CapitalName}}Rows, err error) {
	rows = inputs
	err = execJSONAndReturn(ctx, db, rows.ReceiveRows, sqlSaveAndReturnBy{{$k.CapitalName}}, rows, len(rows))
	if err != nil {
		return rows, fmt.Errorf("%w(SaveAndReturnBy{{$k.CapitalName}}, %w)", ErrPGMG, err)
	}
	return rows, nil
}

var sqlDeleteBy{{$k.CapitalName}} = ` + "`" + `
WITH __key AS (SELECT {{ range $h, $p := $k.Properties }}{{if $h}}, {{end}}{{$p.SQLName}}{{end}} FROM json_populate_recordset(null::"{{$m.Schema}}"."{{$m.SQLName}}", $1))
DELETE FROM "{{$m.Schema}}"."{{$m.SQLName}}" AS __table
	USING __key
	WHERE {{ range $h, $p := $k.Properties }}{{if $h }}  AND {{end}}(__key."{{$p.SQLName}}" = __table."{{$p.SQLName}}")
	{{ end -}}
` + "`" + `

// DeleteBy{{$k.CapitalName}} deletes matching rows by {{$k.CapitalName}} keys from table "{{$m.SQLName}}"
func DeleteBy{{$k.CapitalName}}(ctx context.Context, db PGMGDatabase, keys ...{{$k.CapitalName}}) (affected int64, err error) {
	b, err := json.Marshal(keys);
	if err != nil {
		return affected, fmt.Errorf("%w(DeleteBy{{$k.CapitalName}}, %w)", ErrPGMG, err)
	}
	if affected, err = db.ExecCountingAffected(ctx, sqlDeleteBy{{$k.CapitalName}}, b); err != nil {
		return affected, fmt.Errorf("%w(DeleteBy{{$k.CapitalName}}, %w)", ErrPGMG, err)
	}
	return affected, nil
}

{{end}}{{ end}}


func execJSON(ctx context.Context, db PGMGDatabase, sql string, rows interface{}, ern int) (err error) {
	var arg1 []byte
	if arg1, err = json.Marshal(rows); err != nil {
		return err
	}
	if affected, err := db.ExecCountingAffected(ctx, sql, arg1); err != nil {
		return err
	} else if affected != int64(ern) {
		return ErrUnexpectedRowNumberAffected
	}
	return nil
}

func execJSONAndReturn(ctx context.Context, db PGMGDatabase, receive func(int) []interface{}, sql string, rows interface{}, ern int) (err error) {
	var arg1 []byte
	if arg1, err = json.Marshal(rows); err != nil {
		return err
	}
	if affected, err := db.QueryScan(ctx, receive, sql, arg1); err != nil {
		return err
	} else if affected != ern {
		return ErrUnexpectedRowNumberAffected
	}
	return nil
}

// PGMGDatabase represents PostgresQL database
type PGMGDatabase interface {
	QueryScan(ctx context.Context, receiver func(int) []interface{}, sql string, args ...interface{}) (int, error)
	ExecCountingAffected(ctx context.Context, sql string, args ...interface{}) (int64, error)
}

var ErrUnexpectedRowNumberAffected = fmt.Errorf("pgmg: unexpected row number affected")
var ErrPGMG = fmt.Errorf("pgmg: error")

{{ end -}}
	`,
))
