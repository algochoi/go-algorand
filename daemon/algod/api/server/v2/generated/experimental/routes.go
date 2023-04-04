// Package experimental provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package experimental

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	. "github.com/algorand/go-algorand/daemon/algod/api/server/v2/generated/model"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Returns OK if experimental API is enabled.
	// (GET /v2/experimental)
	ExperimentalCheck(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// ExperimentalCheck converts echo context to params.
func (w *ServerInterfaceWrapper) ExperimentalCheck(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ExperimentalCheck(ctx)
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

	router.GET(baseURL+"/v2/experimental", wrapper.ExperimentalCheck, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
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
