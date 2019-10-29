package dsl

import (
	"bytes"
	"text/template"
)

func DeleteFrom(table Table) *DeleteQuery {
	return &DeleteQuery{Table: table}
}

// DeleteQuery represents SQL Delete
// https://www.postgresql.org/docs/current/sql-delete.html
type DeleteQuery struct {
	CTE        []CTE
	Conditions []Expression
	Table      Table
	UsingList  []FromItem
	Outputs    []Expression
}

var deleteQueryTmpl = template.Must(template.New("Insert").Parse(`DELETE FROM {{.Table.AsTable -}}
{{ range $i, $c := .Conditions}}{{if eq $i 0}} WHERE {{else}} AND {{end}}{{$c.Expr}}{{end -}}
`))

func (q *DeleteQuery) Statement() string {
	var buff bytes.Buffer
	err := deleteQueryTmpl.Execute(&buff, q)
	if err != nil {
		panic(err)
	}
	return string(buff.Bytes())
}
func (q *DeleteQuery) Where(conds ...Expression) *DeleteQuery {
	q.Conditions = append(q.Conditions, conds...)
	return q
}

func (q *DeleteQuery) Wheref(format string, a ...interface{}) *DeleteQuery {
	q.Conditions = append(q.Conditions, NewExprf(format, a...))
	return q
}

func (q *DeleteQuery) Using(items ...FromItem) *DeleteQuery {
	for _, item := range items {
		q.UsingList = append(q.UsingList, item)
	}
	return q
}

func (q *DeleteQuery) Usingf(format string, a ...interface{}) *DeleteQuery {
	q.UsingList = append(q.UsingList, NewFromf(format, a...)) //
	return q
}

func (q *DeleteQuery) Returning(a ...interface{}) *DeleteQuery {
	return q
}
