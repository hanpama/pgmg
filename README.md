# PGMG

Introspection based PostgreSQL model generator

* Select records by keys
* Insert values
* Update with changeset

```
usage: pgmg table <connection_string> <schema_name> <outdir>
	OR pgmg query <connection_string> <sql_glob>

See https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING
for more information about connection string parameters.

Tables generation:
  go run github.com/hanpama/pgmg table 'user=postgres dbname=pgmg sslmode=disable' public public

Query generation:
  go run github.com/hanpama/pgmg query 'user=postgres dbname=pgmg sslmode=disable' 'example/northwind/queries/*.sql'
```

## Example

* [Northwind Database](https://github.com/hanpama/pgmg/tree/master/example/northwind)