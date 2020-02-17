// Code generated by github.com/hanpama/pgmg. DO NOT EDIT.
package tables

import (
	"context"
	"encoding/json"
)

// CampaignRow represents a row for table "campaign"
type CampaignRow struct {
	Data CampaignData
}

type CampaignData struct {
	ID      string  `json:"id"`
	PopName *string `json:"pop_name"`
	PopYear *int32  `json:"pop_year"`
}

func NewCampaignRow(data CampaignData) *CampaignRow {
	return &CampaignRow{data}
}

func NewCampaignRows(data ...CampaignData) CampaignRows {
	rows := make(CampaignRows, len(data))
	for i, d := range data {
		rows[i] = NewCampaignRow(d)
	}
	return rows
}

// GetID gets value of column "id" from "campaign" row
func (r *CampaignRow) GetID() string { return r.Data.ID }

// SetID sets value of column "id" in "campaign" row
func (r *CampaignRow) SetID(id string) { r.Data.ID = id }

// GetPopName gets value of column "pop_name" from "campaign" row
func (r *CampaignRow) GetPopName() string { return *r.Data.PopName }

// SetPopName sets value of column "pop_name" in "campaign" row
func (r *CampaignRow) SetPopName(popName string) { r.Data.PopName = &popName }

// ClearPopName sets value of column "pop_name" null in "campaign" row
func (r *CampaignRow) ClearPopName() { r.Data.PopName = nil }

// HasValidPopName checks to value of column "pop_name" is not null
func (r *CampaignRow) HasValidPopName() bool { return r.Data.PopName != nil }

// GetPopYear gets value of column "pop_year" from "campaign" row
func (r *CampaignRow) GetPopYear() int32 { return *r.Data.PopYear }

// SetPopYear sets value of column "pop_year" in "campaign" row
func (r *CampaignRow) SetPopYear(popYear int32) { r.Data.PopYear = &popYear }

// ClearPopYear sets value of column "pop_year" null in "campaign" row
func (r *CampaignRow) ClearPopYear() { r.Data.PopYear = nil }

// HasValidPopYear checks to value of column "pop_year" is not null
func (r *CampaignRow) HasValidPopYear() bool { return r.Data.PopYear != nil }

// CampaignID represents key defined by PRIMARY KEY constraint "campaign_pkey" for table "campaign"
type CampaignID struct {
	ID string `json:"id"`
}

func (r *CampaignRow) KeyID() CampaignID {
	return CampaignID{r.GetID()}
}

// CampaignRows represents multiple rows for table "campaign"
type CampaignRows []*CampaignRow

func (rs CampaignRows) KeyID() (keys Keys) {
	keys = make(Keys, len(rs))
	for i, r := range rs {
		keys[i] = r.KeyID()
	}
	return keys
}

func (r *CampaignRow) RefPopNamePopYear() PopNameYear {
	return PopNameYear{r.GetPopName(), r.GetPopYear()}
}

func (rs CampaignRows) RefPopNamePopYear() (keys Keys) {
	keys = make(Keys, len(rs))
	for i, r := range rs {
		if !r.HasValidPopName() {
			continue
		}
		if !r.HasValidPopYear() {
			continue
		}
		keys[i] = r.RefPopNamePopYear()
	}
	return keys
}

// NewCampaignTable(h SQLHandle) creates new CampaignTable
func NewCampaignTable(h SQLHandle) *CampaignTable {
	return &CampaignTable{h}
}

// CampaignTable provides access methods for table "campaign"
type CampaignTable struct {
	h SQLHandle
}

func (t *CampaignTable) Find(ctx context.Context, filter CampaignValues) (CampaignRows, error) {
	return FindCampaignRows(ctx, t.h, filter)
}

func (t *CampaignTable) Count(ctx context.Context, filter CampaignValues) (int, error) {
	return CountCampaignRows(ctx, t.h, filter)
}

func (t *CampaignTable) Update(ctx context.Context, changeset, filter CampaignValues) (int64, error) {
	return UpdateCampaignRows(ctx, t.h, changeset, filter)
}

func (t *CampaignTable) Insert(ctx context.Context, rows ...*CampaignRow) (int, error) {
	return InsertReturningCampaignRows(ctx, t.h, rows...)
}

func (t *CampaignTable) Delete(ctx context.Context, filter CampaignValues) (int64, error) {
	return DeleteCampaignRows(ctx, t.h, filter)
}

