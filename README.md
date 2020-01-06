# PGMG

> This library is in development. The API is not stable and may change.

PostgreSQL Model Generator

```
usage: pgmg -database <connection_string> -schema <schema_name> -out <outfile>

See https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING
for more information about connection string parameters.

Tables generation:
	go run github.com/hanpama/pgmg \
		-database 'user=postgres dbname=pgmg sslmode=disable' \
		-schema wise \
		-out example/schema.go
```

## Example

* https://github.com/hanpama/pgmg/tree/master/example/schema.go
