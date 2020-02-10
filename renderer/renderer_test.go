package renderer_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/hanpama/pgmg/testutil"

	"github.com/hanpama/pgmg/introspect"
	"github.com/hanpama/pgmg/model"
	"github.com/hanpama/pgmg/renderer"
)

func TestRenderDataSource(t *testing.T) {
	var err error
	var b []byte

	if b, err = ioutil.ReadFile("../example/introspection.json"); err != nil {
		t.Fatal(err)
	}
	ins := new(introspect.Schema)
	if err = json.Unmarshal(b, ins); err != nil {
		t.Fatal(err)
	}

	schema := model.NewSchemaFromIntrospection(ins)

	if b, err = renderer.RenderSchema("schema", schema); err != nil {
		t.Fatal(err)
	}

	if err = testutil.TestSnapshot(b, "../example/schema/pgmg.go"); err != nil {
		t.Fatal(err)
	}
}
