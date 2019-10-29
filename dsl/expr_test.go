package dsl_test

import (
	"testing"

	"github.com/hanpama/pgmg/dsl"
)

var exprTestCases = testCases{
	{
		name:      "Coalesce",
		test:      func() string { return dsl.Coalesce(dsl.Null, 2).Expr() },
		expecting: `COALESCE(NULL, 2)`,
	},
	{
		name:      "Column",
		test:      func() string { return User.ID.Expr() },
		expecting: `"public"."user"."id"`,
	},
}

func TestExpr(t *testing.T) {
	for _, tc := range exprTestCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.test() != tc.expecting {
				t.Errorf("Exptected %s \nbut got %s", tc.expecting, tc.test())
			}
		})
	}
}

func BenchmarkExpr(b *testing.B) {
	for _, tc := range exprTestCases {
		b.Run(tc.name, func(b *testing.B) {
			if tc.test() != tc.expecting {
				b.Errorf("Exptected %s \nbut got %s", tc.expecting, tc.test())
			}
		})
	}
}
