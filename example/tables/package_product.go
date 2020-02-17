// Code generated by github.com/hanpama/pgmg. DO NOT EDIT.
package tables

import (
	"context"
	"encoding/json"
)

// PackageProductRow represents a row for table "package_product"
type PackageProductRow struct {
	Data PackageProductData
}

type PackageProductData struct {
	PackageID string `json:"package_id"`
	ProductID int32  `json:"product_id"`
}

func NewPackageProductRow(data PackageProductData) *PackageProductRow {
	return &PackageProductRow{data}
}

func NewPackageProductRows(data ...PackageProductData) PackageProductRows {
	rows := make(PackageProductRows, len(data))
	for i, d := range data {
		rows[i] = NewPackageProductRow(d)
	}
	return rows
}

// GetPackageID gets value of column "package_id" from "package_product" row
func (r *PackageProductRow) GetPackageID() string { return r.Data.PackageID }

// SetPackageID sets value of column "package_id" in "package_product" row
func (r *PackageProductRow) SetPackageID(packageID string) { r.Data.PackageID = packageID }

// GetProductID gets value of column "product_id" from "package_product" row
func (r *PackageProductRow) GetProductID() int32 { return r.Data.ProductID }

// SetProductID sets value of column "product_id" in "package_product" row
func (r *PackageProductRow) SetProductID(productID int32) { r.Data.ProductID = productID }

// PackageProductPackageIDProductID represents key defined by UNIQUE constraint "package_product_package_id_product_id_key" for table "package_product"
type PackageProductPackageIDProductID struct {
	PackageID string `json:"package_id"`
	ProductID int32  `json:"product_id"`
}

func (r *PackageProductRow) KeyPackageIDProductID() PackageProductPackageIDProductID {
	return PackageProductPackageIDProductID{r.GetPackageID(), r.GetProductID()}
}

// PackageProductRows represents multiple rows for table "package_product"
type PackageProductRows []*PackageProductRow

func (rs PackageProductRows) KeyPackageIDProductID() (keys Keys) {
	keys = make(Keys, len(rs))
	for i, r := range rs {
		keys[i] = r.KeyPackageIDProductID()
	}
	return keys
}

func (r *PackageProductRow) RefPackageID() PackageID {
	return PackageID{r.GetPackageID()}
}

func (rs PackageProductRows) RefPackageID() (keys Keys) {
	keys = make(Keys, len(rs))
	for i, r := range rs {
		keys[i] = r.RefPackageID()
	}
	return keys
}
func (r *PackageProductRow) RefProductID() ProductID {
	return ProductID{r.GetProductID()}
}

func (rs PackageProductRows) RefProductID() (keys Keys) {
	keys = make(Keys, len(rs))
	for i, r := range rs {
		keys[i] = r.RefProductID()
	}
	return keys
}

// NewPackageProductTable(h SQLHandle) creates new PackageProductTable
func NewPackageProductTable(h SQLHandle) *PackageProductTable {
	return &PackageProductTable{h}
}

// PackageProductTable provides access methods for table "package_product"
type PackageProductTable struct {
	h SQLHandle
}

func (t *PackageProductTable) Find(ctx context.Context, filter PackageProductValues) (PackageProductRows, error) {
	return FindPackageProductRows(ctx, t.h, filter)
}

func (t *PackageProductTable) Count(ctx context.Context, filter PackageProductValues) (int, error) {
	return CountPackageProductRows(ctx, t.h, filter)
}

func (t *PackageProductTable) Update(ctx context.Context, changeset, filter PackageProductValues) (int64, error) {
	return UpdatePackageProductRows(ctx, t.h, changeset, filter)
}

func (t *PackageProductTable) Insert(ctx context.Context, rows ...*PackageProductRow) (int, error) {
	return InsertReturningPackageProductRows(ctx, t.h, rows...)
}

func (t *PackageProductTable) Delete(ctx context.Context, filter PackageProductValues) (int64, error) {
	return DeletePackageProductRows(ctx, t.h, filter)
}

func (t *PackageProductTable) Save(ctx context.Context, rows ...*PackageProductRow) error {
	return SaveReturningPackageProductRows(ctx, t.h, rows...)
}

