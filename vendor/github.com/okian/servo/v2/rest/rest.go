package rest

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type route struct {
	method      string
	path        string
	handler     echo.HandlerFunc
	middlewares []echo.MiddlewareFunc
}

var (
	routes          []route
	middlewares     []echo.MiddlewareFunc
	customValidator echo.Validator
)

func Validator(v echo.Validator) {
	if customValidator != nil {
		panic("multiple Validator call")
	}
	customValidator = v
}

func Use(m echo.MiddlewareFunc) {
	middlewares = append(middlewares, m)
}

func registerRoute(me, p string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) {
	routes = append(routes, route{
		method:      me,
		path:        p,
		handler:     h,
		middlewares: m,
	})
}

func EchoGet(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) {
	registerRoute(http.MethodGet, path, h, m...)
}

func Get(path string, h http.HandlerFunc) {
	EchoGet(path, echo.WrapHandler(h))
}

func EchoPost(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) {
	registerRoute(http.MethodPost, path, h, m...)
}

func Post(path string, h http.HandlerFunc) {
	EchoPost(path, echo.WrapHandler(h))
}

func EchoPut(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) {
	registerRoute(http.MethodPut, path, h, m...)
}

func Put(path string, h http.HandlerFunc) {
	EchoPut(path, echo.WrapHandler(h))
}

func EchoPatch(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) {
	registerRoute(http.MethodPatch, path, h, m...)
}

func Patch(path string, h http.HandlerFunc) {
	EchoPatch(path, echo.WrapHandler(h))
}

func EchoDelete(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) {
	registerRoute(http.MethodDelete, path, h, m...)
}

func Delete(path string, h http.HandlerFunc) {
	EchoDelete(path, echo.WrapHandler(h))
}

func EchoConnect(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) {
	registerRoute(http.MethodConnect, path, h, m...)
}

func Connect(path string, h http.HandlerFunc) {
	EchoConnect(path, echo.WrapHandler(h))
}

func EchoOptions(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) {
	registerRoute(http.MethodOptions, path, h, m...)
}

func Options(path string, h http.HandlerFunc) {
	EchoOptions(path, echo.WrapHandler(h))
}

func EchoTrace(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) {
	registerRoute(http.MethodTrace, path, h, m...)
}

func Trace(path string, h http.HandlerFunc) {
	EchoTrace(path, echo.WrapHandler(h))
}

func EchoHead(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) {
	registerRoute(http.MethodHead, path, h, m...)
}

func Head(path string, h http.HandlerFunc) {
	EchoHead(path, echo.WrapHandler(h))
}

type v10Validator struct {
	validator *validator.Validate
}

func (cv *v10Validator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
