# PGMG

> This library is in development. The API is not stable and may change.

PostgreSQL Model Generator

```
usage: pgmg -database <connection_string> -schema <schema_name> -out <outfile>

See https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING
for more information about connection string parameters.

example:
	go run github.com/hanpama/pgmg \
		-database 'user=postgres dbname=pgmg sslmode=disable' \
		-schema wise \
		-out example/schema.go
```

## PGMG Generates

* structs that represents rows for the tables
* multi-row `Get`, `Save` and `Delete` methods by primary key and unique keys
* `Count`, `Find` and `Delete` methods per table with its EQ filter
* and row constructors for input type safety


### Structs represents rows for the tables

```sql
CREATE TABLE wise.semester (
  id SERIAL PRIMARY KEY,
  year INTEGER NOT NULL,
  season TEXT NOT NULL,

  UNIQUE (year, season)
);
```

```go
// Semester represents a row for table "semester"
type Semester struct {
	ID     *int32 `json:"id"`
	Year   int32  `json:"year"`
	Season string `json:"season"`
}
```

### Multi-row `Get`, `Save` and `Delete` methods per key

```go
// SemesterPkey represents the key defined by UNIQUE constraint "semester_pkey" for table "semester"
type SemesterPkey struct {
	ID int32 `json:"id"`
}

func (r *Semester) SemesterPkey() SemesterPkey {
	k := SemesterPkey{}
	if r.ID != nil {
		k.ID = *r.ID
	}
	return k
}

func (rs SemesterRows) SemesterPkeySlice() (keys []SemesterPkey) {
	keys = make([]SemesterPkey, len(rs))
	for i, r := range rs {
		keys[i] = r.SemesterPkey()
	}
	return keys
}

var SQLGetBySemesterPkey = `
	WITH __key AS (
		SELECT ROW_NUMBER() over () __keyindex,
			id
		FROM json_populate_recordset(null::"wise"."semester", $1)
	)
	SELECT "id", "year", "season"
	FROM __key JOIN "wise"."semester" AS __table USING ("id")
	ORDER BY __keyindex
`

// GetBySemesterPkey gets matching rows for given SemesterPkey keys from table "semester"
func GetBySemesterPkey(ctx context.Context, db PGMGDatabase, keys ...SemesterPkey) (rows SemesterRows, err error) {
	var b []byte
	if b, err = json.Marshal(keys); err != nil {
		return nil, fmt.Errorf("%w(GetBySemesterPkey, %w)", ErrPGMG, err)
	}
	rows = make(SemesterRows, len(keys))
	if _, err = db.QueryScan(ctx, func(i int) []interface{} {
		rows[i] = &Semester{}
		return rows[i].ReceiveRow()
	}, SQLGetBySemesterPkey, b); err != nil {
		return nil, fmt.Errorf("%w(GetBySemesterPkey, %w)", ErrPGMG, err)
	}
	for i := 0; i < len(keys); i++ {
		if rows[i] == nil {
			break
		} else if rows[i].SemesterPkey() != keys[i] {
			copy(rows[i+1:], rows[i:])
			rows[i] = nil
		}
	}
	return rows, nil
}

var SQLSaveBySemesterPkey = SQLInsertSemesterRows + `
	ON CONFLICT ("id") DO UPDATE
		SET ("id", "year", "season") = (
			SELECT "id", "year", "season" FROM __values
			WHERE __values."id" = _t."id"
		)
