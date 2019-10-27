package dsl

import (
	"bytes"
	"fmt"
	"text/template"
)

func InsertInto(table Table, columns ...interface{}) *InsertQuery {
	return &InsertQuery{Table: table, Columns: columns}
}

type InsertQuery struct {
	CTE            []CTE
	Table          Table
	Columns        []interface{}
	ValuesList     []string
	Query          *SelectQuery
	Overriding     string // zero | SYSTEM VALUE | USER VALUE
	ConflictAction *ConflictAction
	Outputs        []Expression
	IsReturning    bool
}

var insertQueryTmpl = template.Must(template.New("Insert").Parse(`INSERT INTO {{.Table.AsTable -}}
{{if len .Columns}} ({{range $i, $c := .Columns}}{{if $i}}, {{end}}{{$c.AsColumn}}{{end}}){{end -}}
{{range $i, $v := .ValuesList}}{{if eq $i 0}} VALUES {{else}}, {{end}}{{$v}}{{end -}}
{{if .Query}} {{.Query.Statement}}{{end -}}
{{if .IsReturning}} RETURNING {{if not .Outputs}}*{{end -}}
{{range $i, $o := .Outputs}}{{if $i}}, {{end}}{{$o.Expr}}{{end -}}
{{end -}}
`))

func (q *InsertQuery) Statement() string {
	var buff bytes.Buffer
	err := insertQueryTmpl.Execute(&buff, q)
	if err != nil {
		panic(err)
	}
	return string(buff.Bytes())
}

func (q *InsertQuery) Valuesf(format string, a ...interface{}) *InsertQuery {
	q.ValuesList = append(q.ValuesList, fmt.Sprintf(format, a...))
	return q
}

func (q *InsertQuery) Select(query *SelectQuery) *InsertQuery {
	q.Query = query
	return q
}

func (q *InsertQuery) OnConflict(colNames ...string) *InsertQuery {
	panic("Unimplemented")

}
func (q *InsertQuery) OnConflictOn(constraintName string) *InsertQuery {
	panic("Unimplemented")
}

func (q *InsertQuery) Returning(exprs ...interface{}) *InsertQuery {
	q.IsReturning = true
	q.Outputs = append(q.Outputs, WrapExprs(exprs...)...)
	return q
}
func (q *InsertQuery) Returningf(format string, a ...interface{}) *InsertQuery {
	panic("Unimplemented")
}
