# openapi-golang

```
docker run -v ./db/migrations:/migrations migrate/migrate -path=/migrations/ -database postgresql://user:password@openapi-db:5432/openapi?sslmode=disable create -ext sql -dir ./migrations -seq stock_item

sudo chown -R k8suser:k8suser db
```

## up
```
docker run --rm -v ./db/migrations:/migrations --network openapi-golang_default migrate/migrate -path=/migrations -database postgresql://user:password@openapi-db:5432/openapi?sslmode=disable up
```

## down
```
 docker run --rm -v ./db/migrations:/migrations --network openapi-golang_default migrate/migrate -path=/migrations -database postgresql://user:password@openapi-db:5432/openapi?sslmode=disable down
```

## sqlboiler
```
docker run --rm -v ./app:/sqlboiler --network openapi-golang_default curvegrid/sqlboiler:psql psql --output ./repository/models --pkgname models --wipe && cd app && go mod tidy
```