`

// SaveBySemesterPkey upserts the given rows for table "semester" checking uniqueness by contstraint "semester_pkey"
func SaveBySemesterPkey(ctx context.Context, db PGMGDatabase, rows ...*Semester) (err error) {
	if err = execJSON(ctx, db, SQLSaveBySemesterPkey, rows, len(rows)); err != nil {
		return fmt.Errorf("%w(SaveBySemesterPkey, %w)", ErrPGMG, err)
	}
	return nil
}

var SQLSaveAndReturnBySemesterPkey = SQLSaveBySemesterPkey + sqlReturningSemesterRows

// SaveAndReturnBySemesterPkey upserts the given rows for table "semester" checking uniqueness by contstraint "semester_pkey"
// It returns the new values and scan them into given row references.
func SaveAndReturnBySemesterPkey(ctx context.Context, db PGMGDatabase, rows ...*Semester) (SemesterRows, error) {
	err := execJSONAndReturn(ctx, db, func(i int) []interface{} { return rows[i].ReceiveRow() }, SQLSaveAndReturnBySemesterPkey, rows, len(rows))
	if err != nil {
		return rows, fmt.Errorf("%w(SaveAndReturnBySemesterPkey, %w)", ErrPGMG, err)
	}
	return rows, nil
}

var SQLDeleteBySemesterPkey = `
WITH __key AS (SELECT id FROM json_populate_recordset(null::"wise"."semester", $1))
DELETE FROM "wise"."semester" AS __table
	USING __key
	WHERE (__key."id" = __table."id")
	`

// DeleteBySemesterPkey deletes matching rows by SemesterPkey keys from table "semester"
func DeleteBySemesterPkey(ctx context.Context, db PGMGDatabase, keys ...SemesterPkey) (affected int64, err error) {
	b, err := json.Marshal(keys)
	if err != nil {
		return affected, fmt.Errorf("%w(DeleteBySemesterPkey, %w)", ErrPGMG, err)
	}
	if affected, err = db.ExecCountingAffected(ctx, SQLDeleteBySemesterPkey, b); err != nil {
		return affected, fmt.Errorf("%w(DeleteBySemesterPkey, %w)", ErrPGMG, err)
	}
	return affected, nil
}
```

### `Insert`, `Find`, `Count` and `Delete` methods per table

```go
var SQLInsertSemesterRows = `
	WITH __values AS (
		SELECT
			COALESCE(__input."id", nextval('wise.semester_id_seq'::regclass)) "id",
			__input."year",
			__input."season"
		FROM json_populate_recordset(null::"wise"."semester", $1) __input
	)
	INSERT INTO "wise"."semester" AS _t ("id", "year", "season")
	SELECT "id", "year", "season" FROM __values`

func InsertSemesterRows(ctx context.Context, db PGMGDatabase, rows ...*Semester) (err error) {
	if err = execJSON(ctx, db, SQLInsertSemesterRows, rows, len(rows)); err != nil {
		return fmt.Errorf("%w( InsertSemesterRows, %w)", ErrPGMG, err)
	}
	return nil
}

var sqlReturningSemesterRows = `
	RETURNING id, year, season
