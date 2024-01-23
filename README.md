# openapi-golang

## migration

### creates
```
docker run -v ./db/migrations:/migrations migrate/migrate -path=/migrations/ -database postgresql://user:password@openapi-db:5432/openapi?sslmode=disable create -ext sql -dir ./migrations -seq stock_item

sudo chown -R k8suser:k8suser db
```

### up
```
docker run --rm -v ./db/migrations:/migrations --network openapi-golang_default migrate/migrate -path=/migrations -database postgresql://user:password@openapi-db:5432/openapi?sslmode=disable up
```

## down
```
 docker run --rm -v ./db/migrations:/migrations --network openapi-golang_default migrate/migrate -path=/migrations -database postgresql://user:password@openapi-db:5432/openapi?sslmode=disable down
```

## sqlboiler
```
docker run --rm -v ./:/sqlboiler --network openapi-golang_default curvegrid/sqlboiler:psql psql --output ./internal/datastore --pkgname datastore --wipe && go mod tidy
```

## oapi-codegen
```
~/go/bin/oapi-codegen -generate server,types -package hello ./internal/presentation/hello.yaml > ./internal/presentation/hello.gen.go && go mod tidy
~/go/bin/oapi-codegen -generate server,types -package stock_item ./internal/presentation/stock_item.yaml > ./internal/presentation/stock_item.gen.go && go mod tidy
```