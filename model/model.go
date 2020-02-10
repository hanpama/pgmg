package model

import (
	"strings"
)

type Schema struct {
	Name   string
	Tables []*Table
}

func (s *Schema) Dependencies() map[string]bool {
	deps := make(map[string]bool)
	for _, t := range s.Tables {
		for module := range t.Dependencies() {
			deps[module] = true
		}
	}
	return deps
}

type Table struct {
	schema       *Schema
	SQLName      string
	Columns      []*Column
	Keys         []*Key
	ForeignKeys  []*ForeignKey
	IsInsertable bool
}

func (m *Table) Schema() string      { return m.schema.Name }
func (m *Table) CapitalName() string { return formatCapitalName(m.SQLName) }
func (m *Table) LowerName() string   { return formatLowerName(m.SQLName) }
func (m *Table) Dependencies() map[string]bool {
	deps := make(map[string]bool)
	for _, c := range m.Columns {
		if c.Type.module != "" {
			deps[c.Type.module] = true
		}
	}
	return deps
}
func (m *Table) UpdatableColumns() (ucs []*Column) {
	for _, c := range m.Columns {
		if c.SQLUpdatable {
			ucs = append(ucs, c)
		}
	}
	return ucs
}

func (m *Table) IsUpdatable() bool {
	return len(m.UpdatableColumns()) > 0
}

func (m *Table) PrimaryKey() *Key {
	for _, k := range m.Keys {
		if k.IsPrimaryKey {
			return k
		}
	}
	if len(m.Keys) > 0 {
		return m.Keys[0] // ?
	}
	return nil
}

type Column struct {
	SQLName      string
	SQLNullable  bool
	SQLDefault   string
	SQLUpdatable bool
	Type         Type
}

func (p *Column) CapitalName() string { return formatCapitalName(p.SQLName) }
func (p *Column) LowerName() string   { return formatLowerName(p.SQLName) }

func (p *Column) Nullable() bool   { return p.SQLDefault == "" && p.SQLNullable }
func (p *Column) HasDefault() bool { return p.SQLDefault != "" }

type Key struct {
	Table        *Table
	SQLName      string
	IsPrimaryKey bool
	Columns      []*Column
}

func (k *Key) IsComposite() bool {
	return len(k.Columns) > 1
}

func (k *Key) CapitalName() string {
	colNames := make([]string, len(k.Columns))
	for i, c := range k.Columns {
		colNames[i] = c.CapitalName()
	}
	return strings.Join(colNames, "")
}

func (k *Key) TypeName() string {
	return k.Table.CapitalName() + k.CapitalName()
}

func (k *Key) LowerName() string {
	return formatLowerName(k.CapitalName())
}

func (k *Key) Column() *Column {
	return k.Columns[0]
}

type ForeignKey struct {
	SQLName   string
	Columns   []*Column
	TargetKey *Key
}

func (fk *ForeignKey) TypeName() string {
	return fk.TargetKey.TypeName()
}

func (fk *ForeignKey) CapitalName() string {
	colNames := make([]string, len(fk.Columns))
	for i, c := range fk.Columns {
		colNames[i] = c.CapitalName()
	}
	return strings.Join(colNames, "")
}

type Type struct {
	BaseType   string
	SelectType string
	InsertType string
	FilterType string
	SQLType    string
	module     string
}