`

var SQLInsertAndReturnSemesterRows = SQLInsertSemesterRows + sqlReturningSemesterRows

func InsertAndReturnSemesterRows(ctx context.Context, db PGMGDatabase, rows ...*Semester) (SemesterRows, error) {
	err := execJSONAndReturn(ctx, db, func(i int) []interface{} { return rows[i].ReceiveRow() }, SQLInsertAndReturnSemesterRows, rows, len(rows))
	if err != nil {
		return rows, fmt.Errorf("%w(SQLInsertAndReturnSemesterRows, %w)", ErrPGMG, err)
	}
	return rows, nil
}

// FindSemesterRows find the rows matching the condition from table "semester"
func FindSemesterRows(ctx context.Context, db PGMGDatabase, cond SemesterCondition) (rows SemesterRows, err error) {
	var arg1 []byte
	if arg1, err = json.Marshal(cond); err != nil {
		return nil, err
	}
	_, err = db.QueryScan(ctx, func(i int) []interface{} {
		rows = append(rows, new(Semester))
		return rows[i].ReceiveRow()
	}, `
		SELECT __t.id, __t.year, __t.season
		FROM "wise"."semester" AS __t
		WHERE TRUE
			AND (($1::json->>'id' IS NULL) OR CAST($1::json->>'id' AS integer) = __t."id")
			AND (($1::json->>'year' IS NULL) OR CAST($1::json->>'year' AS integer) = __t."year")
			AND (($1::json->>'season' IS NULL) OR CAST($1::json->>'season' AS text) = __t."season")
	`, arg1)
	return rows, err
}

// DeleteSemesterRows delete the rows matching the condition from table "semester"
func DeleteSemesterRows(ctx context.Context, db PGMGDatabase, cond SemesterCondition) (afftected int64, err error) {
	var arg1 []byte
	if arg1, err = json.Marshal(cond); err != nil {
		return 0, err
	}
	return db.ExecCountingAffected(ctx, `
		DELETE FROM "wise"."semester" AS __t
		WHERE TRUE
			AND (($1::json->>'id' IS NULL) OR CAST($1::json->>'id' AS integer) = __t."id")
			AND (($1::json->>'year' IS NULL) OR CAST($1::json->>'year' AS integer) = __t."year")
			AND (($1::json->>'season' IS NULL) OR CAST($1::json->>'season' AS text) = __t."season")
	`, arg1)
}

// CountSemesterRows counts the number of rows matching the condition from table "semester"
func CountSemesterRows(ctx context.Context, db PGMGDatabase, cond SemesterCondition) (count int, err error) {
	var arg1 []byte
	if arg1, err = json.Marshal(cond); err != nil {
		return 0, err
	}
	_, err = db.QueryScan(ctx, func(int) []interface{} { return []interface{}{&count} }, `
		SELECT count(*) FROM "wise"."semester" AS __t
		WHERE TRUE
			AND (($1::json->>'id' IS NULL) OR CAST($1::json->>'id' AS integer) = __t."id")
			AND (($1::json->>'year' IS NULL) OR CAST($1::json->>'year' AS integer) = __t."year")
			AND (($1::json->>'season' IS NULL) OR CAST($1::json->>'season' AS text) = __t."season")
	`, arg1)
	return count, err
}
```

### Row constructors for input type safety

You can optionally use the row constructors which force you specify all the columns in proper order.
It can help you statically check the errors that occur while your database schema changes.

```go
// SemesterID represents value type of column "id" of table "semester"
type SemesterID *int32

// SemesterYear represents value type of column "year" of table "semester"
type SemesterYear int32

// SemesterSeason represents value type of column "season" of table "semester"
type SemesterSeason string

// NewSemester creates a new row for table "semester" with all column values
func NewSemester(
	id SemesterID,
	year SemesterYear,
	season SemesterSeason,
) *Semester {
	return &Semester{
		(*int32)(id),
		(int32)(year),
		(string)(season),
	}
}
```

## How to use

Generated PGMG models have no concrete dependency on any SQL driver,
but depend on an interface `PGMGDatabase`.

```go
// PGMGDatabase represents PostgresQL database
type PGMGDatabase interface {
	QueryScan(ctx context.Context, receiver func(int) []interface{}, sql string, args ...interface{}) (int, error)
	Exec(ctx context.Context, sql string, args ...interface{}) (int64, error)
}
```

So your database connection should implement `PGMGDatabase` like below.

The `receiver func(int) []interface{}` parameter is a function returning pointers for `rows.Scan()`-like functions
to scan data into.

```go
type testDB struct {
	b *sql.DB
}

var _ example.PGMGDatabase = (*testDB)(nil)

func (db *testDB) QueryScan(ctx context.Context, r func(int) []interface{}, sql string, args ...interface{}) (rowsReceived int, err error) {
	rows, err := db.b.QueryContext(ctx, sql, args...)
	if err != nil {
		return rowsReceived, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(r(rowsReceived)...)
		if err != nil {
			return rowsReceived, err
		}
		rowsReceived++
	}
	return rowsReceived, err
}

func (db *testDB) Exec(ctx context.Context, sql string, args ...interface{}) (int64, error) {
	res, err := db.b.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
```

## Full Example

* https://github.com/hanpama/pgmg/tree/master/example/schema.go
* https://github.com/hanpama/pgmg/tree/master/example/example_test.go