func (t *CampaignTable) Save(ctx context.Context, rows ...*CampaignRow) error {
	return SaveReturningCampaignRows(ctx, t.h, rows...)
}

func (t *CampaignTable) GetByID(ctx context.Context, keys ...interface{}) (CampaignRows, error) {
	return GetCampaignRowsByID(ctx, t.h, keys...)
}

func (t *CampaignTable) UpdateByID(ctx context.Context, changeset CampaignValues, keys ...interface{}) (int64, error) {
	return UpdateCampaignRowsByID(ctx, t.h, changeset, keys...)
}

func (t *CampaignTable) DeleteByID(ctx context.Context, keys ...interface{}) (int64, error) {
	return DeleteCampaignRowsByID(ctx, t.h, keys...)
}

type CampaignValues struct {
	ID      *string `json:"id"`
	PopName *string `json:"pop_name"`
	PopYear *int32  `json:"pop_year"`
}

// InsertCampaignRows inserts the rows into table "campaign"
func InsertCampaignRows(ctx context.Context, db SQLHandle, rows ...*CampaignRow) (numRows int64, err error) {
	numRows, err = execWithJSONArgs(ctx, db, SQLInsertCampaignRows, rows)
	if err != nil {
		return numRows, formatError("InsertCampaignRows", err)
	}
	return numRows, nil
}

// InsertReturningCampaignRows inserts the rows into table "campaign" and returns the rows.
func InsertReturningCampaignRows(ctx context.Context, db SQLHandle, inputs ...*CampaignRow) (numRows int, err error) {
	rows := CampaignRows(inputs)
	numRows, err = queryWithJSONArgs(ctx, db, rows.ReceiveRows, SQLInsertReturningCampaignRows, rows)
	if err != nil {
		return numRows, formatError("InsertReturningCampaignRows", err)
	}
	return numRows, nil
}

// FindCampaignRows finds the rows matching the condition from table "campaign"
func FindCampaignRows(ctx context.Context, db SQLHandle, cond CampaignValues) (rows CampaignRows, err error) {
	if _, err = queryWithJSONArgs(ctx, db, rows.ReceiveRows, SQLFindCampaignRows, cond); err != nil {
		return nil, formatError("FindCampaignRows", err)
	}
	return rows, nil
}

// DeleteCampaignRows deletes the rows matching the condition from table "campaign"
func DeleteCampaignRows(ctx context.Context, db SQLHandle, cond CampaignValues) (numRows int64, err error) {
	if numRows, err = execWithJSONArgs(ctx, db, SQLDeleteCampaignRows, cond); err != nil {
		return numRows, formatError("DeleteCampaignRows", err)
	}
	return numRows, nil
}

func UpdateCampaignRows(ctx context.Context, db SQLHandle, changeset, filter CampaignValues) (numRows int64, err error) {
	if numRows, err = execWithJSONArgs(ctx, db, SQLUpdateCampaignRows, changeset, filter); err != nil {
		return numRows, formatError("UpdateCampaignRows", err)
	}
	return numRows, nil
}

// CountCampaignRows counts the number of rows matching the condition from table "campaign"
func CountCampaignRows(ctx context.Context, db SQLHandle, cond CampaignValues) (count int, err error) {
	if _, err = queryWithJSONArgs(ctx, db, func(int) []interface{} { return []interface{}{&count} }, SQLCountCampaignRows, cond); err != nil {
		return 0, formatError("CountCampaignRows", err)
	}
	return count, nil
}

// SaveCampaignRows upserts the given rows for table "campaign" checking uniqueness by contstraint "campaign_pkey"
func SaveCampaignRows(ctx context.Context, db SQLHandle, rows ...*CampaignRow) (err error) {
	_, err = execWithJSONArgs(ctx, db, SQLSaveCampaignRows, rows)
	if err != nil {
		return formatError("SaveCampaignRows", err)
	}
	return nil
}

// SaveReturningCampaignRows upserts the given rows for table "campaign" checking uniqueness by contstraint "campaign_pkey"
// It returns the new values and scan them into given row references.
func SaveReturningCampaignRows(ctx context.Context, db SQLHandle, inputs ...*CampaignRow) (err error) {
	rows := CampaignRows(inputs)
	_, err = queryWithJSONArgs(ctx, db, rows.ReceiveRows, SQLSaveReturningCampaignRows, rows)
	if err != nil {
		return formatError("SaveReturningCampaignRows", err)
	}
	return nil
}

