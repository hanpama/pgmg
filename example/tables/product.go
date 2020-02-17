// Code generated by github.com/hanpama/pgmg. DO NOT EDIT.
package tables

import (
	"context"
	"encoding/json"
	"time"
)

// ProductRow represents a row for table "product"
type ProductRow struct {
	Data ProductData
}

type ProductData struct {
	ID      *int32     `json:"id"`
	Price   string     `json:"price"`
	Name    string     `json:"name"`
	Alias   string     `json:"alias"`
	Stocked time.Time  `json:"stocked"`
	Sold    *time.Time `json:"sold"`
}

func NewProductRow(data ProductData) *ProductRow {
	return &ProductRow{data}
}

func NewProductRows(data ...ProductData) ProductRows {
	rows := make(ProductRows, len(data))
	for i, d := range data {
		rows[i] = NewProductRow(d)
	}
	return rows
}

// GetID gets value of column "id" from "product" row
func (r *ProductRow) GetID() int32 { return *r.Data.ID }

// SetID sets value of column "id" in "product" row
func (r *ProductRow) SetID(id int32) { r.Data.ID = &id }

// ClearID sets value of column "id" null in "product" row
func (r *ProductRow) ClearID() { r.Data.ID = nil }

// HasValidID checks to value of column "id" is not null
func (r *ProductRow) HasValidID() bool { return r.Data.ID != nil }

// GetPrice gets value of column "price" from "product" row
func (r *ProductRow) GetPrice() string { return r.Data.Price }

// SetPrice sets value of column "price" in "product" row
func (r *ProductRow) SetPrice(price string) { r.Data.Price = price }

// GetName gets value of column "name" from "product" row
func (r *ProductRow) GetName() string { return r.Data.Name }

// SetName sets value of column "name" in "product" row
func (r *ProductRow) SetName(name string) { r.Data.Name = name }

// GetAlias gets value of column "alias" from "product" row
func (r *ProductRow) GetAlias() string { return r.Data.Alias }

// SetAlias sets value of column "alias" in "product" row
func (r *ProductRow) SetAlias(alias string) { r.Data.Alias = alias }

// GetStocked gets value of column "stocked" from "product" row
func (r *ProductRow) GetStocked() time.Time { return r.Data.Stocked }

// SetStocked sets value of column "stocked" in "product" row
func (r *ProductRow) SetStocked(stocked time.Time) { r.Data.Stocked = stocked }

// GetSold gets value of column "sold" from "product" row
func (r *ProductRow) GetSold() time.Time { return *r.Data.Sold }

// SetSold sets value of column "sold" in "product" row
func (r *ProductRow) SetSold(sold time.Time) { r.Data.Sold = &sold }

// ClearSold sets value of column "sold" null in "product" row
func (r *ProductRow) ClearSold() { r.Data.Sold = nil }

// HasValidSold checks to value of column "sold" is not null
func (r *ProductRow) HasValidSold() bool { return r.Data.Sold != nil }

// ProductID represents key defined by PRIMARY KEY constraint "product_pkey" for table "product"
type ProductID struct {
	ID int32 `json:"id"`
}

func (r *ProductRow) KeyID() ProductID {
	return ProductID{r.GetID()}
}

// ProductAlias represents key defined by UNIQUE constraint "product_alias_key" for table "product"
type ProductAlias struct {
	Alias string `json:"alias"`
}

func (r *ProductRow) KeyAlias() ProductAlias {
	return ProductAlias{r.GetAlias()}
}

// ProductRows represents multiple rows for table "product"
type ProductRows []*ProductRow

func (rs ProductRows) KeyID() (keys Keys) {
	keys = make(Keys, len(rs))
	for i, r := range rs {
		keys[i] = r.KeyID()
	}
	return keys
}

func (rs ProductRows) KeyAlias() (keys Keys) {
	keys = make(Keys, len(rs))
	for i, r := range rs {
		keys[i] = r.KeyAlias()
	}
	return keys
}