func (t *PackageProductTable) GetByPackageIDProductID(ctx context.Context, keys ...interface{}) (PackageProductRows, error) {
	return GetPackageProductRowsByPackageIDProductID(ctx, t.h, keys...)
}

func (t *PackageProductTable) UpdateByPackageIDProductID(ctx context.Context, changeset PackageProductValues, keys ...interface{}) (int64, error) {
	return UpdatePackageProductRowsByPackageIDProductID(ctx, t.h, changeset, keys...)
}

func (t *PackageProductTable) DeleteByPackageIDProductID(ctx context.Context, keys ...interface{}) (int64, error) {
	return DeletePackageProductRowsByPackageIDProductID(ctx, t.h, keys...)
}

type PackageProductValues struct {
	PackageID *string `json:"package_id"`
	ProductID *int32  `json:"product_id"`
}

// InsertPackageProductRows inserts the rows into table "package_product"
func InsertPackageProductRows(ctx context.Context, db SQLHandle, rows ...*PackageProductRow) (numRows int64, err error) {
	numRows, err = execWithJSONArgs(ctx, db, SQLInsertPackageProductRows, rows)
	if err != nil {
		return numRows, formatError("InsertPackageProductRows", err)
	}
	return numRows, nil
}

// InsertReturningPackageProductRows inserts the rows into table "package_product" and returns the rows.
func InsertReturningPackageProductRows(ctx context.Context, db SQLHandle, inputs ...*PackageProductRow) (numRows int, err error) {
	rows := PackageProductRows(inputs)
	numRows, err = queryWithJSONArgs(ctx, db, rows.ReceiveRows, SQLInsertReturningPackageProductRows, rows)
	if err != nil {
		return numRows, formatError("InsertReturningPackageProductRows", err)
	}
	return numRows, nil
}

// FindPackageProductRows finds the rows matching the condition from table "package_product"
func FindPackageProductRows(ctx context.Context, db SQLHandle, cond PackageProductValues) (rows PackageProductRows, err error) {
	if _, err = queryWithJSONArgs(ctx, db, rows.ReceiveRows, SQLFindPackageProductRows, cond); err != nil {
		return nil, formatError("FindPackageProductRows", err)
	}
	return rows, nil
}

// DeletePackageProductRows deletes the rows matching the condition from table "package_product"
func DeletePackageProductRows(ctx context.Context, db SQLHandle, cond PackageProductValues) (numRows int64, err error) {
	if numRows, err = execWithJSONArgs(ctx, db, SQLDeletePackageProductRows, cond); err != nil {
		return numRows, formatError("DeletePackageProductRows", err)
	}
	return numRows, nil
}

func UpdatePackageProductRows(ctx context.Context, db SQLHandle, changeset, filter PackageProductValues) (numRows int64, err error) {
	if numRows, err = execWithJSONArgs(ctx, db, SQLUpdatePackageProductRows, changeset, filter); err != nil {
		return numRows, formatError("UpdatePackageProductRows", err)
	}
	return numRows, nil
}

// CountPackageProductRows counts the number of rows matching the condition from table "package_product"
func CountPackageProductRows(ctx context.Context, db SQLHandle, cond PackageProductValues) (count int, err error) {
	if _, err = queryWithJSONArgs(ctx, db, func(int) []interface{} { return []interface{}{&count} }, SQLCountPackageProductRows, cond); err != nil {
		return 0, formatError("CountPackageProductRows", err)
	}
	return count, nil
}

// SavePackageProductRows upserts the given rows for table "package_product" checking uniqueness by contstraint "package_product_package_id_product_id_key"
func SavePackageProductRows(ctx context.Context, db SQLHandle, rows ...*PackageProductRow) (err error) {
	_, err = execWithJSONArgs(ctx, db, SQLSavePackageProductRows, rows)
	if err != nil {
		return formatError("SavePackageProductRows", err)
	}
	return nil
}

// SaveReturningPackageProductRows upserts the given rows for table "package_product" checking uniqueness by contstraint "package_product_package_id_product_id_key"
// It returns the new values and scan them into given row references.
func SaveReturningPackageProductRows(ctx context.Context, db SQLHandle, inputs ...*PackageProductRow) (err error) {
	rows := PackageProductRows(inputs)
	_, err = queryWithJSONArgs(ctx, db, rows.ReceiveRows, SQLSaveReturningPackageProductRows, rows)
	if err != nil {
		return formatError("SaveReturningPackageProductRows", err)
	}
	return nil
}

