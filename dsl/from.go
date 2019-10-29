package dsl

import "fmt"

type FromItem interface {
	From() string
}

func NewFromf(format string, a ...interface{}) *Fromf {
	return &Fromf{format, WrapExprs(a...)}
}

type Fromf struct {
	Format      string
	Expressions []Expression
}

func (f Fromf) From() string {
	strs := make([]interface{}, len(f.Expressions))
	for i, val := range f.Expressions {
		strs[i] = val.Expr()
	}
	return fmt.Sprintf(f.Format, strs...)

}
