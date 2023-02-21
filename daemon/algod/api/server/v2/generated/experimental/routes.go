// Package experimental provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package experimental

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	. "github.com/algorand/go-algorand/daemon/algod/api/server/v2/generated/model"
	"github.com/algorand/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Simulates a raw transaction or transaction group as it would be evaluated on the network. WARNING: This endpoint is experimental and under active development. There are no guarantees in terms of functionality or future support.
	// (POST /v2/transactions/simulate)
	SimulateTransaction(ctx echo.Context, params SimulateTransactionParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// SimulateTransaction converts echo context to params.
func (w *ServerInterfaceWrapper) SimulateTransaction(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params SimulateTransactionParams
	// ------------- Optional query parameter "format" -------------

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.SimulateTransaction(ctx, params)
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
func RegisterHandlers(router EchoRouter, si ServerInterface, m ...echo.MiddlewareFunc) {
	RegisterHandlersWithBaseURL(router, si, "", m...)
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/v2/transactions/simulate", wrapper.SimulateTransaction, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"atR9aZtZKdkVZsp9+PR/AwAA//87chVMC74AAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
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
	var res = make(map[string]func() ([]byte, error))
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
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
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
