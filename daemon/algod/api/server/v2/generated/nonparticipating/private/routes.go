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

	"H4sIAAAAAAAC/+x9/ZPbNrLgv4LSe1WOfeLM+CPZtau23k3sJDsXJ3F5Jtl7z/ZlIbIlYYcCuASokeKb",
	"//0K3QAJkqBEzSj2pu79ZI+Ij0aj0ehu9MfHSapWhZIgjZ68+DgpeMlXYKDEv3iaqkqaRGT2rwx0WorC",
	"CCUnL/w3pk0p5GIynQj7a8HNcjKdSL6Cpo3tP52U8M9KlJBNXpiygulEp0tYcTuw2Ra2dT3SJlmoxA1x",
	"TkNcvJrc7vjAs6wErftQ/iTzLRMyzasMmCm51Dy1nzS7EWbJzFJo5jozIZmSwNScmWWrMZsLyDN94hf5",
	"zwrKbbBKN/nwkm4bEJNS5dCH86VazYQEDxXUQNUbwoxiGcyx0ZIbZmewsPqGRjENvEyXbK7KPaASECG8",
	"IKvV5MW7iQaZQYm7lYJY43/nJcBvkBheLsBMPkxji5sbKBMjVpGlXTjsl6Cr3GiGbXGNC7EGyWyvE/ZD",
	"pQ2bAeOSvf32JXv69Olzu5AVNwYyR2SDq2pmD9dE3ScvJhk34D/3aY3nC1VymSV1+7ffvsT5L90Cx7bi",
	"WkP8sJzbL+zi1dACfMcICQlpYIH70KJ+2yNyKJqfZzBXJYzcE2p81E0J5/+su5Jyky4LJaSJ7AvDr4w+",
	"R3lY0H0XD6sBaLUvLKZKO+i7s+T5h4+Pp4/Pbv/t3XnyX+7PL5/ejlz+y3rcPRiINkyrsgSZbpNFCRxP",
	"y5LLPj7eOnrQS1XlGVvyNW4+XyGrd32Z7Uusc83zytKJSEt1ni+UZtyRUQZzXuWG+YlZJXPLpuxojtqZ",
	"0Kwo1VpkkE0t971ZinTJUq5pCGzHbkSeWxqsNGRDtBZf3Y7DdBuixMJ1J3zggv51kdGsaw8mYIPcIElz",
	"pSExas/15G8cLjMWXijNXaUPu6zY1RIYTm4/0GWLuJOWpvN8ywzua8a4Zpz5q2nKxJxtVcVucHNycY39",
	"3Wos1lbMIg03p3WP2sM7hL4eMiLImymVA5eIPH/u+iiTc7GoStDsZglm6e68EnShpAamZv+A1Nht/1+X",
	"P/3IVMl+AK35At7w9JqBTFUG2Qm7mDOpTEAajpYQh7bn0DocXLFL/h9aWZpY6UXB0+v4jZ6LlYis6ge+",
	"EatqxWS1mkFpt9RfIUaxEkxVyiGAaMQ9pLjim/6kV2UlU9z/ZtqWLGepTegi51tE2Ipv/nI2deBoxvOc",
	"FSAzIRfMbOSgHGfn3g9eUqpKZiPEHGP3NLhYdQGpmAvIWD3KDkjcNPvgEfIweBrhKwDHDzIITj3LHnAk",
	"bCI0Y0+3/cIKvoCAZE7Yz4654VejrkHWhM5mW/xUlLAWqtJ1pwEYcerdErhUBpKihLmI0NilQ4dlMNTG",
	"ceCVk4FSJQ0XEjLLnBFoZYCY1SBMwYS79Z3+LT7jGr56NnTHN19H7v5cdXd9546P2m1slNCRjFyd9qs7",
	"sHHJqtV/hH4Yzq3FIqGfexspFlf2tpmLHG+if9j982ioNDKBFiL83aTFQnJTlfDivXxk/2IJuzRcZrzM",
	"7C8r+umHKjfiUizsTzn99FotRHopFgPIrGGNKlzYbUX/2PHi7NhsonrFa6WuqyJcUNpSXGdbdvFqaJNp",
	"zEMJ87zWdkPF42rjlZFDe5hNvZEDQA7iruC24TVsS7DQ8nSO/2zmSE98Xv5m/ymK3PY2xTyGWkvH7kpG",
	"84EzK5wXRS5SbpH41n22Xy0TAFIkeNPiFC/UFx8DEItSFVAaQYPyokhylfI80YYbHOnfS5hPXkz+7bSx",
	"v5xSd30aTP7a9rrETlZkJTEo4UVxwBhvrOijdzALy6DxE7IJYnsoNAlJm2hJSVgWnMOaS3PSqCwtflAf",
	"4HdupgbfJO0Qvjsq2CDCGTWcgSYJmBo+0CxAPUO0MkQrCqSLXM3qH744L4oGg/j9vCgIHyg9gkDBDDZC",
	"G/0Ql8+bkxTOc/HqhH0Xjo2iuJL51l4OJGrYu2Hubi13i9W2JbeGZsQHmuF2qvLEbo1HgxXzj0FxqFYs",
	"VW6lnr20Yhv/1bUNycz+PqrzH4PEQtwOExcqWg5zpOPgL4Fy80WHcvqE48w9J+y82/duZGNHiRPMnWhl",
	"537SuDvwWKPwpuQFAei+0F0qJCpp1IhgvSc3HcnoojAHZzigNYTqzmdt73mIQoKk0IHh61yl13/lenmE",
	"Mz/zY/WPH07DlsAzKNmS6+XJJCZlhMerGW3MEbMNUcFns2Cqk3qJx1renqVl3PBgaQ7euFhCqMd+yPSg",
	"jOguP+F/eM7sZ3u2LeunYU/YFTIwTcfZPTJkVtsnBYFmsg3QCqHYihR8ZrXug6B82Uwe36dRe/QN2RTc",
	"DrlF1Dt0tRGZPtY24WBDexUKqBevSKMzsNIRra1eFS9Lvo2vneYag4ArVbAc1pB3QSCWhaMRQtTm6Hzh",
	"a7WJwfS12vR4gtrAUXbCjoNytcfuHvheOchUuR/zOPYYpNsFWlleI3uQoQhkZ2ms1eczVd6NHXf4rGSN",
	"DZ5xO2pwG007SMKmVZG4sxmx41GDzkDNs+duLtodPoaxFhYuDf8dsKDtqMfAQnugY2NBrQqRwxFIfxm9",
	"BWdcw9Mn7PKv518+fvLrky+/siRZlGpR8hWbbQ1o9oVTVpk22xwe9leG6mKVm/joXz3zltv2uLFxtKrK",
	"FFa86A9FFmGSCakZs+36WGujGVddAziKI4K92gjtjB47LGivhLYi52p2lM0YQljWzJIxB0kGe4np0OU1",
	"02zDJZbbsjqGbg9lqcro1VWUyqhU5ckaSi1U5HnpjWvBXAsv7xfd3wladsM1s3OjLbySKGFFKMts5Hi+",
	"T0NfbWSDm52cn9YbWZ2bd8y+tJHvTauaFVAmZiNZBrNq0VIN56VaMc4y7Ih39HdgSG4RK7g0fFX8NJ8f",
	"R3dWOFBEhxUr0HYmRi2s1KAhVZJcQ/aoq27UMejpIsbbLM0wAA4jl1uZouH1GMd2WJNfCYmvQHor00Ct",
	"tzDmkC1aZHl/9X0IHTTVAx0Bx6LjNX5Gy88ryA3/VpVXjdj3Xamq4uhCXnfOscvhbjHOtpTZvt6oIOQi",
	"b7sjLSzsJ7E1fpYFvfTH160BoUeKfC0WSxPoWW9KpebHhzE2SwxQ/EBaam779HXVH1VmmYmp9BFEsGaw",
	"hsNZug35Gp+pyjDOpMoAN7/SceFswIEFX87xwd+E8p5ZkuI5A0tdKa/saquC4XN2775oOiY8pROaIGr0",
	"wGNe/QpLrWg6co7IS+DZls0AJFMz92Lm3vJwkRzf4o0Xb5xoGOEXLbiKUqWgNWSJs9TtBc23o6vD7MAT",
	"Ao4A17Mwrdicl/cG9nq9F85r2CboOaLZF9//oh9+BniNMjzfg1hsE0Nvbfdwz6J9qMdNv4vgupOHZMdL",
	"YP5eYUahNJuDgSEUHoSTwf3rQtTbxfujZQ0lPlD+rhTvJ7kfAdWg/s70fl9oq2LAH9Kpt1bCsxsmuVRe",
	"sIoNlnNtkn1s2TZq6eB2BQEnjHFiHHhA8HrNtaFHdSEztAXSdYLzkBBmpxgGeFANsSP/4jWQ/tipvQel",
	"rnStjuiqKFRpIIutQcJmx1w/wqaeS82DsWudxyhWadg38hCWgvEdsmglhCBu6rcn53XSXxy+0Nh7fhtF",
	"ZQuIBhG7ALn0rQLshj5hA4AI3SCaCEfoDuXUjmjTiTaqKCy3MEkl635DaLqk1ufm56Ztn7i4ae7tTIFG",
	"VzTX3kF+Q5glb8Al18zBwVb82soeaAah1/8+zPYwJlrIFJJdlI8qnm0VHoG9h7QqFiXPIMkg59v+oD/T",
	"Z0afdw2AO96ou8pAQm5d8U1vKNl70ewYWuF4OiY8MvzCUnsErSrQEIjrvWfkDHDsGHNydPSgHgrnim6R",
	"Hw+XTVsdGRFvw7UydscdPSDIjqOPAXgAD/XQd0cFdk4a3bM7xX+CdhPUcsThk2xBDy2hGf+gBQzYUJ3H",
	"fHBeOuy9w4GjbHOQje3hI0NHdsCg+4aXRqSiQF3ne9geXfXrThB9d2UZGC5yyFjwgdTAIuzPyCGpO+bd",
	"VMFRtrc++D3jW2Q5udAo8rSBv4Yt6txvyNM1MHUcQ5eNjGrvJy4ZAur956wIHjaBDU9NvrWCmlnClt1A",
	"CUxXs5UwhjzY26quUUUSDhB919gxo3vVjL4p7nxmvcShguX1t2I6IZ1gN3xXHcWghQ6nCxRK5SMsZD1k",
	"RCEY5QDDCmV3XThneu9O7SmpBaRj2vikXV//D3QLzbgC9p+qYimXqHJVBmqZRpUoKKAAaWewIlg9p3N1",
	"aTAEOayANEn88uhRd+GPHrk9F5rN4cZHoNiGXXQ8eoR2nDdKm9bhOoI91B63i8j1gQ8+9uJzWkiXp+x3",
	"tXAjj9nJN53B61cie6a0doRrl39vBtA5mZsxaw9pZJybCY476i2n9WTfXzfu+6VYVTk3x3i1gjXPE7WG",
	"shQZ7OXkbmKh5Ddrnv9Ud8PoGkgtjaaQpBgTMnIsuLJ9KIxkn27YuNeJ1QoywQ3kW1aUkAKFPViRT9cw",
	"njByiEyXXC5Q0i9VtXAeeTQOcupKk02lrGRviKg0ZDYyQet0jHM7L2wf+WLlIOBWF+uatknzuOH1fC7Y",
	"acyVGiCva+qPvm5NJ4OqqkXqulFVCTnt8J0RXLwlqAX4aSYe+QaCqLNCSx9f4bbYU2A39/extTdDx6Ds",
	"Txz4CDYfh9wErZ6cb48grdBArISiBI13S2hf0vRVzcNQPXf56K02sOqb4KnrrwPH7+2goqdkLiQkKyVh",
	"G41OFxJ+wI/R44T320BnlDSG+naVhxb8HbDa84yhxvviF3e7e0K7T036W1Ue6y2TBhwtl494Otz7Tu6m",
	"vOsDJ8/zyJugC+TpMgA9rRMHiJJxrVUqUNi6yPSUDpp7RnRRP230v6ndk49w9rrjdh6/whhRNO5CXjDO",
	"0lyg6VdJbcoqNe8lR+NSsNSI15LXoofNjS99k7h9M2J+dEO9lxw91mqTU9TTYg4R+8q3AN7qqKvFArTp",
	"KClzgPfStRKSVVIYnGtlj0tC56WAEl2HTqjlim/Z3NKEUew3KBWbVaYttmOcmjYiz91LnJ2Gqfl7yQ3L",
	"gWvDfhDyaoPD+dd6f2QlmBtVXtdYiN/uC5CghU7i3lXf0Vf0BHbLXzqvYMwrQJ+9l2UTODuxy2zFyv+f",
	"L/7jxbvz5L948ttZ8vx/nH74+Oz24aPej09u//KX/9v+6entXx7+x7/HdsrDHouicpBfvHIq7cUr1Fua",
	"x5se7J/McL8SMokSWeiG0aEt9gVGDDsCeti2apklvJdmIy0hrXkuMstb7kIO3RumdxbpdHSoprURHSuW",
	"X+uB2sA9uAyLMJkOa7yzFNV3SIzHK+JrogtBxPMyryRtpZe+KRzHO4ap+bSOSaV0NS8YBiwuufdqdH8+",
	"+fKrybQJNKy/T6YT9/VDhJJFtomFk2awiSl57oDgwXigWcG3GkyceyDsUR84csoIh13BagalXori03MK",
	"bcQszuF8kIMzFm3khSSPdnt+8G1y65481PzTw21KgAwKs4ylsWgJatiq2U2Ajr9IUao1yCkTJ3DSNdZk",
	"Vl903ng58DmmU0DtU43RhupzQITmqSLAeriQURaRGP10/Pnd5a+Prg65gWNwdeesHyL930axB999c8VO",
	"HcPUDyiymYYOYlEjqrQLt2p5ElluRsl7SMh7L9/LVzAXUtjvL97LjBt+OuNapPq00lB+zXMuUzhZKPbC",
	"R3C94oa/lz1JazC/VhA7x4pqlouUXYcKSUOelDOlP8L79+94vlDv33/oOVX01Qc3VZS/0ASJFYRVZRKX",
	"8SEp4YaXsUcrXUf848iU0mXXrCRkq4osmz6jhBs/zvN4Uehu5G9/+UWR2+UHZKhdXKvdMqaNKr0sYgUU",
	"ggb390flLoaS33i7SqVBs7+vePFOSPOBJe+rs7OnwFqhsH93V76lyW0Bo60rg5HJXaMKLpzUStiYkicF",
	"X8Text6/f2eAF7j7KC+v0MaR5wy7tUJwvUc9DtUswONjeAMIjoPDCXFxl9TLZ/eKLwE/4RZiGytuNC/2",
	"d92vICj3ztvVCezt7VJllok929FVaUvifmfqpD8LK2R5NwotFqituvxIM2DpEtJrl7gGVoXZTlvdvaeO",
	"EzQ96xCaUhpRSB0m1cCXhRmwqsi4E8W53HazG2gwxvsDv4Vr2F6pJifHIekM2tH1euigIqUG0qUl1vDY",
	"ujG6m+/cwVCxLwofpI7Rip4sXtR04fsMH2QSeY9wiGNE0Yr+HkIELyOIIOIfQMEdFmrHuxfpx5ZntYwZ",
	"3XyR9Eae9zPXpFGenOdWuBq0utP3FWB+NHWj2YxbuV251F4UQR5wsUrzBQxIyOHjzsg47daDEA6y796L",
	"3nRq3r3QevdNFGRqnNg1RykF7BdLKqjMdPz1/Ez0fuheJjBjp0PYLEcxqXZsJKbDy9YjG6UgHAItTsBQ",
	"ykbg8GC0MRJKNkuufdYxTM7mz/IoGeB3zIiwKw/OReBqFmRgq7PceJ7bPac97dJlw/EpcHzem1C1HJHD",
	"xkr46N0e2w4lUQDKIIcFLZwae0JpsjM0G2Th+Gk+z4UElsS81gIzaHDNuDnAysePGCMLPBs9QoyMA7Dx",
	"XRwHZj+q8GzKxSFASpddgvux8UU9+BvicV/kx21FHlVYFi4GXrVSzwG4c3Ws76+Owy0Ow4ScMsvm1jy3",
	"bM5pfM0gvXQsKLZ2kq84z4yHQ+LsjgcQulgOWhNdRXdZTSgzeaDjAt0OiGdqk1DgZ1TinW1mlt6jru0Y",
	"hho7mJT45oFmM7VBbx+8WsiVeg8sw3B4MAINfyM00iv2G7rNCZhd0+6WpmJUqJFknDmvJpchcWLM1AMS",
	"zBC5fBHksrkTAB1jR5MY2im/e5XUtnjSv8ybW23a5GjzUUOx4z90hKK7NIC/vhWmzj7zpiuxRO0UbaeV",
	"duKdQISMEb1lE/1Hmv5TkIYcUClIWkJUch17ObW6DeCNc+m7BcYLTO/D5fZh4AlVwkJoA40R3ftJfA7z",
	"JMesgkrNh1dninJu1/dWqfqaomdE7Nha5idfAboSz0WpTYIvENEl2EbfalSqv7VN47JS29eKcvCKLM4b",
	"cNpr2CaZyKs4vbp5v39lp/2xZom6miG/FZIcVmaYMzrqgbljanLS3bng17Tg1/xo6x13GmxTO3FpyaU9",
	"xx/kXHQ47y52ECHAGHH0d20QpTsYZBA52+eOgdwUvPGf7LK+9g5T5sfe67Xj43eH7igaKbqWwGCwcxUC",
	"n4msWCJMkHK5H9I6cAZ4UYhs07GF0qiDGjM/yODhE9V1sIC76wbbg4HA7hmLqilBt3MSNgI+Jc9uZcA5",
	"GYWZq3bmwJAhhFMJ7Us/9BFVR93tw9UV8Px72P5i2+JyJrfTyf1MpzFcuxH34PpNvb1RPOPTPJnSWi8h",
	"B6KcF0Wp1jxPnIF5iDRLtXakic29PfoTs7q4GfPqm/PXbxz4t9NJmgMvk1pUGFwVtiv+MKui9IcDB8Sn",
	"lrc6n5fZSZQMNr/O2RYapW+W4HJ0B9JoL5lo8+AQHEVnpJ7HPYT2mpzd2wgtcccbCRT1E0ljvqMXkvar",
	"CF9zkXu7mYd2wJsHFzcuI22UK4QD3Pt1JXgkS47KbnqnO346Guraw5PCuXZkEV9RonzNlOw+oaPP87Zw",
	"r+4rjqlAySrSZ06yWqElIdG5SOM2VjnTljgkvZ3ZxgwbDwijdsRKDDzFykoEY9lmY3LbdIAM5ogiU0fT",
	"6zS4mylXBKmS4p8VMJGBNPZTiaeyc1AxTYqztvevUys79OdyA5OFvhn+PjJGmAa3e+MhELsFjPClrgfu",
	"q1pl9gutLVL2h+BJ4oAH/3DG3pW447He0YejZnJeXLZf3MKaRX3+ZwmDktfvL5jklVeXj3dgjmgBJKGT",
	"eal+g7ieh+pxJGDJJ/4V6OXyG4SBDmHZjxaLqa07TR2nZvbB7R6SbkIrVNtJYYDqceeDZznMQOot1FzS",
	"VlMgScvXLU4woVfpKY3fEIyDueeJm/ObGY+lZ7VChoXpvHkAbtnSjWK+s8e9rqMtaHYWvCXXbQUFoxdQ",
	"NrGE/cQ2dxQYaNrRokIjGSDVhjLBlN7/cq0iw1Tyhksqa2P70VFyvTWQ8cv2ulElppLQcbN/BqlY8Twu",
	"OWRp38SbiYWgii2VhqAkiBuIqmERFbmyKnUMkUPNxZydTYO6RG43MrEWWsxywBaPqcWMa+TktSGq7mKX",
	"B9IsNTZ/MqL5spJZCZlZakKsVqwW6lC9qR+vZmBuACQ7w3aPn7Mv8NlOizU8tFh09/PkxePnaHSlP85i",
	"F4CruLOLm2TITv7m2EmcjvHdksawjNuNehKNuqeSe8OMa8dpoq5jzhK2dLxu/1lacckXEPcUWe2Bifri",
	"bqIhrYMXmVG9KG1KtWXCxOcHwy1/GvA+t+yPwGCpWq2EWbnHHa1Wlp6aeh80qR+Oik+5VM0eLv8R30gL",
	"/0TUUSI/rdGU7rfYqvEl+0e+gjZap4xT/pBcNN4LPoE8u/DpiTB3dZ2ymnBj57JLRzEHnRnmrCiFNKhY",
	"VGae/JmlS17y1LK/kyFwk9lXzyI5oNtpUuVhgH9yvJegoVzHUV8OkL2XIVxf9oVUMllZjpI9bKI9glM5",
	"+Jgbf7YbejvcPfRYocyOkgySW9UiNx5w6nsRntwx4D1JsV7PQfR48Mo+OWVWZZw8eGV36Oe3r52UsVJl",
	"LOdgc9ydxFGCKQWs0Xcvvkl2zHvuRZmP2oX7QP95Xx68yBmIZf4sRxWB9eoXb5Yd9Nm3IvwvP7j6kj3Z",
	"e8DPgBwJ6j6fOBYh6pJEEhq68TFcNfv747+zEuauYuSjRwj0o0dTJ8z9/Un7MzGpR4/imXiiNg37a4OF",
	"g1hhN1OB7Rvbw69VxMLg097XryEu3iBi4RlitfaDPcozN9SUtVOMf/q78DiebPHXyvgpeP/+HX7xeMA/",
	"uoj4zEceN7Dxx6CVDBBKUGIhSjJZ/T3wk+Dsa7UZSzgdTuqJ518ARVGUVCLPfmmidzusreQyXUbfPWe2",
	"469N8cF6cXR4oykgl1xKyKPDkc7wq9ctItrPP9TYeVZCjmzbLapBy+0srgG8DaYHyk9o0StMbicIsdoO",
	"jKwd7/OFyhjO0+QbbI5rvzpNkDL/nxVoE7uw8AM5/6F927IDytjOQGZoVThh31F98SWwVjIp1OZ9to92",
	"5HtV5IpnU8xCcvXN+WtGs1IfKqFFGeMXqMy2V9GxawapVMe5kftqWPEQl/Hj7Pa5t6vWJqkTvMeCiG2L",
	"JgW96Lz1oJobYueEvQoqBVO8sR2CYRKacmU183o0knGRJux/jOHpElX3FmsdJvnxpQ48Veqg3mpdN63O",
	"L4rnzsLtqh1QsYMpU2YJ5Y3QVFYa1tCOW66D+J3pyMcxt5dXVlISpZwccMvV2UQPRbsHjq5I/xwUhayD",
	"+AMVN6oUcmjlh0vsFU131i0j0Su0SlGwdT2sH3ypXC6VFCkmG4td0a7+9Ji30hF52brGeH/E3QmNHK5o",
	"8YrandJhcbCchWeEDnH9x5rgq91Uog7602Ch4yU3bAFGO84G2dTXYHH2YiE1uHyxWK084JOqbL0/I4eM",
	"ujQk9dPXgWSE4VMDBoBv7bcfnXkI4wquhURF0KHNCX5k0cXyuMZqj8KwhQLt1tOOIdfvbJ8TDKfOYPPh",
	"xJfTxTHo+dYum3wV+kOde88F5ylg2760bV2Sq/rnlqc6TXpeFG7S4Qo98bJkGzmI4MgLdOKfAAPk1uOH",
	"o+0gt50uR3ifWkKDNTosQIH3cI8w6mo1ndJwVmglisIWjFz9opkuhIyA8VpIaIo9Ry6INHol4MbgeR3o",
	"p9OSGxIBR/G0K+A5KdQRhqaNe6K671DdFF8WJbhGP8fwNjaFdgYYR92gEdy43NY1pi11B8LESyxu7xDZ",
	"L5uDUpUTojKMPOkU0okxDsu4famu9gWwpzrftOmO+e4OvYmGgolnVbYAk/Asi6Xv/Rq/MvzKsgolB9hA",
	"WtVpXouCpZg7p51MqE9tbqJUSV2tdszlG9xzuqAyVYQawupYfocxWGm2xX8PqZtYO+sc7C7qPXOywzJo",
	"9d1fY1KvpelEi0UyHhN4p9wfHc3UdyP0pv9RKT1XizYgn8NsN8Dlwj2K8bdv7MURZtjoJe6lq6VOgIHO",
	"mcoXWEW1sQ7dbnMlvMp6mXzxUbCuV7jbADFceXCKl9+Ai3ZohKX7lQyTQ47a6WBcATcuwtFwtpMFDUaN",
	"kZdXx6zbt7APeXaRY9fxzKFurTsR6l0G+wB97/2RWcGFc6FomEUfsy5yoR9LMsanudng7iJcPMCgxe77",
	"9ZDvvk+oh9+7lcmuwaU9KEpYC1V55wTvveZVQvq1Veerjp6Irr9veMWpPq85dNB4e+UqRNAynU7+/S/k",
	"68hAmnL7L2DK7W16r+ZZX9ol81TThNXJxUclG2/dimOSTcbyGjrZsFV1bU/NuB5ZvRojDvRrwE0nF9lB",
	"F2YsN+aERokdu3hFt+HUYU26MDxihdKiyfEfK/U20k30Cqu1BanP+mN5H601pAYLOzS+JyXAIYnQ7GRB",
	"8dj/TiE2oE7X3rQuc9iudGH9ag577vheRF8QlUqZ8E/GJ8c6rz0MkU9jRusFSFe/tR2rMzpiYD6H1Ij1",
	"ngjKvy1BBtF5U2+XocL0QUClqD3QMQHP4VbHBqBdAY474QkSYd4bnKH4qWvYPtCsRQ3R1PxTf9XeJfcK",
	"YgC5Q2JJROmYBw8Zkp1ThdA1ZSAWvMccdYcmi91gVa8gHviOc3mStBdHEyO8Y8p4WaFRc9muB0XOozP1",
	"UJBlvyrJsP7xCovA6Lrips/dEmrp7KKf4fLG5X7BeNf67cRngQHtf/PB7TRLLq4hrDuGL1U3vMx8i6jp",
	"xVt1kh33US8y0lfU6AI9r2cWjX9zPxYukjMNvdjTXFkxIhkKBWi7FNf+OA80OU5RCn90lrZwzaF09RlR",
	"/s2VhsQo7w+9C45dqCDvsDshQQ/mKSXgBrMHvW3SI2G+Zo7ZgrhzCgsXyEpYcQtdGSQxGp5zF7Jf0ncf",
	"/OXz9e61MNX0ur9whPdsF7qHxJDq58zdlvuDyu5ibBJSUg1wHctoJKFsv4YUpcqqlC7o8GDUBrnR+cJ2",
	"sJKonSbtr7KjIwSRudewPSUlyFfc8DsYAk2SE4EeZMLobPJRzW86BvfiKOB9TsvVdFIolScDjx0X/TRM",
	"XYq/Fuk1ZMzeFN4DdKAKEvsCbez1a/bNcuvTDhUFSMgenjB2Lsnn3j9st/OAdyaXD8yu+Tc4a1ZRZjRn",
	"VDt5L+POy5izrLwnN/PD7OZhGiyru+dUNMieJD+bgRRQJb+J1AQ7GauV95+au3WaGqIiKGIyySW9WL3E",
	"gx4zHN2UwoBzbKBL3G4kcy9dTOcq5iQIN+Pi92uHUrsjuRq4uMPJECADckycZw2FGzyKgLoG0x5HodpH",
	"qClf0/gJ9cWjPFc3CR6jpE5iF1O6bLv2LeHT9jbdLLnNIHA44tpJEFu25BlLVVlCGvaIx+kQUCtVQpIr",
	"9D+KPY3OjRUIV+icL1muFkwVVs+nXJD+ESlaWymY61h1pCjmnCBI6MVrIKsHaBdj7sClxn14d5RyOrxM",
	"1NUyYrjCDfO7dXAtKEdwB5dwCcAcQej7jXbnsVJX7XV1i64NlUA0aiXSOLr/WO46g042MeqNocJlUaYo",
	"TmyGBzzkKfXrLJ6ePppB8lke5dXu+LlXKqRz+1+8wrvjsjk45jLAzyI1m4kNJ+ngZdEBACGl0CJTlZR6",
	"OWTldUE3taBQRHxj6wI6kuGgK8P9YLMjHBOo292EEqv4FjkI9e64gnQ+lnrgUEWdJHb7JFAV0NlYz4Q6",
	"0/xI/hkAMOyr0IJhlMfCoWDMsapuwiNIvqj1xGmr6LnoXBI+Cygxw5STnWgJzI5dleBie6n8Z6feWMHN",
	"0suNtnnfmiMz2IDGwFsqmsQ12R69DdTVHu0K5KpIclhDy4XDBRxXaQpaizWEdUupM8sACnwR6OqpMd+E",
	"8DrsKC9u7Unwuj0Gu1FthhBLO8X2qCpRxWojEzomeuxRshCtRVbxFv70PSo4DhVvjNzXHtYP4zjFwUwi",
	"vrhdLGKvNxHSfPRcyrgzURjvXpshcbasfq4gImxOti74jRxW2/tE2Yib42ufBoj9ZgMpXt1tb5n744Th",
	"YEx3clkMypllvcN3Nf8MUtkuIutVgo3rYeAreYdpp7yu4PpGrkYyVAsdGUDohjeg7y00vp1BsxXfskzM",
	"51DSU5w2XGa8zMLmQrIUSsOFZDd8q++uk1loywqme9Uyy6lxUM+sYgoaWpUJkHzrFP4hlWmEqoPvrhE1",
	"h65to4aK1PZ2JR4MxDdWNUSvyAEicKkoUDGkw6okSuVsxa/hwHm0+A12T4MJopzl3iicdcwUtztp/SdE",
	"HR74n6UwO6md5L2umyq9IxIxehqUi8aZgTanT4Mxz+IrKpUWehd3K4/4vSajJs0HA5lU22L6wC6iWce5",
	"pYcyuR6vrrYsRzH/ZeLhCfJ2vcNdAXRQqy115ua+WNK7FAgpU+f9faDUQuoCzzIxVBp/CS5duTtb7Wlr",
	"E6AdZ7ylO7B3xSEqVJGkY96wMsjBshrSWhykbRhH2MiKdM+1EL0kB7hSW0VSc+QPeCxINEBvn/pCnHb9",
	"0NpCQH3wsO5yWpUoxt7w7f6UmI0gEHfhp5G9Du49k2qo3QbTEddUyieacfIQATHCdWLVbPq5/o6/GIpN",
	"aV7Pf7/luPex+ALOpVOUsEbhLnprVClPKhFa43IbYxr+BegOCxySD0d4Vx9tq+rT8ntsUPSSvFsK6FGg",
	"9T1tI9gMarbvdn4KM8Q3aQtKcthGZwmvkXb5xQ+NpjquerzvsAe80CcuqB/vnycdOJ85/v+HGinBUj4M",
	"UUJr+fvc7NwCG9U+2CInLRsDVK+DYkbb+xL4UOqXtWviwNXc82DEdPBWPMvziOcjCfBUXDwgHHsvlmue",
	"f3rvRawTcI74gOztsL9D6P4WIplQqe8WfPuaj5o7cHU73tTyDXpb/g3sHkWvBTeUsxn0mD+qXzynp6m5",
	"rzS8BslucEyy2D7+is1cgqmihFTori3ixhcBrL29sCauC3jemD3uZfvW+Ysy9yDjuTftsR+bgmL4+rKQ",
	"DYTNEf3MTGXg5EapPEZ9PbKI4C/Go8JMz3uui+tWDEcj1QU3mirhyLEcQVTmgbEc/RzWY5dH8Qr20qk0",
	"9Nc5+rZu4TZyUTdrGxuINDobFFZ7GhM/FM/cZLtjANNRUjgdlMDpdwhdIhy5Mdy8MYr5ZSiZBSVsGMib",
	"0tmPSuTZPsJoZcG5rWvkY56XX12+tE97l3oIyJ26f1Rdyep7xIAQYiJrbU0eTBXktxmR2sZ1iySyQVel",
	"tCqF2WIad6/xil+jQVbf1Q77LuCjNqK6u8+oa6gLATTu/ZX2t+t3iud4H5FtV9pbSOUn7JsNXxW5s4mw",
	"vzyY/Qme/vlZdvb08Z9mfz778iyFZ18+Pzvjz5/xx8+fPoYnf/7y2Rk8nn/1fPYke/LsyezZk2dfffk8",
	"ffrs8ezZV8//9MDyIQsyATrxSUMn/zs5zxcqOX9zkVxZYBuc8EJ8D1sqX27J2BdG5ymeRFhxkU9e+J/+",
	"pz9hJ6laNcP7XycuJ+FkaUyhX5ye3tzcnIRdThfoz5sYVaXLUz9Pr3L6+ZuL+t2cnl1wR2uPKfLFcaRw",
	"jt/efnN5xc7fXJw0BDN5MTk7OTt5bMdXBUheiMmLyVP8CU/PEvf91BHb5MXH2+nkdAk8x/AX+8cKTClS",
	"/6kEnm3d//UNXyygPHHV4u1P6yenXqw4/ej8mm93fTsNCy+efmy5f2d7emJhttOPPt/47tathN7O7T3o",
	"MBKKXc1OZ5gCb2xT0EHj4aWgsqFPP6K4PPj7qcvZFf+Iagudh1MfIxFv2cLSR7OxsO7psRFZsJKUm3RZ",
	"Facf8T9IvbfETnKIxUtQIizOmuZTJgzjM1ViGnCTLi0H8fmHhQ5aTpCm6ThcZPYY2F4vCQJfaYBKL714",
	"1/fcwIGYHwl5hj0QzZFuzdRwbXzECaoB1XdSq31zM707S55/+Ph4+vjs9t/szeP+/PLp7Uh/opf1uOyy",
	"vlZGNvyAyXvxDQ9P+pOzM8/enPIQkOapO8nB4npKVLNI2qQ6kr1/6ztaGH7Bd1vVGYjVyNiTZLQzfF94",
	"QY7+7MAV77Q0taL7cfhu3sGMeZdPnPvxp5v7QmLYmb0BGN1wt9PJl59y9RfSkjzPGbYMssb3t/5neS3V",
	"jfQtrThSrVa83PpjrFtMgbnNxkuPLzS+MZRizVEKlEq2SmFPPqDze8ztdoDfaMPvwG8uba//5jefit/g",
	"Jh2D37QHOjK/eXLgmf/jr/j/bw777OzPnw4CHzVwJVagKvNH5fCXxG7vxeGdwEkpmU7NRp6iR8rpx5b4",
	"7D73xOf27033sMV6pTLw8q6az6lg167Ppx/p32Ai2BRQihVIKmTgfqV0FaeYRn/b/3kr0+iP/XUUndrT",
	"sZ9PP7artbYQpJeVydQN5R2OXplYkoznrn4JGpNrxdQo5gdocgOwn1w6o3yLFnSRAeOYZ1VVprEc2M61",
	"12j9tmNHYHrpjOgLIXECNNLjLFSohwceCxpSJTPUhzvXs4PsR5VB/3rGC/ifFZTb5gZ2ME6mLf7sCDxS",
	"Fufe112fnd4eRv74mEAvYX3icEXRO3+f3nBh7CXugvQRo/3OBnh+6jJydn5tkmD1vmBmr+DH0PU1+utp",
	"XVku+rGrqMe+OkV1oJH3nvOfG6NdaARDkqjNX+8+2J3FuiWOWhqbzovTUwx8XSptTie3048de0/48UO9",
	"mT5Reb2ptx9u/18AAAD//9/FM7B32gAA",
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
