// Code generated by scripts/fstr/gen.go. DO NOT EDIT.
package templates

import (
	"fmt"
	"text/template"
)

var funcs = template.FuncMap{
	"sqlParam":   func(n int) string { return fmt.Sprintf("$%d", n+1) },
	"goQueryArg": func(n int) string { return fmt.Sprintf("a%d", n+1) },
}

var Tmpl = template.Must(template.New("pgmg").Funcs(funcs).Parse(string(content)))