// NewProductTable(h SQLHandle) creates new ProductTable
func NewProductTable(h SQLHandle) *ProductTable {
	return &ProductTable{h}
}

// ProductTable provides access methods for table "product"
type ProductTable struct {
	h SQLHandle
}

func (t *ProductTable) Find(ctx context.Context, filter ProductValues) (ProductRows, error) {
	return FindProductRows(ctx, t.h, filter)
}

func (t *ProductTable) Count(ctx context.Context, filter ProductValues) (int, error) {
	return CountProductRows(ctx, t.h, filter)
}

func (t *ProductTable) Update(ctx context.Context, changeset, filter ProductValues) (int64, error) {
	return UpdateProductRows(ctx, t.h, changeset, filter)
}

func (t *ProductTable) Insert(ctx context.Context, rows ...*ProductRow) (int, error) {
	return InsertReturningProductRows(ctx, t.h, rows...)
}

func (t *ProductTable) Delete(ctx context.Context, filter ProductValues) (int64, error) {
	return DeleteProductRows(ctx, t.h, filter)
}

func (t *ProductTable) Save(ctx context.Context, rows ...*ProductRow) error {
	return SaveReturningProductRows(ctx, t.h, rows...)
}

func (t *ProductTable) GetByID(ctx context.Context, keys ...interface{}) (ProductRows, error) {
	return GetProductRowsByID(ctx, t.h, keys...)
}

func (t *ProductTable) UpdateByID(ctx context.Context, changeset ProductValues, keys ...interface{}) (int64, error) {
	return UpdateProductRowsByID(ctx, t.h, changeset, keys...)
}

func (t *ProductTable) DeleteByID(ctx context.Context, keys ...interface{}) (int64, error) {
	return DeleteProductRowsByID(ctx, t.h, keys...)
}

func (t *ProductTable) GetByAlias(ctx context.Context, keys ...interface{}) (ProductRows, error) {
	return GetProductRowsByAlias(ctx, t.h, keys...)
}

func (t *ProductTable) UpdateByAlias(ctx context.Context, changeset ProductValues, keys ...interface{}) (int64, error) {
	return UpdateProductRowsByAlias(ctx, t.h, changeset, keys...)
}

func (t *ProductTable) DeleteByAlias(ctx context.Context, keys ...interface{}) (int64, error) {
	return DeleteProductRowsByAlias(ctx, t.h, keys...)
}

type ProductValues struct {
	ID      *int32     `json:"id"`
	Price   *string    `json:"price"`
	Name    *string    `json:"name"`
	Alias   *string    `json:"alias"`
	Stocked *time.Time `json:"stocked"`
	Sold    *time.Time `json:"sold"`
}

// InsertProductRows inserts the rows into table "product"
func InsertProductRows(ctx context.Context, db SQLHandle, rows ...*ProductRow) (numRows int64, err error) {
	numRows, err = execWithJSONArgs(ctx, db, SQLInsertProductRows, rows)
	if err != nil {
		return numRows, formatError("InsertProductRows", err)
	}
	return numRows, nil
}

// InsertReturningProductRows inserts the rows into table "product" and returns the rows.
func InsertReturningProductRows(ctx context.Context, db SQLHandle, inputs ...*ProductRow) (numRows int, err error) {
	rows := ProductRows(inputs)
	numRows, err = queryWithJSONArgs(ctx, db, rows.ReceiveRows, SQLInsertReturningProductRows, rows)
	if err != nil {
		return numRows, formatError("InsertReturningProductRows", err)
	}
	return numRows, nil
}

// FindProductRows finds the rows matching the condition from table "product"
func FindProductRows(ctx context.Context, db SQLHandle, cond ProductValues) (rows ProductRows, err error) {
	if _, err = queryWithJSONArgs(ctx, db, rows.ReceiveRows, SQLFindProductRows, cond); err != nil {
		return nil, formatError("FindProductRows", err)
	}
	return rows, nil
}

