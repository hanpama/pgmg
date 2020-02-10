package tables_test

import (
	"context"
	"testing"
	"time"

	"github.com/hanpama/pgmg/example"

	"github.com/hanpama/pgmg/example/tables"
	_ "github.com/lib/pq"
)

func TestGetByID(t *testing.T) {
	ctx := context.Background()
	tdb := example.NewTestDB(ctx, t, "..")
	fxt := prepareTestFixture(ctx, t, tdb)
	packages := tables.NewPackageTable(tdb)

	// Get
	rows, err := packages.GetByID(ctx, fxt.packageRows[0].KeyID(), fxt.packageRows[1].KeyID())
	if err != nil {
		t.Fatal(err)
	}
loop:
	for _, r := range rows {
		for i := range fxt.packageRows {
			if r.GetID() == fxt.packageRows[i].GetID() {
				if r.GetName() != fxt.packageRows[i].GetName() {
					t.Fatalf("Unexpected name: %s", r.GetName())
				}
				if !r.GetAvailable() {
					t.Fatalf("Unexpected available: false")
				}
				continue loop
			}
		}
		t.Fatalf("Unexpected id: %s", r.GetID())
	}
}

func TestFind(t *testing.T) {
	ctx := context.Background()
	tdb := example.NewTestDB(ctx, t, "..")
	prepareTestFixture(ctx, t, tdb)
	packages := tables.NewPackageTable(tdb)

	// Find
	queryName := "Package for baz"
	filter := tables.PackageValues{
		Name: &queryName,
	}
	rows, err := packages.Find(ctx, filter)

	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 2 {
		t.Fatalf("Unexpected count: %d", len(rows))
	}
}

func TestCount(t *testing.T) {
	ctx := context.Background()
	tdb := example.NewTestDB(ctx, t, "..")
	prepareTestFixture(ctx, t, tdb)
	packages := tables.NewPackageTable(tdb)

	queryName := "Package for baz"
	filter := tables.PackageValues{
		Name: &queryName,
	}
	// Count
	count, err := packages.Count(ctx, filter)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Fatalf("Unexpected count: %d", count)
	}
}

func TestUpdate(t *testing.T) {
	ctx := context.Background()
	tdb := example.NewTestDB(ctx, t, "..")
	prepareTestFixture(ctx, t, tdb)
	packages := tables.NewPackageTable(tdb)

	filterName := "Package for baz"
	newName := "Package for foo"

	filter := tables.PackageValues{Name: &filterName}
	changeset := tables.PackageValues{Name: &newName}

	affected, err := packages.Update(ctx, changeset, filter)
	if err != nil {
		t.Fatal(err)
	}
	if affected != 2 {
		t.Fatalf("Unexpected count: %d", affected)
	}

}

func TestSave(t *testing.T) {
	ctx := context.Background()
	tdb := example.NewTestDB(ctx, t, "..")
	fxt := prepareTestFixture(ctx, t, tdb)
	packages := tables.NewPackageTable(tdb)

	// Save
	fxt.packageRows[1].SetName("Package for foo")
	if err := packages.Save(ctx, fxt.packageRows...); err != nil {
		t.Fatal(err)
	}
}

func TestUpdateBy(t *testing.T) {
	ctx := context.Background()
	tdb := example.NewTestDB(ctx, t, "..")
	fxt := prepareTestFixture(ctx, t, tdb)
	packages := tables.NewPackageTable(tdb)

	newName := "Package for bar"
	changeset := tables.PackageValues{Name: &newName}
	affected, err := packages.UpdateByID(ctx, changeset, fxt.packageRows[0].KeyID(), fxt.packageRows[1].KeyID())
	if err != nil {
		t.Fatal(err)
	}
	if affected != 2 {
		t.Fatalf("Unexpected count: %d", affected)
	}

}

func TestDelete(t *testing.T) {
	ctx := context.Background()
	tdb := example.NewTestDB(ctx, t, "..")
	prepareTestFixture(ctx, t, tdb)
	packages := tables.NewPackageTable(tdb)

	newName := "Package for baz"
	filter := tables.PackageValues{Name: &newName}

	affected, err := packages.Delete(ctx, filter)
	if err != nil {
		t.Fatal(err)
	}
	if affected != 2 {
		t.Fatalf("Unexpected count: %d", affected)
	}

	result, err := packages.Find(ctx, tables.PackageValues{})
	if err != nil {
		t.Fatal(err)
	}
	if len(result) != 2 {
		t.Fatal()
	}
}

func TestDeleteBy(t *testing.T) {
	ctx := context.Background()
	tdb := example.NewTestDB(ctx, t, "..")
	fxt := prepareTestFixture(ctx, t, tdb)
	packages := tables.NewPackageTable(tdb)

	affected, err := packages.DeleteByID(ctx, fxt.packageRows[0].KeyID(), fxt.packageRows[1].KeyID())
	if err != nil {
		t.Fatal(err)
	}
	if affected != 2 {
		t.Fatalf("Unexpected count: %d", affected)
	}
}

func prepareTestFixture(ctx context.Context, t *testing.T, tdb *example.TestDB) (fixture testFixture) {

	time1, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	if err != nil {
		t.Fatal(err)
	}
	packages := tables.NewPackageTable(tdb)
	products := tables.NewProductTable(tdb)
	joinTable := tables.NewPackageProductTable(tdb)

	fixture.packageRows = tables.NewPackageRows(
		tables.PackageData{ID: "955074e7-71fb-43f7-b33e-e3825900d4a3", Name: "Package for foo"},
		tables.PackageData{ID: "dca8850a-5a0a-478b-a833-1a7e2d202b97", Name: "Package for bar"},
		tables.PackageData{ID: "00ab6a7c-b9d2-422b-8d5d-58b894fc22d4", Name: "Package for baz"},
		tables.PackageData{ID: "6beececb-3d0e-42b2-9270-b96f75e5c346", Name: "Package for baz"},
	)

	fixture.productRows = tables.NewProductRows([]tables.ProductData{
		{nil, "1.99", "Product1", "product-1", time1, nil},
		{nil, "2.99", "Product2", "product-2", time1, nil},
	}...)

	var n int
	if n, err = packages.Insert(ctx, fixture.packageRows...); err != nil {
		t.Fatal(err)
	} else if n != len(fixture.packageRows) {
		t.Fatalf("Unexpected number of affected rows: %d", n)
	}
	if n, err = products.Insert(ctx, fixture.productRows...); err != nil {
		t.Fatal(err)
	} else if n != len(fixture.productRows) {
		t.Fatalf("Unexpected number of affected rows: %d", n)
	}

	fixture.packageProductRows = tables.NewPackageProductRows([]tables.PackageProductData{
		{PackageID: fixture.packageRows[0].GetID(), ProductID: fixture.productRows[0].GetID()},
		{PackageID: fixture.packageRows[0].GetID(), ProductID: fixture.productRows[1].GetID()},
	}...)

	if n, err = joinTable.Insert(ctx, fixture.packageProductRows...); err != nil {
		t.Fatal(err)
	} else if n != len(fixture.packageProductRows) {
		t.Fatalf("Unexpected number of affected rows: %d", n)
	}
	return fixture
}

type testFixture struct {
	packageRows        tables.PackageRows
	productRows        tables.ProductRows
	packageProductRows tables.PackageProductRows
}
