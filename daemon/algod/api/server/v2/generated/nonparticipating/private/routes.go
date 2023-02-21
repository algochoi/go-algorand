// Package private provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package private

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
	// Aborts a catchpoint catchup.
	// (DELETE /v2/catchup/{catchpoint})
	AbortCatchup(ctx echo.Context, catchpoint string) error
	// Starts a catchpoint catchup.
	// (POST /v2/catchup/{catchpoint})
	StartCatchup(ctx echo.Context, catchpoint string) error

	// (POST /v2/shutdown)
	ShutdownNode(ctx echo.Context, params ShutdownNodeParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AbortCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) AbortCatchup(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameterWithLocation("simple", false, "catchpoint", runtime.ParamLocationPath, ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AbortCatchup(ctx, catchpoint)
	return err
}

// StartCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) StartCatchup(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameterWithLocation("simple", false, "catchpoint", runtime.ParamLocationPath, ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.StartCatchup(ctx, catchpoint)
	return err
}

// ShutdownNode converts echo context to params.
func (w *ServerInterfaceWrapper) ShutdownNode(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ShutdownNodeParams
	// ------------- Optional query parameter "timeout" -------------

	err = runtime.BindQueryParameter("form", true, false, "timeout", ctx.QueryParams(), &params.Timeout)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter timeout: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ShutdownNode(ctx, params)
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

	router.DELETE(baseURL+"/v2/catchup/:catchpoint", wrapper.AbortCatchup, m...)
	router.POST(baseURL+"/v2/catchup/:catchpoint", wrapper.StartCatchup, m...)
	router.POST(baseURL+"/v2/shutdown", wrapper.ShutdownNode, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"Ha90DzS9AAA=",
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