// GetPackageProductRowsByPackageIDProductID gets matching rows for given PackageIDProductID keys from table "package_product"
func GetPackageProductRowsByPackageIDProductID(ctx context.Context, db SQLHandle, keys ...interface{}) (rows PackageProductRows, err error) {
	rows = make(PackageProductRows, 0, len(keys))
	if _, err = queryWithJSONArgs(ctx, db, rows.ReceiveRows, SQLGetPackageProductRowsByPackageIDProductID, Keys(keys)); err != nil {
		return nil, formatError("GetPackageProductRowsByPackageIDProductID", err)
	}
	return rows, nil
}

// DeletePackageProductRowsByPackageIDProductID deletes matching rows by PackageProductPackageIDProductID keys from table "package_product"
func DeletePackageProductRowsByPackageIDProductID(ctx context.Context, db SQLHandle, keys ...interface{}) (numRows int64, err error) {
	numRows, err = execWithJSONArgs(ctx, db, SQLDeletePackageProductRowsByPackageIDProductID, keys)
	if err != nil {
		return numRows, formatError("DeletePackageProductRowsByPackageIDProductID", err)
	}
	return numRows, nil
}

// UpdatePackageProductRowsByPackageIDProductID deletes matching rows by PackageProductPackageIDProductID keys from table "package_product"
func UpdatePackageProductRowsByPackageIDProductID(ctx context.Context, db SQLHandle, changeset PackageProductValues, keys ...interface{}) (numRows int64, err error) {
	numRows, err = execWithJSONArgs(ctx, db, SQLUpdatePackageProductRowsByPackageIDProductID, changeset, keys)
	if err != nil {
		return numRows, formatError("UpdatePackageProductRowsByPackageIDProductID", err)
	}
	return numRows, nil
}

// ReceiveRow returns all pointers of the column values for scanning
func (r *PackageProductRow) ReceiveRow() []interface{} {
	return []interface{}{&r.Data.PackageID, &r.Data.ProductID}
}

// ReceiveRows returns pointer slice to receive data for the row on index i
func (rs *PackageProductRows) ReceiveRows(i int) []interface{} {
	if len(*rs) <= i {
		*rs = append(*rs, new(PackageProductRow))
	} else if (*rs)[i] == nil {
		(*rs)[i] = new(PackageProductRow)
	}
	return (*rs)[i].ReceiveRow()
}

func (r *PackageProductRow) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Data)
}

