package dsl

import (
	"bytes"
	"fmt"
	"text/template"
)

// Select creates a new SelectQuery. If no selectors are given, it uses asterisk(*) as its selection set.
func Select(selectors ...interface{}) *SelectQuery {
	q := &SelectQuery{}
	q.Select = WrapExprs(selectors...)
	return q
}

func SelectDistinct() {}

type SelectQuery struct {
	CTE         []CTE
	Conditions  []Expression
	FromList    []FromItem
	Select      []Expression
	SortOptions []Sortable
	Count       Expression
	Start       Expression
	Locking     *Locking
}

var selectQueryTmpl = template.Must(template.New("Select").Parse(`SELECT {{if .Select}}{{range $i, $c := .Select -}}
{{if $i}}, {{end}}{{$c.Expr}}{{ end -}}{{else}}*{{end -}}
{{range $i, $f := .FromList -}}{{if eq $i 0}} FROM {{else}}, {{end}}{{$f.From}}{{ end -}}
{{range $i, $c := .Conditions}}{{ if eq $i 0 }} WHERE {{else}} AND{{end}}{{$c.Expr}}{{end -}}
{{range $i, $s := .SortOptions}}{{ if eq $i 0 }} ORDER BY {{else}}, {{end -}}
{{$s.Expr}}{{if $s.Order}} {{$s.Order}}{{end}}{{if $s.Nulls}} NULLS {{$s.Nulls}}{{end}}
{{- end -}}
{{if .Count }} LIMIT {{.Count.Expr}}{{end -}}
{{if .Start }} OFFSET {{.Start.Expr}}{{end -}}
{{if .Locking}} FOR {{.Locking.Strength -}}
	{{range $i, $t := .Locking.TableNames}}{{if eq $i 0}} OF{{else}}, {{end}}{{$t}}{{end -}}
	{{if .Locking.Concurrency}} {{.Locking.Concurrency}}{{end -}}
{{end -}}
`))

func (q *SelectQuery) Statement() string {
	var buff bytes.Buffer
	err := selectQueryTmpl.Execute(&buff, q)
	if err != nil {
		panic(err)
	}
	return string(buff.Bytes())
}

func (q *SelectQuery) As(alias string) SelectQueryAlias {
	return SelectQueryAlias{q, alias}
}

func (q *SelectQuery) Expr() string {
	return fmt.Sprintf(`(%s)`, q.Statement())
}

func (q *SelectQuery) From(items ...FromItem) *SelectQuery {
	for _, item := range items {
		q.FromList = append(q.FromList, item)
	}
	return q
}

func (q *SelectQuery) Fromf(format string, a ...interface{}) *SelectQuery {
	q.FromList = append(q.FromList, NewFromf(format, a...)) //
	return q
}

func (q *SelectQuery) Where(conds ...Expression) *SelectQuery {
	q.Conditions = append(q.Conditions, conds...)
	return q
}

func (q *SelectQuery) Wheref(format string, a ...interface{}) *SelectQuery {
	q.Conditions = append(q.Conditions, NewExprf(format, a...))
	return q
}

func (q *SelectQuery) OrderBy(options ...Sortable) *SelectQuery {
	q.SortOptions = append(q.SortOptions, options...)
	return q
}

func (q *SelectQuery) Limit(count interface{}) *SelectQuery {
	q.Count = WrapExpr(count)
	return q
}
func (q *SelectQuery) Limitf(format string, a ...interface{}) *SelectQuery {
	q.Count = NewExprf(format, a...)
	return q
}
func (q *SelectQuery) Offset(start interface{}) *SelectQuery {
	q.Start = WrapExpr(start)
	return q
}
func (q *SelectQuery) Offsetf(format string, a ...interface{}) *SelectQuery {
	q.Start = NewExprf(format, a...)
	return q
}
func (q *SelectQuery) For() *SelectQuery {
	panic("Unimeplemented")
}

type Locking struct {
	Strength    string
	TableNames  []string
	Concurrency string
}

type Sortable interface {
	Expr() string
	Order() string // Order [ ASC | DESC | USING operator]
	Nulls() string // Null can be zero | "FIRST" | "LAST"
}

type ConflictAction struct {
	Setters    []Setter
	Conditions string
}
