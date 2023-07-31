// Package public provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package public

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
	// Get a list of unconfirmed transactions currently in the transaction pool by address.
	// (GET /v2/accounts/{address}/transactions/pending)
	GetPendingTransactionsByAddress(ctx echo.Context, address string, params GetPendingTransactionsByAddressParams) error
	// Broadcasts a raw transaction or transaction group to the network.
	// (POST /v2/transactions)
	RawTransaction(ctx echo.Context) error
	// Get a list of unconfirmed transactions currently in the transaction pool.
	// (GET /v2/transactions/pending)
	GetPendingTransactions(ctx echo.Context, params GetPendingTransactionsParams) error
	// Get a specific pending transaction.
	// (GET /v2/transactions/pending/{txid})
	PendingTransactionInformation(ctx echo.Context, txid string, params PendingTransactionInformationParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetPendingTransactionsByAddress converts echo context to params.
func (w *ServerInterfaceWrapper) GetPendingTransactionsByAddress(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameterWithLocation("simple", false, "address", runtime.ParamLocationPath, ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPendingTransactionsByAddressParams
	// ------------- Optional query parameter "max" -------------

	err = runtime.BindQueryParameter("form", true, false, "max", ctx.QueryParams(), &params.Max)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter max: %s", err))
	}

	// ------------- Optional query parameter "format" -------------

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPendingTransactionsByAddress(ctx, address, params)
	return err
}

// RawTransaction converts echo context to params.
func (w *ServerInterfaceWrapper) RawTransaction(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.RawTransaction(ctx)
	return err
}

// GetPendingTransactions converts echo context to params.
func (w *ServerInterfaceWrapper) GetPendingTransactions(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPendingTransactionsParams
	// ------------- Optional query parameter "max" -------------

	err = runtime.BindQueryParameter("form", true, false, "max", ctx.QueryParams(), &params.Max)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter max: %s", err))
	}

	// ------------- Optional query parameter "format" -------------

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPendingTransactions(ctx, params)
	return err
}

// PendingTransactionInformation converts echo context to params.
func (w *ServerInterfaceWrapper) PendingTransactionInformation(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "txid" -------------
	var txid string

	err = runtime.BindStyledParameterWithLocation("simple", false, "txid", runtime.ParamLocationPath, ctx.Param("txid"), &txid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter txid: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params PendingTransactionInformationParams
	// ------------- Optional query parameter "format" -------------

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PendingTransactionInformation(ctx, txid, params)
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

	router.GET(baseURL+"/v2/accounts/:address/transactions/pending", wrapper.GetPendingTransactionsByAddress, m...)
	router.POST(baseURL+"/v2/transactions", wrapper.RawTransaction, m...)
	router.GET(baseURL+"/v2/transactions/pending", wrapper.GetPendingTransactions, m...)
	router.GET(baseURL+"/v2/transactions/pending/:txid", wrapper.PendingTransactionInformation, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"AAA=",
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