// GetCampaignRowsByID gets matching rows for given ID keys from table "campaign"
func GetCampaignRowsByID(ctx context.Context, db SQLHandle, keys ...interface{}) (rows CampaignRows, err error) {
	rows = make(CampaignRows, 0, len(keys))
	if _, err = queryWithJSONArgs(ctx, db, rows.ReceiveRows, SQLGetCampaignRowsByID, Keys(keys)); err != nil {
		return nil, formatError("GetCampaignRowsByID", err)
	}
	return rows, nil
}

// DeleteCampaignRowsByID deletes matching rows by CampaignID keys from table "campaign"
func DeleteCampaignRowsByID(ctx context.Context, db SQLHandle, keys ...interface{}) (numRows int64, err error) {
	numRows, err = execWithJSONArgs(ctx, db, SQLDeleteCampaignRowsByID, keys)
	if err != nil {
		return numRows, formatError("DeleteCampaignRowsByID", err)
	}
	return numRows, nil
}

// UpdateCampaignRowsByID deletes matching rows by CampaignID keys from table "campaign"
func UpdateCampaignRowsByID(ctx context.Context, db SQLHandle, changeset CampaignValues, keys ...interface{}) (numRows int64, err error) {
	numRows, err = execWithJSONArgs(ctx, db, SQLUpdateCampaignRowsByID, changeset, keys)
	if err != nil {
		return numRows, formatError("UpdateCampaignRowsByID", err)
	}
	return numRows, nil
}

// ReceiveRow returns all pointers of the column values for scanning
func (r *CampaignRow) ReceiveRow() []interface{} {
	return []interface{}{&r.Data.ID, &r.Data.PopName, &r.Data.PopYear}
}

// ReceiveRows returns pointer slice to receive data for the row on index i
func (rs *CampaignRows) ReceiveRows(i int) []interface{} {
	if len(*rs) <= i {
		*rs = append(*rs, new(CampaignRow))
	} else if (*rs)[i] == nil {
		(*rs)[i] = new(CampaignRow)
	}
	return (*rs)[i].ReceiveRow()
}

func (r *CampaignRow) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Data)
}

