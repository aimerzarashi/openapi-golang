~/go/bin/oapi-codegen -generate types,server,spec -package oapicodegen ./../../api/hello.yaml > ./../../internal/infra/oapicodegen/hello/hello.go && go mod tidy

~/go/bin/oapi-codegen -generate types,server,spec -package oapicodegen ./../../api/stockitem.yaml > ./../../internal/infra/oapicodegen/stockitem/stockitem.go && go mod tidy