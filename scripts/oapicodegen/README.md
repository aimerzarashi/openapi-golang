~/go/bin/oapi-codegen -generate types,server,spec -package oapicodegen ./../../api/hello.yaml > ./../../internal/infrastructure/oapicodegen/hello/hello.go && go mod tidy

~/go/bin/oapi-codegen -generate types,server,spec -package oapicodegen ./../../api/stock.yaml > ./../../internal/infrastructure/oapicodegen/stock/stock.go && go mod tidy