// DeleteProductRows deletes the rows matching the condition from table "product"
func DeleteProductRows(ctx context.Context, db SQLHandle, cond ProductValues) (numRows int64, err error) {
	if numRows, err = execWithJSONArgs(ctx, db, SQLDeleteProductRows, cond); err != nil {
		return numRows, formatError("DeleteProductRows", err)
	}
	return numRows, nil
}

func UpdateProductRows(ctx context.Context, db SQLHandle, changeset, filter ProductValues) (numRows int64, err error) {
	if numRows, err = execWithJSONArgs(ctx, db, SQLUpdateProductRows, changeset, filter); err != nil {
		return numRows, formatError("UpdateProductRows", err)
	}
	return numRows, nil
}

// CountProductRows counts the number of rows matching the condition from table "product"
func CountProductRows(ctx context.Context, db SQLHandle, cond ProductValues) (count int, err error) {
	if _, err = queryWithJSONArgs(ctx, db, func(int) []interface{} { return []interface{}{&count} }, SQLCountProductRows, cond); err != nil {
		return 0, formatError("CountProductRows", err)
	}
	return count, nil
}

// SaveProductRows upserts the given rows for table "product" checking uniqueness by contstraint "product_pkey"
func SaveProductRows(ctx context.Context, db SQLHandle, rows ...*ProductRow) (err error) {
	_, err = execWithJSONArgs(ctx, db, SQLSaveProductRows, rows)
	if err != nil {
		return formatError("SaveProductRows", err)
	}
	return nil
}

// SaveReturningProductRows upserts the given rows for table "product" checking uniqueness by contstraint "product_pkey"
// It returns the new values and scan them into given row references.
func SaveReturningProductRows(ctx context.Context, db SQLHandle, inputs ...*ProductRow) (err error) {
	rows := ProductRows(inputs)
	_, err = queryWithJSONArgs(ctx, db, rows.ReceiveRows, SQLSaveReturningProductRows, rows)
	if err != nil {
		return formatError("SaveReturningProductRows", err)
	}
	return nil
}

// GetProductRowsByID gets matching rows for given ID keys from table "product"
func GetProductRowsByID(ctx context.Context, db SQLHandle, keys ...interface{}) (rows ProductRows, err error) {
	rows = make(ProductRows, 0, len(keys))
	if _, err = queryWithJSONArgs(ctx, db, rows.ReceiveRows, SQLGetProductRowsByID, Keys(keys)); err != nil {
		return nil, formatError("GetProductRowsByID", err)
	}
	return rows, nil
}

// DeleteProductRowsByID deletes matching rows by ProductID keys from table "product"
func DeleteProductRowsByID(ctx context.Context, db SQLHandle, keys ...interface{}) (numRows int64, err error) {
	numRows, err = execWithJSONArgs(ctx, db, SQLDeleteProductRowsByID, keys)
	if err != nil {
		return numRows, formatError("DeleteProductRowsByID", err)
	}
	return numRows, nil
}

// UpdateProductRowsByID deletes matching rows by ProductID keys from table "product"
func UpdateProductRowsByID(ctx context.Context, db SQLHandle, changeset ProductValues, keys ...interface{}) (numRows int64, err error) {
	numRows, err = execWithJSONArgs(ctx, db, SQLUpdateProductRowsByID, changeset, keys)
	if err != nil {
		return numRows, formatError("UpdateProductRowsByID", err)
	}
	return numRows, nil
}

// GetProductRowsByAlias gets matching rows for given Alias keys from table "product"
func GetProductRowsByAlias(ctx context.Context, db SQLHandle, keys ...interface{}) (rows ProductRows, err error) {
	rows = make(ProductRows, 0, len(keys))
	if _, err = queryWithJSONArgs(ctx, db, rows.ReceiveRows, SQLGetProductRowsByAlias, Keys(keys)); err != nil {
		return nil, formatError("GetProductRowsByAlias", err)
	}
	return rows, nil
}

