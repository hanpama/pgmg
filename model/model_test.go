package model_test

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/hanpama/pgmg/model"

	"github.com/hanpama/pgmg/introspect"
)

func TestNewSchemaFromIntrospection(t *testing.T) {
	b, err := ioutil.ReadFile("../example/introspection.json")
	if err != nil {
		t.Fatal(err)
	}

	ins := new(introspect.Schema)
	if err = json.Unmarshal(b, ins); err != nil {
		t.Fatal(err)
	}

	schema := model.NewSchemaFromIntrospection(ins)

	var packageProductTable, packageTable *model.Table
	for _, t := range schema.Tables {
		if t.SQLName == "package_product" {
			packageProductTable = t
		}
		if t.SQLName == "package" {
			packageTable = t
		}
	}

	for i, test := range []struct {
		Got      interface{}
		Expected interface{}
	}{
		{schema.Name, "wise"},
		{len(schema.Tables), 6},
		{len(packageProductTable.ForeignKeys), 2},
		{len(packageProductTable.ForeignKeys[0].Columns), 1},
		{len(packageProductTable.ForeignKeys[1].Columns), 1},
		{packageProductTable.ForeignKeys[0].TargetKey, packageTable.Keys[0]},
	} {
		if !reflect.DeepEqual(test.Got, test.Expected) {
			t.Fatalf("Assertion failed at index %d: expected %#v, but got %#v", i, test.Expected, test.Got)
		}
	}

}
