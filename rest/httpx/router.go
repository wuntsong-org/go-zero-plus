package httpx

import "net/http"

type MiddlewareFunc func(http.Handler) http.Handler

// Router interface represents a http router that handles http requests.
type Router interface {
	http.Handler
	Handle(method, path string, handler http.Handler) error
	SetNotFoundHandler(handler http.Handler)
	SetNotAllowedHandler(handler http.Handler)
	SetOptionsHandler(handler http.Handler)
	SetMiddleware(middleware MiddlewareFunc)
}
