package internal

import (
	"bytes"
	"go/format"

	"github.com/hanpama/pgmg/templates"

	"github.com/knq/snaker"
)

// var funcs = template.FuncMap{
// 	"sqlParam": func(n int) string { return fmt.Sprintf("$%d", n+1) },
// }
// var modelTmpl = template.Must(
// 	template.New("model.go.tmpl").Funcs(funcs).ParseFiles("templates/model.go.tmpl"),
// )
// var queryTmpl = template.Must(
// 	template.New("query.go.tmpl").Funcs(funcs).ParseFiles("templates/query.go.tmpl"),
// )

func RenderModel(t *Table) ([]byte, error) {
	var buff bytes.Buffer
	err := templates.Tmpl.ExecuteTemplate(&buff, "model", &model{t})
	if err != nil {
		return nil, err
	}
	return format.Source(buff.Bytes())
}

func RenderQuery(t *Table) ([]byte, error) {
	var buff bytes.Buffer
	err := templates.Tmpl.ExecuteTemplate(&buff, "query", &model{t})
	if err != nil {
		return nil, err
	}
	return format.Source(buff.Bytes())
}

type model struct {
	t *Table
}

func (m *model) Dependencies() (mods []string) {
	for _, col := range m.t.Columns {
		if mod := pgTypeToGoType(col.DataType).Module; mod != "" {
			mods = append(mods, mod)
		}
	}
	return mods
}

func (m *model) Name() string { return m.t.Name }

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
func (p *property) LowerName() string   { return snaker.ForceLowerCamelIdentifier(p.c.Name) }
func (p *property) SQLName() string     { return p.c.Name }
func (p *property) BaseType() string    { return p.t.Name }
func (p *property) SelectType() string {
	if p.c.IsNullable {
		return p.t.NullableName
	}
	return p.t.Name
}
func (p *property) SelectNullable() bool { return p.c.IsNullable }
func (p *property) NullableType() string { return p.t.NullableName }

type key struct{ k *Key }

func (k *key) CapitalName() string { return snaker.ForceCamelIdentifier(k.k.Name) }

func (k *key) Properties() (props []property) {
	for i := range k.k.Columns {
		props = append(props, property{&k.k.Columns[i], pgTypeToGoType(k.k.Columns[i].DataType)})
	}
	return props
}
