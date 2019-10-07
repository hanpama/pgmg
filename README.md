# PGMG

Introspection based PostgreSQL model generator

* Select records by keys
* Insert values
* Update with changeset

```
usage: pgmg <connection_string> <schema_name> <outdir>

See https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING
for more information about connection string parameters.

example: go run github.com/hanpama/pgmg 'user=postgres dbname=pgmg sslmode=disable' public public
```

## Example

* [Northwind Database](https://github.com/hanpama/pgmg/tree/master/example/northwind)