// DeleteProductRowsByAlias deletes matching rows by ProductAlias keys from table "product"
func DeleteProductRowsByAlias(ctx context.Context, db SQLHandle, keys ...interface{}) (numRows int64, err error) {
	numRows, err = execWithJSONArgs(ctx, db, SQLDeleteProductRowsByAlias, keys)
	if err != nil {
		return numRows, formatError("DeleteProductRowsByAlias", err)
	}
	return numRows, nil
}

// UpdateProductRowsByAlias deletes matching rows by ProductAlias keys from table "product"
func UpdateProductRowsByAlias(ctx context.Context, db SQLHandle, changeset ProductValues, keys ...interface{}) (numRows int64, err error) {
	numRows, err = execWithJSONArgs(ctx, db, SQLUpdateProductRowsByAlias, changeset, keys)
	if err != nil {
		return numRows, formatError("UpdateProductRowsByAlias", err)
	}
	return numRows, nil
}

// ReceiveRow returns all pointers of the column values for scanning
func (r *ProductRow) ReceiveRow() []interface{} {
	return []interface{}{&r.Data.ID, &r.Data.Price, &r.Data.Name, &r.Data.Alias, &r.Data.Stocked, &r.Data.Sold}
}

// ReceiveRows returns pointer slice to receive data for the row on index i
func (rs *ProductRows) ReceiveRows(i int) []interface{} {
	if len(*rs) <= i {
		*rs = append(*rs, new(ProductRow))
	} else if (*rs)[i] == nil {
		(*rs)[i] = new(ProductRow)
	}
	return (*rs)[i].ReceiveRow()
}

func (r *ProductRow) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Data)
}

