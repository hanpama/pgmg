package dsl

import (
	"bytes"
	"text/template"
)

// JoinClause [ NATURAL ] join_type from_item [ ON join_condition | USING ( join_column [, ...] ) ]
type JoinClause struct {
	Left    FromItem
	Right   FromItem
	Natural bool
	Type    string
	Conds   []Expression
	Outer   bool
}

var joinTmpl = template.Must(template.New("Join").Parse(`
{{- .Left.From}}
{{- if .Natural}} NATURAL{{end -}}
{{- if .Type}} {{.Type}}{{end -}}
{{- if .Outer}} OUTER{{end}} JOIN {{.Right.From }}
{{- range $i, $c := .Conds -}}
{{- if eq $i 0 }} ON {{else}}, {{end}}{{$c.Expr}}{{end}}`))

func (jc *JoinClause) From() string {
	var buff bytes.Buffer
	err := joinTmpl.Execute(&buff, jc)
	if err != nil {
		panic(err)
	}
	return string(buff.Bytes())
}

func (jc *JoinClause) On(exprs ...Expression) *JoinClause {
	jc.Conds = append(jc.Conds, exprs...)
	return jc
}

func (jc *JoinClause) Onf(format string, a ...interface{}) *JoinClause {
	jc.Conds = append(jc.Conds, NewExprf(format, a...))
	return jc
}

func (jc *JoinClause) Join(t FromItem) *JoinClause {
	return &JoinClause{Left: jc, Right: t}
}
func (jc *JoinClause) InnerJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: jc, Right: t, Type: "INNER"}
}
func (jc *JoinClause) LeftJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: jc, Right: t, Type: "LEFT"}
}
func (jc *JoinClause) RightJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: jc, Right: t, Type: "RIGHT"}
}
func (jc *JoinClause) FullJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: jc, Right: t, Type: "FULL"}
}
func (jc *JoinClause) LeftOuterJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: jc, Right: t, Type: "LEFT", Outer: true}
}
func (jc *JoinClause) RightOuterJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: jc, Right: t, Type: "RIGHT", Outer: true}
}
func (jc *JoinClause) FullOuterJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: jc, Right: t, Type: "FULL", Outer: true}
}
func (jc *JoinClause) CrossJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: jc, Right: t, Type: "CROSS"}
}
func (jc *JoinClause) NaturalJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: jc, Right: t, Natural: true}
}
func (jc *JoinClause) NaturalInnerJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: jc, Right: t, Type: "INNER", Natural: true}
}
func (jc *JoinClause) NaturalLeftJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: jc, Right: t, Type: "LEFT", Natural: true}
}
func (jc *JoinClause) NaturalRightJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: jc, Right: t, Type: "RIGHT", Natural: true}
}
func (jc *JoinClause) NaturalFullJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: jc, Right: t, Type: "FULL", Natural: true}
}
func (jc *JoinClause) NaturalLeftOuterJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: jc, Right: t, Type: "LEFT", Outer: true, Natural: true}
}
func (jc *JoinClause) NaturalRightOuterJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: jc, Right: t, Type: "RIGHT", Outer: true, Natural: true}
}
func (jc *JoinClause) NaturalFullOuterJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: jc, Right: t, Type: "FULL", Outer: true, Natural: true}
}
func (jc *JoinClause) NaturalCrossJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: jc, Right: t, Type: "CROSS", Natural: true}
}

func (tr TableReference) Join(t FromItem) *JoinClause {
	return &JoinClause{Left: tr, Right: t}
}
func (tr TableReference) InnerJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: tr, Right: t, Type: "INNER"}
}
func (tr TableReference) LeftJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: tr, Right: t, Type: "LEFT"}
}
func (tr TableReference) RightJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: tr, Right: t, Type: "RIGHT"}
}
func (tr TableReference) FullJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: tr, Right: t, Type: "FULL"}
}
func (tr TableReference) LeftOuterJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: tr, Right: t, Type: "LEFT", Outer: true}
}
func (tr TableReference) RightOuterJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: tr, Right: t, Type: "RIGHT", Outer: true}
}
func (tr TableReference) FullOuterJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: tr, Right: t, Type: "FULL", Outer: true}
}
func (tr TableReference) CrossJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: tr, Right: t, Type: "CROSS"}
}
func (tr TableReference) NaturalJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: tr, Right: t, Natural: true}
}
func (tr TableReference) NaturalInnerJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: tr, Right: t, Type: "INNER", Natural: true}
}
func (tr TableReference) NaturalLeftJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: tr, Right: t, Type: "LEFT", Natural: true}
}
func (tr TableReference) NaturalRightJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: tr, Right: t, Type: "RIGHT", Natural: true}
}
func (tr TableReference) NaturalFullJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: tr, Right: t, Type: "FULL", Natural: true}
}
func (tr TableReference) NaturalLeftOuterJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: tr, Right: t, Type: "LEFT", Outer: true, Natural: true}
}
func (tr TableReference) NaturalRightOuterJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: tr, Right: t, Type: "RIGHT", Outer: true, Natural: true}
}
func (tr TableReference) NaturalFullOuterJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: tr, Right: t, Type: "FULL", Outer: true, Natural: true}
}
func (tr TableReference) NaturalCrossJoin(t FromItem) *JoinClause {
	return &JoinClause{Left: tr, Right: t, Type: "CROSS", Natural: true}
}

type joiner interface {
	Join(t FromItem) *JoinClause
	InnerJoin(t FromItem) *JoinClause
	LeftJoin(t FromItem) *JoinClause
	RightJoin(t FromItem) *JoinClause
	FullJoin(t FromItem) *JoinClause
	LeftOuterJoin(t FromItem) *JoinClause
	RightOuterJoin(t FromItem) *JoinClause
	FullOuterJoin(t FromItem) *JoinClause
	CrossJoin(t FromItem) *JoinClause
	NaturalJoin(t FromItem) *JoinClause
	NaturalInnerJoin(t FromItem) *JoinClause
	NaturalLeftJoin(t FromItem) *JoinClause
	NaturalRightJoin(t FromItem) *JoinClause
	NaturalFullJoin(t FromItem) *JoinClause
	NaturalLeftOuterJoin(t FromItem) *JoinClause
	NaturalRightOuterJoin(t FromItem) *JoinClause
	NaturalFullOuterJoin(t FromItem) *JoinClause
	NaturalCrossJoin(t FromItem) *JoinClause
}

var _ joiner = TableReference{}
var _ joiner = &JoinClause{}
