package main

import (
	"bytes"
	"go/format"

	"github.com/knq/snaker"
)

func RenderTableModel(packageName string, ts []Table) ([]byte, error) {
	tmplArgs := struct {
		PackageName  string
		Dependencies map[string]bool
		Models       []*model
	}{
		packageName,
		make(map[string]bool),
		make([]*model, len(ts)),
	}

	for i, t := range ts {
		for _, col := range t.Columns {
			if mod := pgTypeToGoType(col.DataType).Module; mod != "" {
				tmplArgs.Dependencies[mod] = true
			}
		}
		tmplArgs.Models[i] = &model{t}
	}

	var buff bytes.Buffer
	err := Tmpl.ExecuteTemplate(&buff, "table_model", tmplArgs)
	if err != nil {
		return nil, err
	}
	return format.Source(buff.Bytes())
}

type schema struct {
	Dependencies map[string]bool
	Models       []*model
}

type model struct {
	t Table
}

func (m *model) CapitalName() string { return snaker.ForceCamelIdentifier(m.t.Name) }
func (m *model) SQLName() string     { return m.t.Name }
func (m *model) Schema() string      { return m.t.Schema }

func (m *model) Properties() (props []property) {
	for i := range m.t.Columns {
		props = append(props, property{&m.t.Columns[i], pgTypeToGoType(m.t.Columns[i].DataType)})
	}
	return props
}

func (m *model) Keys() (keys []key) {
	for i := range m.t.Keys {
		keys = append(keys, key{&m.t.Keys[i]})
	}
	return keys
}

type property struct {
	c *Column
	t typeMapping
}

func (p *property) CapitalName() string { return snaker.ForceCamelIdentifier(p.c.Name) }
func (p *property) LowerName() string {
	name := snaker.ForceLowerCamelIdentifier(p.c.Name)
	if _, ok := reservedWords[name]; ok {
		return "_" + name
	}
	return name
}
func (p *property) SQLName() string { return p.c.Name }
func (p *property) GoBaseType() string {
	return p.t.Name
}
func (p *property) GoSelectType() string {
	if p.c.IsNullable {
		return p.t.NullableName
	}
	return p.t.Name
}
func (p *property) CanBeNull() bool { return p.c.IsNullable || p.c.Default != "" }
func (p *property) GoInsertType() string {
	if p.c.IsNullable || p.c.Default != "" {
		return p.t.NullableName
	}
	return p.t.Name
}
func (p *property) FilterType() string { return p.t.NullableName }
func (p *property) SQLType() string    { return p.t.PGType }
func (p *property) Default() string    { return p.c.Default }

type key struct{ k *Key }

func (k *key) CapitalName() string { return snaker.ForceCamelIdentifier(k.k.Name) }
func (k *key) SQLName() string     { return k.k.Name }

func (k *key) Properties() (props []property) {
	for i := range k.k.Columns {
		props = append(props, property{&k.k.Columns[i], pgTypeToGoType(k.k.Columns[i].DataType)})
	}
	return props
}

var reservedWords = map[string]bool{
	"break": true, "default": true, "func": true, "interface": true, "select": true,
	"case": true, "defer": true, "go": true, "map": true, "struct": true,
	"chan": true, "else": true, "goto": true, "package": true, "switch": true,
	"const": true, "fallthrough": true, "if": true, "range": true, "type": true,
	"continue": true, "for": true, "import": true, "return": true, "var": true,
}
