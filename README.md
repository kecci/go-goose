# go-goose

## Environment
DBSTRING=postgres://username:password@host:port/dbname?sslmode=disable

## Create new file
```
go run ./cmd/main.go -dir migrations create table_payment_transactions go
```

## Check Current Version
```
go run ./cmd/main.go -dir migrations postgres '[DBSTRING]' version
```

## Migrate Up
```
go run ./cmd/main.go -dir migrations postgres '[DBSTRING]' up
go run ./cmd/main.go -dir migrations postgres '[DBSTRING]' up-to 20221017133145
```

## Migrate Down
```
go run ./cmd/main.go -dir migrations postgres '[DBSTRING]' down
```

## Rebuild Binary (Optional,If you want to using binary, should rebuild every time migrate)
```
go build -o goose ./cmd/main.go
```