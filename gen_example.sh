go run github.com/hanpama/pgmg/cmd/pgmg \
  -database 'user=postgres dbname=pgmg sslmode=disable' \
  -schema wise \
  -out example/tables/ \
&& go run github.com/hanpama/pgmg/cmd/pgmg \
  -database 'user=postgres dbname=pgmg sslmode=disable' \
  -schema wise \
  -out example/schema/pgmg.go
