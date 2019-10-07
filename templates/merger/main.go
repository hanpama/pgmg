package main

import (
	"bytes"
	"encoding/base64"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var tmpl = template.Must(template.New("fstr").Parse(`// Code generated by templates/merger/main.go. DO NOT EDIT.
package templates

import "encoding/base64"

{{.Comment}}
var content = mustDecode(
	{{- range $i, $line := .Base64Slice }}{{ if $i }} +{{end}}
		"{{$line}}"
	{{- end -}},
)

func mustDecode(b64 string) []byte {
  dat, err := base64.StdEncoding.DecodeString(b64)
  if err != nil {
    panic(err)
  }
  return dat
}
`))

const glob = "templates/*.go.tmpl"
const out = "templates/merged_gen.go"

func main() {
	files, err := filepath.Glob(glob)
	if err != nil {
		panic(err)
	}
	merged := []byte{}
	for _, filename := range files {
		b, err := ioutil.ReadFile(filename)

		if err != nil {
			panic(err)
		}
		merged = append(merged, b...)
		merged = append(merged, []byte("\n")...)
	}

	var r bytes.Buffer
	enc := base64.NewEncoder(base64.StdEncoding, &r)

	_, err = enc.Write(merged)
	if err != nil {
		panic(err)
	}

	var args tmplArgs

	args.Comment = "// " + strings.ReplaceAll(string(merged), "\n", "\n// ")
	b := make([]byte, 100)
	for {
		n, err := r.Read(b)
		if err == io.EOF {
			break
		}
		args.Base64Slice = append(args.Base64Slice, string(b[:n]))
	}

	f, err := os.OpenFile(out, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	if err := tmpl.ExecuteTemplate(f, "fstr", args); err != nil {
		panic(err)
	}
}

type tmplArgs struct {
	Base64Slice []string
	Comment     string
}