var (
	SQLFindProductRows = `
		WITH __f AS (SELECT "id", "price", "name", "alias", "stocked", "sold" FROM json_populate_record(null::"wise"."product", $1))
		SELECT __t.id, __t.price, __t.name, __t.alias, __t.stocked, __t.sold
		FROM "wise"."product" AS __t
		WHERE ((SELECT __f."id" IS NULL FROM __f) OR (SELECT __f."id" = __t."id" FROM __f))
			AND ((SELECT __f."price" IS NULL FROM __f) OR (SELECT __f."price" = __t."price" FROM __f))
			AND ((SELECT __f."name" IS NULL FROM __f) OR (SELECT __f."name" = __t."name" FROM __f))
			AND ((SELECT __f."alias" IS NULL FROM __f) OR (SELECT __f."alias" = __t."alias" FROM __f))
			AND ((SELECT __f."stocked" IS NULL FROM __f) OR (SELECT __f."stocked" = __t."stocked" FROM __f))
			AND ((SELECT __f."sold" IS NULL FROM __f) OR (SELECT __f."sold" = __t."sold" FROM __f))`
	SQLCountProductRows = `
		WITH __f AS (SELECT "id", "price", "name", "alias", "stocked", "sold" FROM json_populate_record(null::"wise"."product", $1))
		SELECT count(*) FROM "wise"."product" AS __t
		WHERE ((SELECT __f."id" IS NULL FROM __f) OR (SELECT __f."id" = __t."id" FROM __f))
			AND ((SELECT __f."price" IS NULL FROM __f) OR (SELECT __f."price" = __t."price" FROM __f))
			AND ((SELECT __f."name" IS NULL FROM __f) OR (SELECT __f."name" = __t."name" FROM __f))
			AND ((SELECT __f."alias" IS NULL FROM __f) OR (SELECT __f."alias" = __t."alias" FROM __f))
			AND ((SELECT __f."stocked" IS NULL FROM __f) OR (SELECT __f."stocked" = __t."stocked" FROM __f))
			AND ((SELECT __f."sold" IS NULL FROM __f) OR (SELECT __f."sold" = __t."sold" FROM __f))`
	SQLReturningProductRows = `
		RETURNING "id", "price", "name", "alias", "stocked", "sold"`
	SQLInsertProductRows = `
		WITH __v AS (SELECT * FROM json_populate_recordset(null::"wise"."product", $1))
		INSERT INTO "wise"."product" AS __t ("id", "price", "name", "alias", "stocked", "sold")
		SELECT 
			COALESCE(__v."id", nextval('wise.product_id_seq'::regclass)), 
			__v."price", 
			__v."name", 
			__v."alias", 
			__v."stocked", 
			__v."sold" FROM __v`
	SQLInsertReturningProductRows = SQLInsertProductRows + SQLReturningProductRows
	SQLDeleteProductRows          = `
		DELETE FROM "wise"."product" AS __t
		WHERE TRUE
			AND (($1::json->>'id' IS NULL) OR CAST($1::json->>'id' AS integer) = __t."id")
			AND (($1::json->>'price' IS NULL) OR CAST($1::json->>'price' AS money) = __t."price")
			AND (($1::json->>'name' IS NULL) OR CAST($1::json->>'name' AS character varying) = __t."name")
			AND (($1::json->>'alias' IS NULL) OR CAST($1::json->>'alias' AS character varying) = __t."alias")
			AND (($1::json->>'stocked' IS NULL) OR CAST($1::json->>'stocked' AS timestamp with time zone) = __t."stocked")
			AND (($1::json->>'sold' IS NULL) OR CAST($1::json->>'sold' AS timestamp with time zone) = __t."sold")`
	SQLDeleteReturningProductRows = SQLDeleteProductRows + SQLReturningProductRows
	SQLUpdateProductRows          = `
		WITH __v AS (SELECT * FROM json_populate_record(null::"wise"."product", $1)),
			__f AS (SELECT * FROM json_populate_record(null::"wise"."product", $2))
		UPDATE "wise"."product" AS __t
		SET ("id", "price", "name", "alias", "stocked", "sold") = (SELECT 
			COALESCE(__v."id", __t."id"), 
			COALESCE(__v."price", __t."price"), 
			COALESCE(__v."name", __t."name"), 
			COALESCE(__v."alias", __t."alias"), 
			COALESCE(__v."stocked", __t."stocked"), 
			COALESCE(__v."sold", __t."sold") FROM __v)
		WHERE ((SELECT __f."id" IS NULL FROM __f) OR (SELECT __f."id" = __t."id" FROM __f))
			AND ((SELECT __f."price" IS NULL FROM __f) OR (SELECT __f."price" = __t."price" FROM __f))
			AND ((SELECT __f."name" IS NULL FROM __f) OR (SELECT __f."name" = __t."name" FROM __f))
			AND ((SELECT __f."alias" IS NULL FROM __f) OR (SELECT __f."alias" = __t."alias" FROM __f))
			AND ((SELECT __f."stocked" IS NULL FROM __f) OR (SELECT __f."stocked" = __t."stocked" FROM __f))
			AND ((SELECT __f."sold" IS NULL FROM __f) OR (SELECT __f."sold" = __t."sold" FROM __f))`
	SQLUpdateReturningProductRows = SQLUpdateProductRows + SQLReturningProductRows
	SQLReplaceProductRows         = `
		WITH __v AS (SELECT * FROM json_populate_recordset(null::"wise"."product", $1))
		UPDATE "wise"."product" AS __t
			SET ("id", "price", "name", "alias", "stocked", "sold") = (SELECT 
				COALESCE(__v."id", __t."id"), 
				COALESCE(__v."price", __t."price"), 
				COALESCE(__v."name", __t."name"), 
				COALESCE(__v."alias", __t."alias"), 
				COALESCE(__v."stocked", __t."stocked"), 
				COALESCE(__v."sold", __t."sold")
			FROM __v WHERE __v."id" = __t."id")
		FROM __v WHERE __v."id" = __t."id"`
	SQLReplaceReturningProductRows = SQLReplaceProductRows + SQLReturningProductRows
	SQLSaveProductRows             = `
		WITH __v AS (SELECT * FROM json_populate_recordset(null::"wise"."product", $1))
		INSERT INTO "wise"."product" AS __t ("id", "price", "name", "alias", "stocked", "sold")
		SELECT 
			COALESCE(__v."id", nextval('wise.product_id_seq'::regclass)), 
			__v."price", 
			__v."name", 
			__v."alias", 
			__v."stocked", 
			__v."sold" FROM __v
		ON CONFLICT ("id") DO UPDATE
		SET ("id", "price", "name", "alias", "stocked", "sold") = (
			SELECT "id", "price", "name", "alias", "stocked", "sold" FROM __v
			WHERE __v."id" = __t."id"
		)`
	SQLSaveReturningProductRows = SQLSaveProductRows + SQLReturningProductRows
	SQLGetProductRowsByID       = `
		WITH __key AS (SELECT DISTINCT "id" FROM json_populate_recordset(null::"wise"."product", $1))
		SELECT "id", "price", "name", "alias", "stocked", "sold"
		FROM __key JOIN "wise"."product" AS __t USING ("id")`
	SQLUpdateProductRowsByID = `
		WITH __v AS (SELECT * FROM json_populate_record(null::"wise"."product", $1)),
		  __key AS (SELECT id FROM json_populate_recordset(null::"wise"."product", $2))
		UPDATE "wise"."product" AS __t
		SET ("id", "price", "name", "alias", "stocked", "sold") = (SELECT
			COALESCE(__v."id", __t."id"), 
			COALESCE(__v."price", __t."price"), 
			COALESCE(__v."name", __t."name"), 
			COALESCE(__v."alias", __t."alias"), 
			COALESCE(__v."stocked", __t."stocked"), 
			COALESCE(__v."sold", __t."sold")
		FROM __v)
		FROM __key WHERE (__key."id" = __t."id")`
	SQLDeleteProductRowsByID = `
		WITH __key AS (SELECT id FROM json_populate_recordset(null::"wise"."product", $1))
		DELETE FROM "wise"."product" AS __t USING __key WHERE (__key."id" = __t."id")`
	SQLDeleteReturningProductRowsByID = SQLDeleteProductRowsByID + SQLReturningProductRows
	SQLGetProductRowsByAlias          = `
		WITH __key AS (SELECT DISTINCT "alias" FROM json_populate_recordset(null::"wise"."product", $1))
		SELECT "id", "price", "name", "alias", "stocked", "sold"
		FROM __key JOIN "wise"."product" AS __t USING ("alias")`
	SQLUpdateProductRowsByAlias = `
		WITH __v AS (SELECT * FROM json_populate_record(null::"wise"."product", $1)),
		  __key AS (SELECT alias FROM json_populate_recordset(null::"wise"."product", $2))
		UPDATE "wise"."product" AS __t
		SET ("id", "price", "name", "alias", "stocked", "sold") = (SELECT
			COALESCE(__v."id", __t."id"), 
			COALESCE(__v."price", __t."price"), 
			COALESCE(__v."name", __t."name"), 
			COALESCE(__v."alias", __t."alias"), 
			COALESCE(__v."stocked", __t."stocked"), 
			COALESCE(__v."sold", __t."sold")
		FROM __v)
		FROM __key WHERE (__key."alias" = __t."alias")`
	SQLDeleteProductRowsByAlias = `
		WITH __key AS (SELECT alias FROM json_populate_recordset(null::"wise"."product", $1))
		DELETE FROM "wise"."product" AS __t USING __key WHERE (__key."alias" = __t."alias")`
	SQLDeleteReturningProductRowsByAlias = SQLDeleteProductRowsByAlias + SQLReturningProductRows
)