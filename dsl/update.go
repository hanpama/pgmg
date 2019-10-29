package dsl

import (
	"bytes"
	"text/template"
)

func Update(table Table) *UpdateQuery {
	return &UpdateQuery{Table: table}
}

type Setter interface {
	AsSetter() string
}

type UpdateQuery struct {
	CTE        []CTE
	Conditions []Expression
	Table      Table
	Setters    []Setter
	FromList   []FromItem
	Outputs    []Expression
}

var updateQueryTmpl = template.Must(template.New("Select").Parse(`UPDATE {{.Table.AsTable -}}
{{ range $i, $s := .Setters}}{{if $i}}, {{else}} SET {{end}}{{$s.AsSetter}}{{end -}}
{{ range $i, $f := .FromList}}{{if eq $i 0}} FROM {{else}}, {{end}}{{$f.From}}{{end -}}
{{ range $i, $c := .Conditions}}{{if eq $i 0}} WHERE {{else}} AND {{end}}{{$c.Expr}}{{end -}}
`))

func (q *UpdateQuery) Statement() string {
	var buff bytes.Buffer
	err := updateQueryTmpl.Execute(&buff, q)
	if err != nil {
		panic(err)
	}
	return string(buff.Bytes())
}

func (q *UpdateQuery) Set(setters ...Setter) *UpdateQuery {
	q.Setters = append(q.Setters, setters...)
	return q
}

func (q *UpdateQuery) Where(conds ...Expression) *UpdateQuery {
	q.Conditions = append(q.Conditions, conds...)
	return q
}

func (q *UpdateQuery) From(items ...FromItem) *UpdateQuery {
	for _, item := range items {
		q.FromList = append(q.FromList, item)
	}
	return q
}

func (q *UpdateQuery) Fromf(format string, a ...interface{}) *UpdateQuery {
	q.FromList = append(q.FromList, NewFromf(format, a...)) //
	return q
}
