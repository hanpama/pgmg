package tables_test

import (
	"context"
	"fmt"
	"testing"
	"time"

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
	selectKeys := func(num int) (keys []interface{}) {
		for i := 0; i < num; i++ {
			keys = append(keys, fxt.productRows[i].KeyID())
		}
		return keys
	}
	for _, rows := range []int{1, 2, 5, 20, 100} {
		b.Run(fmt.Sprintf("%d Row", rows), func(b *testing.B) {
			keys := selectKeys(rows)
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
		var (
			id      int
			price   string
			name    string
			alias   string
			stocked time.Time
			sold    *time.Time
		)
		for rows.Next() {
			if err = rows.Scan(&id, &price, &name, &alias, &stocked, &sold); err != nil {
				b.Fatal(err)
			}
		}
	}
}
