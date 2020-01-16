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
CREATE TABLE wise.product (
  id SERIAL PRIMARY KEY,
  price NUMERIC NOT NULL CHECK(price > 0),
  stocked TIMESTAMPTZ NOT NULL,
  sold TIMESTAMPTZ
);
```

```go
// Product represents a row for table "product"
type Product struct {
	ID      *int32     `json:"id"`
	Price   float64    `json:"price"`
	Stocked time.Time  `json:"stocked"`
	Sold    *time.Time `json:"sold"`
}
```

### Multi-row `Get`, `Save` and `Delete` methods

```go
// NewProductRepository creates a new ProductRepository
func NewProductRepository(db PGMGDatabase) *ProductRepository {
	return &ProductRepository{db}
}

// ProductRepository gets, saves and deletes rows of table "product"
type ProductRepository struct {
	db PGMGDatabase
}

// GetByProductPkey gets matching rows for given ProductPkey keys from table "product"
func (rep *ProductRepository) GetByProductPkey(ctx context.Context, keys ...ProductPkey) (rows []*Product, err error) {
	return GetByProductPkey(ctx, rep.db, keys...)
}

// SaveByProductPkey upserts the given rows for table "product" checking uniqueness by contstraint "product_pkey"
func (rep *ProductRepository) SaveByProductPkey(ctx context.Context, rows ...*Product) error {
	return SaveByProductPkey(ctx, rep.db, rows...)
}

// SaveAndReturnByProductPkey upserts the given rows for table "product" checking uniqueness by contstraint "product_pkey"
// It returns the new values and scan them into given row references.
func (rep *ProductRepository) SaveAndReturnByProductPkey(ctx context.Context, rows ...*Product) ([]*Product, error) {
	return SaveAndReturnByProductPkey(ctx, rep.db, rows...)
}

// DeleteByProductPkey deletes matching rows by ProductPkey keys from table "product"
func (rep *ProductRepository) DeleteByProductPkey(ctx context.Context, keys ...ProductPkey) (int64, error) {
	return DeleteByProductPkey(ctx, rep.db, keys...)
}
```

### `Count`, `Find` and `Delete` methods

```go
// ProductCondition is used for quering table "product"
type ProductCondition struct {
	ID      *int32     `json:"id"`
	Price   *float64   `json:"price"`
	Stocked *time.Time `json:"stocked"`
	Sold    *time.Time `json:"sold"`
}

// FindProductRows find the rows matching the condition from table "product"
func (rep *ProductRepository) FindProductRows(ctx context.Context, cond ProductCondition) ([]*Product, error) {
	return FindProductRows(ctx, rep.db, cond)
}

// DeleteProductRows delete the rows matching the condition from table "product"
func (rep *ProductRepository) DeleteProductRows(ctx context.Context, cond ProductCondition) (afftected int64, err error) {
	return DeleteProductRows(ctx, rep.db, cond)
}

// CountProductRows counts the number of rows matching the condition from table "product"
func (rep *ProductRepository) CountProductRows(ctx context.Context, cond ProductCondition) (int, error) {
	return CountProductRows(ctx, rep.db, cond)
}
```

### Row constructors for input type safety

You can optionally use the row constructors which force you specify all the columns in proper order.
It can help you statically check the errors that occur while your database schema changes.

```go
// ProductID represents column "id" of table "product"
type ProductID *int32

// ProductPrice represents column "price" of table "product"
type ProductPrice float64

// ProductStocked represents column "stocked" of table "product"
type ProductStocked time.Time

// ProductSold represents column "sold" of table "product"
type ProductSold *time.Time

// NewProduct creates a new row for table "product" with all column values
func NewProduct(
	id ProductID,
	price ProductPrice,
	stocked ProductStocked,
	sold ProductSold,
) *Product {
	return &Product{
		(*int32)(id),
		(float64)(price),
		(time.Time)(stocked),
		(*time.Time)(sold),
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
