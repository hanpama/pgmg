go run github.com/hanpama/pgmg \
  -database 'user=postgres dbname=pgmg sslmode=disable' \
  -schema wise \
  -out example/schema.go
