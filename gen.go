package gen

//go:generate go run ./cmd/tools/terndotenv/main.go
//go:generate go sqlc generate -f ./internal/store/pgstore/sqlc.yaml
