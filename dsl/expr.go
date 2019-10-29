package dsl

import (
	"fmt"
	"strings"
)

type Expression interface {
	Expr() string
}

func NewExprf(format string, a ...interface{}) *Exprf {
	return &Exprf{format, WrapExprs(a...)}
}

var Null = NewExprf("NULL")

type Exprf struct {
	Format      string
	Expressions []Expression
}

func (c *Exprf) Expr() string {
	strs := make([]interface{}, len(c.Expressions))
	for i, val := range c.Expressions {
		strs[i] = val.Expr()
	}
	return fmt.Sprintf(c.Format, strs...)
}

type AnyExpr struct{ val interface{} }

func (s AnyExpr) Expr() string { return fmt.Sprintf("%v", s.val) }

func WrapExpr(val interface{}) Expression {
	exp, ok := val.(Expression)
	if ok {
		return exp
	}
	return AnyExpr{val}
}

func WrapExprs(vals ...interface{}) (res []Expression) {
	res = make([]Expression, len(vals))
	for i, val := range vals {
		res[i] = WrapExpr(val)
	}
	return res
}

type FuncExpr struct {
	Name        string
	Expressions []Expression
}

func (c *FuncExpr) Expr() string {
	var b strings.Builder
	b.WriteString(c.Name)
	b.WriteString("(")
	for i, e := range c.Expressions {
		if i != 0 {
			b.WriteString(", ")
		}
		b.WriteString(e.Expr())
	}
	b.WriteString(")")
	return b.String()
}

type SelectQueryAlias struct {
	SelectQuery *SelectQuery
	Alias       string
}

func (a SelectQueryAlias) Expr() string {
	return fmt.Sprintf(`%s AS "%s"`, a.SelectQuery.Expr(), a.Alias)
}
func (a SelectQueryAlias) From() string {
	return fmt.Sprintf(`%s AS "%s"`, a.SelectQuery.Expr(), a.Alias)
}
