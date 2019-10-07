package internal

import (
	"bytes"
	"go/format"
	"strings"

	"github.com/hanpama/pgmg/templates"
)

func RenderQuery(qi *Query) ([]byte, error) {
	var buff bytes.Buffer
	err := templates.Tmpl.ExecuteTemplate(&buff, "query", &qi)
	if err != nil {
		return nil, err
	}
	return format.Source(buff.Bytes())
}

type Query struct {
	Name       string
	SQL        string
	ParamTypes []string
	Returning  []Column
}

func (qi *Query) Dependencies() (mods []string) {
	for _, col := range qi.Returning {
		if mod := pgTypeToGoType(col.DataType).Module; mod != "" {
			mods = append(mods, mod)
		}
	}
	return mods
}

func (qi *Query) QueryComment() string {
	res := "// Query queries"
	for _, line := range strings.Split(qi.SQL, "\n") {
		res = "//  " + line
	}
	return res
}

func (qi *Query) Properties() (props []property) {
	for i := range qi.Returning {
		props = append(props, property{&qi.Returning[i], pgTypeToGoType(qi.Returning[i].DataType)})
	}
	return props
}

func (qi *Query) GoParamTypes() []string {
	types := make([]string, len(qi.ParamTypes))
	for i, t := range qi.ParamTypes {
		types[i] = pgTypeToGoType(t).NullableName
	}
	return types
}
