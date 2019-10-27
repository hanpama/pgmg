# PGMG

> This library is in development. The API is not stable and may change.

Introspection based PostgreSQL model generator

```
usage: pgmg table <connection_string> <schema_name> <outdir>

See https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING
for more information about connection string parameters.

Tables generation:
  go run github.com/hanpama/pgmg table 'user=postgres dbname=pgmg sslmode=disable' public public
```

## Example

* https://github.com/hanpama/pgmg/tree/master/example/wise
