package dsl

import (
	"fmt"
)

type Table interface {
	AsTable() string
}
type Column interface {
	AsColumn() string
}

type TableReference struct {
	TableSchema string
	TableName   string
	Alias       string
}

func (tr TableReference) Expr() string {
	if tr.Alias == "" {
		return fmt.Sprintf(`"%s"."%s"`, tr.TableSchema, tr.TableName)
	}
	return fmt.Sprintf(`"%s"`, tr.Alias)
}
func (tr TableReference) From() string {
	if tr.Alias == "" {
		return fmt.Sprintf(`"%s"."%s"`, tr.TableSchema, tr.TableName)
	}
	return fmt.Sprintf(`"%s"."%s" AS "%s"`, tr.TableSchema, tr.TableName, tr.Alias)
}
func (tr TableReference) AsTable() string {
	return tr.From()
}

type ColumnReference struct {
	TableReference TableReference
	ColumnName     string
}

func (cr ColumnReference) Expr() string {
	return fmt.Sprintf(`%s."%s"`, cr.TableReference.Expr(), cr.ColumnName)
}
func (cr ColumnReference) Set(val interface{}) ColumnSetter {
	return ColumnSetter{cr, WrapExpr(val)}
}
func (cr ColumnReference) AsColumn() string { return fmt.Sprintf(`"%s"`, cr.ColumnName) }
func (cr ColumnReference) Order() string    { return "" }
func (cr ColumnReference) Nulls() string    { return "" }

type ColumnSetter struct {
	column ColumnReference
	value  Expression
}

func (cs ColumnSetter) AsSetter() string {
	return fmt.Sprintf("%s = %s", cs.column.AsColumn(), cs.value.Expr())
}
