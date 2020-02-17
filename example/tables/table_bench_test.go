package tables_test

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	"github.com/hanpama/pgmg/example/tables"
)

func BenchmarkGetByID(b *testing.B) {
	ctx := context.Background()
	fxt, err := prepareTestEnv(ctx, "..")
	if err != nil {
		b.Fatal(err)
	}
	products := tables.NewProductTable(fxt.tdb)
	if err = fxt.tdb.Prepare(ctx, tables.SQLGetProductRowsByID); err != nil {
		b.Fatal(err)
	}
	for _, rows := range []int{1, 2, 5, 20, 100} {
		b.Run(fmt.Sprintf("%d Row", rows), func(b *testing.B) {
			keys := selectKeys(fxt, rows)
			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_, err := products.GetByID(ctx, keys...)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func BenchmarkGetByIDBaseline(b *testing.B) {
	ctx := context.Background()
	fxt, err := prepareTestEnv(ctx, "..")
	if err != nil {
		b.Fatal(err)
	}
	query := "SELECT * FROM wise.product WHERE id = any($1::int[])"
	if err = fxt.tdb.Prepare(ctx, query); err != nil {
		b.Fatal(err)
	}

	for _, rows := range []int{1, 2, 5, 20, 100} {
		b.Run(fmt.Sprintf("%d Row", rows), func(b *testing.B) {
			keys := selectKeys(fxt, rows)
			var buff bytes.Buffer
			buff.WriteRune('{')
			for i, k := range keys {
				if i > 0 {
					buff.WriteRune(',')
				}
				buff.WriteString(fmt.Sprintf("%d", k.(tables.ProductID).ID))
			}
			buff.WriteRune('}')
			arg1 := buff.String()

			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				var datas []*tables.ProductData
				rows, err := fxt.tdb.DB.QueryContext(ctx, query, arg1)
				if err != nil {
					b.Fatal(err)
				}
				for rows.Next() {
					var d tables.ProductData
					if err = rows.Scan(&d.ID, &d.Price, &d.Name, &d.Alias, &d.Stocked, &d.Sold); err != nil {
						b.Fatal(err)
					}
					datas = append(datas, &d)
				}
				rows.Close()
			}
		})
	}
}

func BenchmarkFind(b *testing.B) {
	ctx := context.Background()
	fxt, err := prepareTestEnv(ctx, "..")
	if err != nil {
		b.Fatal(err)
	}
	products := tables.NewProductTable(fxt.tdb)
	if err = fxt.tdb.Prepare(ctx, tables.SQLFindProductRows); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		rows, err := products.Find(ctx, tables.ProductValues{})
		if err != nil {
			b.Fatal(err)
		}
		if len(rows) != 100 {
			b.FailNow()
		}
	}
}

func BenchmarkFindAndScan(b *testing.B) {
	ctx := context.Background()
	fxt, err := prepareTestEnv(ctx, "..")
	if err != nil {
		b.Fatal(err)
	}
	if err = fxt.tdb.Prepare(ctx, "SELECT * FROM wise.product"); err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var rows tables.ProductRows
		_, err := fxt.tdb.QueryAndReceive(ctx, rows.ReceiveRows, "SELECT * FROM wise.product")
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFindBase(b *testing.B) {
	ctx := context.Background()
	fxt, err := prepareTestEnv(ctx, "..")
	if err != nil {
		b.Fatal(err)
	}
	if err = fxt.tdb.Prepare(ctx, "SELECT * FROM wise.product"); err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		rows, err := fxt.tdb.DB.QueryContext(ctx, "SELECT * FROM wise.product")
		if err != nil {
			b.Fatal(err)
		}
		var datas []*tables.ProductData
		for rows.Next() {
			var d tables.ProductData
			if err = rows.Scan(&d.ID, &d.Price, &d.Name, &d.Alias, &d.Stocked, &d.Sold); err != nil {
				b.Fatal(err)
			}
			datas = append(datas, &d)
		}
		rows.Close()
	}
}

func selectKeys(fxt *testEnv, num int) (keys []interface{}) {
	for i := 0; i < num; i++ {
		keys = append(keys, fxt.productRows[i].KeyID())
	}
	return keys
}
