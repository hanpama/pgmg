package model

import "github.com/knq/snaker"

func formatCapitalName(name string) string {
	return snaker.ForceCamelIdentifier(name)
}

func formatLowerName(name string) string {
	name = snaker.ForceLowerCamelIdentifier(name)
	if _, ok := reservedWords[name]; ok {
		return "_" + name
	}
	return name
}

var reservedWords = map[string]bool{
	"break": true, "default": true, "func": true, "interface": true, "select": true,
	"case": true, "defer": true, "go": true, "map": true, "struct": true,
	"chan": true, "else": true, "goto": true, "package": true, "switch": true,
	"const": true, "fallthrough": true, "if": true, "range": true, "type": true,
	"continue": true, "for": true, "import": true, "return": true, "var": true,
}