var (
	SQLFindPackageProductRows = `
		WITH __f AS (SELECT "package_id", "product_id" FROM json_populate_record(null::"wise"."package_product", $1))
		SELECT __t.package_id, __t.product_id
		FROM "wise"."package_product" AS __t
		WHERE ((SELECT __f."package_id" IS NULL FROM __f) OR (SELECT __f."package_id" = __t."package_id" FROM __f))
			AND ((SELECT __f."product_id" IS NULL FROM __f) OR (SELECT __f."product_id" = __t."product_id" FROM __f))`
	SQLCountPackageProductRows = `
		WITH __f AS (SELECT "package_id", "product_id" FROM json_populate_record(null::"wise"."package_product", $1))
		SELECT count(*) FROM "wise"."package_product" AS __t
		WHERE ((SELECT __f."package_id" IS NULL FROM __f) OR (SELECT __f."package_id" = __t."package_id" FROM __f))
			AND ((SELECT __f."product_id" IS NULL FROM __f) OR (SELECT __f."product_id" = __t."product_id" FROM __f))`
	SQLReturningPackageProductRows = `
		RETURNING "package_id", "product_id"`
	SQLInsertPackageProductRows = `
		WITH __v AS (SELECT * FROM json_populate_recordset(null::"wise"."package_product", $1))
		INSERT INTO "wise"."package_product" AS __t ("package_id", "product_id")
		SELECT 
			__v."package_id", 
			__v."product_id" FROM __v`
	SQLInsertReturningPackageProductRows = SQLInsertPackageProductRows + SQLReturningPackageProductRows
	SQLDeletePackageProductRows          = `
		DELETE FROM "wise"."package_product" AS __t
		WHERE TRUE
			AND (($1::json->>'package_id' IS NULL) OR CAST($1::json->>'package_id' AS uuid) = __t."package_id")
			AND (($1::json->>'product_id' IS NULL) OR CAST($1::json->>'product_id' AS integer) = __t."product_id")`
	SQLDeleteReturningPackageProductRows = SQLDeletePackageProductRows + SQLReturningPackageProductRows
	SQLUpdatePackageProductRows          = `
		WITH __v AS (SELECT * FROM json_populate_record(null::"wise"."package_product", $1)),
			__f AS (SELECT * FROM json_populate_record(null::"wise"."package_product", $2))
		UPDATE "wise"."package_product" AS __t
		SET ("package_id", "product_id") = (SELECT 
			COALESCE(__v."package_id", __t."package_id"), 
			COALESCE(__v."product_id", __t."product_id") FROM __v)
		WHERE ((SELECT __f."package_id" IS NULL FROM __f) OR (SELECT __f."package_id" = __t."package_id" FROM __f))
			AND ((SELECT __f."product_id" IS NULL FROM __f) OR (SELECT __f."product_id" = __t."product_id" FROM __f))`
	SQLUpdateReturningPackageProductRows = SQLUpdatePackageProductRows + SQLReturningPackageProductRows
	SQLReplacePackageProductRows         = `
		WITH __v AS (SELECT * FROM json_populate_recordset(null::"wise"."package_product", $1))
		UPDATE "wise"."package_product" AS __t
			SET ("package_id", "product_id") = (SELECT 
				COALESCE(__v."package_id", __t."package_id"), 
				COALESCE(__v."product_id", __t."product_id")
			FROM __v WHERE __v."package_id" = __t."package_id" AND __v."product_id" = __t."product_id")
		FROM __v WHERE __v."package_id" = __t."package_id" AND __v."product_id" = __t."product_id"`
	SQLReplaceReturningPackageProductRows = SQLReplacePackageProductRows + SQLReturningPackageProductRows
	SQLSavePackageProductRows             = `
		WITH __v AS (SELECT * FROM json_populate_recordset(null::"wise"."package_product", $1))
		INSERT INTO "wise"."package_product" AS __t ("package_id", "product_id")
		SELECT 
			__v."package_id", 
			__v."product_id" FROM __v
		ON CONFLICT ("package_id", "product_id") DO UPDATE
		SET ("package_id", "product_id") = (
			SELECT "package_id", "product_id" FROM __v
			WHERE __v."package_id" = __t."package_id"
				AND __v."product_id" = __t."product_id"
		)`
	SQLSaveReturningPackageProductRows           = SQLSavePackageProductRows + SQLReturningPackageProductRows
	SQLGetPackageProductRowsByPackageIDProductID = `
		WITH __key AS (SELECT DISTINCT "package_id", "product_id" FROM json_populate_recordset(null::"wise"."package_product", $1))
		SELECT "package_id", "product_id"
		FROM __key JOIN "wise"."package_product" AS __t USING ("package_id", "product_id")`
	SQLUpdatePackageProductRowsByPackageIDProductID = `
		WITH __v AS (SELECT * FROM json_populate_record(null::"wise"."package_product", $1)),
		  __key AS (SELECT package_id, product_id FROM json_populate_recordset(null::"wise"."package_product", $2))
		UPDATE "wise"."package_product" AS __t
		SET ("package_id", "product_id") = (SELECT
			COALESCE(__v."package_id", __t."package_id"), 
			COALESCE(__v."product_id", __t."product_id")
		FROM __v)
		FROM __key WHERE (__key."package_id" = __t."package_id")AND (__key."product_id" = __t."product_id")`
	SQLDeletePackageProductRowsByPackageIDProductID = `
		WITH __key AS (SELECT package_id, product_id FROM json_populate_recordset(null::"wise"."package_product", $1))
		DELETE FROM "wise"."package_product" AS __t USING __key WHERE (__key."package_id" = __t."package_id") AND (__key."product_id" = __t."product_id")`
	SQLDeleteReturningPackageProductRowsByPackageIDProductID = SQLDeletePackageProductRowsByPackageIDProductID + SQLReturningPackageProductRows
)
