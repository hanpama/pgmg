# PGMG

> This package is in development. The API is not stable and may change.

PostgreSQL Model Generator

```
usage: pgmg -database <connection_string> -schema <schema_name> -out <outpath>

See https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING
for more information about connection string parameters.

1. generate a single file:
	go run github.com/hanpama/pgmg \
		-database 'user=postgres dbname=pgmg sslmode=disable' \
		-schema wise \
		-out example/schema.go # file

2. generate each file per table:
	go run github.com/hanpama/pgmg \
		-database 'user=postgres dbname=pgmg sslmode=disable' \
		-schema wise \
		-out example/ # directory

```

## PGMG Generates

* https://github.com/hanpama/pgmg/tree/master/example/tables/


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

* https://github.com/hanpama/pgmg/tree/master/example/tables/tables_test.go
