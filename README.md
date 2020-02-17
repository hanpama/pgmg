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
but depend on an interface `SQLHandle`.

```go
type SQLHandle interface {
	QueryAndReceive(ctx context.Context, receiver func(int) []interface{}, sql string, args ...interface{}) (int, error)
	ExecAndCount(ctx context.Context, sql string, args ...interface{}) (int64, error)
}
```

So your database connection should implement `SQLHandle` like the [test database](https://github.com/hanpama/pgmg/tree/master/example/testdb.go).

The `receiver func(int) []interface{}` parameter is a function returning pointers for `rows.Scan()`-like functions
to scan data into.

## Full Example

* https://github.com/hanpama/pgmg/tree/master/example/tables/tables_test.go
