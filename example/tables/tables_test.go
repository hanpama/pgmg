package tables_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/hanpama/pgmg/example"

	"github.com/hanpama/pgmg/example/tables"
	_ "github.com/lib/pq"
)

func TestGetByID(t *testing.T) {
	ctx := context.Background()
	fxt, err := prepareTestEnv(ctx, "..")
	if err != nil {
		t.Fatal(err)
	}
	packages := tables.NewPackageTable(fxt.tdb)

	row1, row2 := fxt.packageRows[0], fxt.packageRows[1]
	rows, err := packages.GetByID(ctx, row1.KeyID(), row2.KeyID(), row1.KeyID())
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 2 {
		t.Fatalf("Unexpected length: %d", len(rows))
	}
	if !reflect.DeepEqual(row1, rows[*row1.KeyID()]) {
		t.Fatal()
	}
	if !reflect.DeepEqual(row2, rows[*row2.KeyID()]) {
		t.Fatal()
	}
}

func TestFind(t *testing.T) {
	ctx := context.Background()
	fxt, err := prepareTestEnv(ctx, "..")
	if err != nil {
		t.Fatal(err)
	}
	packages := tables.NewPackageTable(fxt.tdb)

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
	fxt, err := prepareTestEnv(ctx, "..")
	if err != nil {
		t.Fatal(err)
	}
	packages := tables.NewPackageTable(fxt.tdb)

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
	fxt, err := prepareTestEnv(ctx, "..")
	if err != nil {
		t.Fatal(err)
	}
	packages := tables.NewPackageTable(fxt.tdb)

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
	fxt, err := prepareTestEnv(ctx, "..")
	if err != nil {
		t.Fatal(err)
	}
	packages := tables.NewPackageTable(fxt.tdb)

	// Save
	fxt.packageRows[1].SetName("Package for foo")
	if err := packages.Save(ctx, fxt.packageRows...); err != nil {
		t.Fatal(err)
	}
}

func TestUpdateBy(t *testing.T) {
	ctx := context.Background()
	fxt, err := prepareTestEnv(ctx, "..")
	if err != nil {
		t.Fatal(err)
	}
	packages := tables.NewPackageTable(fxt.tdb)

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
	fxt, err := prepareTestEnv(ctx, "..")
	if err != nil {
		t.Fatal(err)
	}
	packages := tables.NewPackageTable(fxt.tdb)

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
	fxt, err := prepareTestEnv(ctx, "..")
	if err != nil {
		t.Fatal(err)
	}
	packages := tables.NewPackageTable(fxt.tdb)

	affected, err := packages.DeleteByID(ctx, fxt.packageRows[0].KeyID(), fxt.packageRows[1].KeyID())
	if err != nil {
		t.Fatal(err)
	}
	if affected != 2 {
		t.Fatalf("Unexpected count: %d", affected)
	}
}

func TestForeignKey(t *testing.T) {
	ctx := context.Background()
	fxt, err := prepareTestEnv(ctx, "..")
	if err != nil {
		t.Fatal(err)
	}

	pops := tables.NewPopTable(fxt.tdb)

	rows, err := pops.GetByNameYear(ctx, fxt.campaignRows.RefPopNamePopYear()...)
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 1 {
		t.Fatal()
	}
	if rows[*fxt.popRows[0].KeyNameYear()] == nil {
		t.Fatal()
	}
}

func prepareTestEnv(ctx context.Context, migdir string) (fixture *testEnv, err error) {
	fixture = new(testEnv)

	if fixture.tdb, err = example.NewTestDB(ctx, migdir); err != nil {
		return nil, err
	}

	time1, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	if err != nil {
		return nil, err
	}
	packages := tables.NewPackageTable(fixture.tdb)
	products := tables.NewProductTable(fixture.tdb)
	joinTable := tables.NewPackageProductTable(fixture.tdb)

	fixture.packageRows = tables.NewPackageRows(
		tables.PackageData{ID: "955074e7-71fb-43f7-b33e-e3825900d4a3", Name: "Package for foo"},
		tables.PackageData{ID: "dca8850a-5a0a-478b-a833-1a7e2d202b97", Name: "Package for bar"},
		tables.PackageData{ID: "00ab6a7c-b9d2-422b-8d5d-58b894fc22d4", Name: "Package for baz"},
		tables.PackageData{ID: "6beececb-3d0e-42b2-9270-b96f75e5c346", Name: "Package for baz"},
	)
	fixture.productRows = make(tables.ProductRows, 100)

	for i := 0; i < 100; i++ {
		fixture.productRows[i] = tables.NewProductRow(tables.ProductData{
			nil, fmt.Sprintf("%d.99", i), fmt.Sprintf("Product%d", i), fmt.Sprintf("product-%d", i), time1, nil,
		})
	}

	var n int
	if n, err = packages.Insert(ctx, fixture.packageRows...); err != nil {
		return nil, err
	} else if n != len(fixture.packageRows) {
		return nil, fmt.Errorf("Unexpected number of affected rows: %d", n)
	}
	if n, err = products.Insert(ctx, fixture.productRows...); err != nil {
		return nil, err
	} else if n != len(fixture.productRows) {
		return nil, fmt.Errorf("Unexpected number of affected rows: %d", n)
	}

	fixture.packageProductRows = tables.NewPackageProductRows([]tables.PackageProductData{
		{PackageID: fixture.packageRows[0].GetID(), ProductID: fixture.productRows[0].GetID()},
		{PackageID: fixture.packageRows[0].GetID(), ProductID: fixture.productRows[1].GetID()},
	}...)

	if n, err = joinTable.Insert(ctx, fixture.packageProductRows...); err != nil {
		return nil, err
	} else if n != len(fixture.packageProductRows) {
		return nil, fmt.Errorf("Unexpected number of affected rows: %d", n)
	}

	pops := tables.NewPopTable(fixture.tdb)
	popRows := tables.NewPopRows(
		tables.PopData{Name: "Foo", Year: 2015, Description: "Foo"},
		tables.PopData{Name: "Foo", Year: 2016, Description: "Foo"},
	)
	if _, err = pops.Insert(ctx, popRows...); err != nil {
		return nil, err
	}
	fixture.popRows = popRows

	campaigns := tables.NewCampaignTable(fixture.tdb)
	campaignRows := tables.NewCampaignRows(
		tables.CampaignData{ID: "2642476b-8a76-4a97-b2cf-dd0a882ad151", PopName: nil, PopYear: nil},
		tables.CampaignData{ID: "95c9abd0-8344-49a1-92e9-2b3d9bc1a8f8", PopName: &popRows[0].Data.Name, PopYear: &popRows[0].Data.Year},
	)
	if _, err = campaigns.Insert(ctx, campaignRows...); err != nil {
		return nil, err
	}
	fixture.campaignRows = campaignRows

	return fixture, nil
}

type testEnv struct {
	tdb                *example.TestDB
	packageRows        tables.PackageRows
	productRows        tables.ProductRows
	packageProductRows tables.PackageProductRows
	popRows            tables.PopRows
	campaignRows       tables.CampaignRows
}