var (
	SQLFindCampaignRows = `
		WITH __f AS (SELECT "id", "pop_name", "pop_year" FROM json_populate_record(null::"wise"."campaign", $1))
		SELECT __t.id, __t.pop_name, __t.pop_year
		FROM "wise"."campaign" AS __t
		WHERE ((SELECT __f."id" IS NULL FROM __f) OR (SELECT __f."id" = __t."id" FROM __f))
			AND ((SELECT __f."pop_name" IS NULL FROM __f) OR (SELECT __f."pop_name" = __t."pop_name" FROM __f))
			AND ((SELECT __f."pop_year" IS NULL FROM __f) OR (SELECT __f."pop_year" = __t."pop_year" FROM __f))`
	SQLCountCampaignRows = `
		WITH __f AS (SELECT "id", "pop_name", "pop_year" FROM json_populate_record(null::"wise"."campaign", $1))
		SELECT count(*) FROM "wise"."campaign" AS __t
		WHERE ((SELECT __f."id" IS NULL FROM __f) OR (SELECT __f."id" = __t."id" FROM __f))
			AND ((SELECT __f."pop_name" IS NULL FROM __f) OR (SELECT __f."pop_name" = __t."pop_name" FROM __f))
			AND ((SELECT __f."pop_year" IS NULL FROM __f) OR (SELECT __f."pop_year" = __t."pop_year" FROM __f))`
	SQLReturningCampaignRows = `
		RETURNING "id", "pop_name", "pop_year"`
	SQLInsertCampaignRows = `
		WITH __v AS (SELECT * FROM json_populate_recordset(null::"wise"."campaign", $1))
		INSERT INTO "wise"."campaign" AS __t ("id", "pop_name", "pop_year")
		SELECT 
			__v."id", 
			__v."pop_name", 
			__v."pop_year" FROM __v`
	SQLInsertReturningCampaignRows = SQLInsertCampaignRows + SQLReturningCampaignRows
	SQLDeleteCampaignRows          = `
		DELETE FROM "wise"."campaign" AS __t
		WHERE TRUE
			AND (($1::json->>'id' IS NULL) OR CAST($1::json->>'id' AS uuid) = __t."id")
			AND (($1::json->>'pop_name' IS NULL) OR CAST($1::json->>'pop_name' AS text) = __t."pop_name")
			AND (($1::json->>'pop_year' IS NULL) OR CAST($1::json->>'pop_year' AS integer) = __t."pop_year")`
	SQLDeleteReturningCampaignRows = SQLDeleteCampaignRows + SQLReturningCampaignRows
	SQLUpdateCampaignRows          = `
		WITH __v AS (SELECT * FROM json_populate_record(null::"wise"."campaign", $1)),
			__f AS (SELECT * FROM json_populate_record(null::"wise"."campaign", $2))
		UPDATE "wise"."campaign" AS __t
		SET ("id", "pop_name", "pop_year") = (SELECT 
			COALESCE(__v."id", __t."id"), 
			COALESCE(__v."pop_name", __t."pop_name"), 
			COALESCE(__v."pop_year", __t."pop_year") FROM __v)
		WHERE ((SELECT __f."id" IS NULL FROM __f) OR (SELECT __f."id" = __t."id" FROM __f))
			AND ((SELECT __f."pop_name" IS NULL FROM __f) OR (SELECT __f."pop_name" = __t."pop_name" FROM __f))
			AND ((SELECT __f."pop_year" IS NULL FROM __f) OR (SELECT __f."pop_year" = __t."pop_year" FROM __f))`
	SQLUpdateReturningCampaignRows = SQLUpdateCampaignRows + SQLReturningCampaignRows
	SQLReplaceCampaignRows         = `
		WITH __v AS (SELECT * FROM json_populate_recordset(null::"wise"."campaign", $1))
		UPDATE "wise"."campaign" AS __t
			SET ("id", "pop_name", "pop_year") = (SELECT 
				COALESCE(__v."id", __t."id"), 
				COALESCE(__v."pop_name", __t."pop_name"), 
				COALESCE(__v."pop_year", __t."pop_year")
			FROM __v WHERE __v."id" = __t."id")
		FROM __v WHERE __v."id" = __t."id"`
	SQLReplaceReturningCampaignRows = SQLReplaceCampaignRows + SQLReturningCampaignRows
	SQLSaveCampaignRows             = `
		WITH __v AS (SELECT * FROM json_populate_recordset(null::"wise"."campaign", $1))
		INSERT INTO "wise"."campaign" AS __t ("id", "pop_name", "pop_year")
		SELECT 
			__v."id", 
			__v."pop_name", 
			__v."pop_year" FROM __v
		ON CONFLICT ("id") DO UPDATE
		SET ("id", "pop_name", "pop_year") = (
			SELECT "id", "pop_name", "pop_year" FROM __v
			WHERE __v."id" = __t."id"
		)`
	SQLSaveReturningCampaignRows = SQLSaveCampaignRows + SQLReturningCampaignRows
	SQLGetCampaignRowsByID       = `
		WITH __key AS (SELECT DISTINCT "id" FROM json_populate_recordset(null::"wise"."campaign", $1))
		SELECT "id", "pop_name", "pop_year"
		FROM __key JOIN "wise"."campaign" AS __t USING ("id")`
	SQLUpdateCampaignRowsByID = `
		WITH __v AS (SELECT * FROM json_populate_record(null::"wise"."campaign", $1)),
		  __key AS (SELECT id FROM json_populate_recordset(null::"wise"."campaign", $2))
		UPDATE "wise"."campaign" AS __t
		SET ("id", "pop_name", "pop_year") = (SELECT
			COALESCE(__v."id", __t."id"), 
			COALESCE(__v."pop_name", __t."pop_name"), 
			COALESCE(__v."pop_year", __t."pop_year")
		FROM __v)
		FROM __key WHERE (__key."id" = __t."id")`
	SQLDeleteCampaignRowsByID = `
		WITH __key AS (SELECT id FROM json_populate_recordset(null::"wise"."campaign", $1))
		DELETE FROM "wise"."campaign" AS __t USING __key WHERE (__key."id" = __t."id")`
	SQLDeleteReturningCampaignRowsByID = SQLDeleteCampaignRowsByID + SQLReturningCampaignRows
)