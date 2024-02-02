// Package stock provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package stock

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// BadRequestResponse defines model for BadRequestResponse.
type BadRequestResponse struct {
	Message string `json:"message"`
}

// NewStockItem defines model for NewStockItem.
type NewStockItem struct {
	Name string `json:"name" validate:"required,lt=100"`
}

// BadRequest defines model for BadRequest.
type BadRequest = BadRequestResponse

// Created defines model for Created.
type Created struct {
	Id openapi_types.UUID `json:"id" validate:"required"`
}

// PostStockItemJSONRequestBody defines body for PostStockItem for application/json ContentType.
type PostStockItemJSONRequestBody = NewStockItem

// PutStockItemJSONRequestBody defines body for PutStockItem for application/json ContentType.
type PutStockItemJSONRequestBody = NewStockItem

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create Stock Item
	// (POST /stock/items)
	PostStockItem(ctx echo.Context) error
	// Delete Stock Item
	// (DELETE /stock/items/{stockitemId})
	DeleteStockItem(ctx echo.Context, stockitemId openapi_types.UUID) error
	// Update Stock Item
	// (PUT /stock/items/{stockitemId})
	PutStockItem(ctx echo.Context, stockitemId openapi_types.UUID) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostStockItem converts echo context to params.
func (w *ServerInterfaceWrapper) PostStockItem(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostStockItem(ctx)
	return err
}

// DeleteStockItem converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteStockItem(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "stockitemId" -------------
	var stockitemId openapi_types.UUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "stockitemId", runtime.ParamLocationPath, ctx.Param("stockitemId"), &stockitemId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter stockitemId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteStockItem(ctx, stockitemId)
	return err
}

// PutStockItem converts echo context to params.
func (w *ServerInterfaceWrapper) PutStockItem(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "stockitemId" -------------
	var stockitemId openapi_types.UUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "stockitemId", runtime.ParamLocationPath, ctx.Param("stockitemId"), &stockitemId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter stockitemId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PutStockItem(ctx, stockitemId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/stock/items", wrapper.PostStockItem)
	router.DELETE(baseURL+"/stock/items/:stockitemId", wrapper.DeleteStockItem)
	router.PUT(baseURL+"/stock/items/:stockitemId", wrapper.PutStockItem)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8yUSU8jOxDHv0qr3js63Z2Ed2npHYBhpIgRINCcUA5Wd5F4pr1gVxNQ5O8+sp2VzrBo",
	"kJibl9r8/7lqCbWWRitU5KBagkVntHIYNye8ucb7Dh2FXa0VoYpLbkwrak5Cq+KH0yqcuXqOkofVvxbv",
	"oIJ/im3oIt26YhvyepUJvPcMGnS1FSZEhCokztaZPYNTi5yweVcRxmqDlkR6iYjOd9pKTlBB14kGGNCT",
	"QajAkRVqBgweB5obMah1gzNUA3wkywfEZzHEA29Fwyk4WLzvhMUmlb7ZVbchz/TAe9YP8AwmitAq3t6g",
	"fUB7Zq22Ifq+/dooS1ZZMvMMLjR91Z1q+i4XmrJ05RlcnvcNLs8hFLYCsY93w6InnETn+Cxe7Kv1/OVr",
	"w2moEhc3pOufE0LZD6m4jPEkfxSyk1ANy/LPWbCW/h+WZR9JTJegCHWn+8LEUrNQa3Z8NQEGrahxJUaq",
	"FY4Nr+eYjfISGHS2hQrmRKYqisVikfN4m2s7K1aurvg2OT27uDkbjPIyn5NsAxUS1OKhhA9oXaplmJd5",
	"GWy1QcWNgArGeZmPgYHhNI+PL1zwLwShjHujU38e+nLZNhfEoDb2y6SBCq60oy2mJBo6OtHN04e1+95P",
	"eIaGbIfxYGfijMrh70Ju7Iqdbjoqy9ftd8aYZ/DfW1wOdWnsnk5Kbp8O6hvud9kUy7gJ60njE6EWCfus",
	"vsTzl1gli11ahlsukdA6qG6XIEKc8EWArT/tTnZ4rjvb4ffyVPR+2mP0BgHDtAl4jl433cy0j4PTF9Qz",
	"MN2BNvlumtfapKNP1P1vaMl34H53N37WD+lzj1K46JDQ7s35Vte8nWtH1XA8GoOf+l8BAAD//6QMgeE6",
	"CQAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
