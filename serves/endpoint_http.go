package serves

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type HTTPEndpoint struct {
	EndpointDesc

	Method      string
	Path        string
	Handler     http.Handler
	HandlerFunc func(http.ResponseWriter, *http.Request)
	Middleware  func(http.Handler) http.Handler

	Children []*HTTPEndpoint

	Methods     []string
	Middlewares []func(http.Handler) http.Handler
}

func (e *HTTPEndpoint) Normalize() {
	if e.Method != "" {
		e.Methods = append(e.Methods, e.Method)
		e.Method = ""
	}
	if e.HandlerFunc != nil {
		e.Handler = http.HandlerFunc(e.HandlerFunc)
		e.HandlerFunc = nil
	}
	if e.Middleware != nil {
		e.Middlewares = append(e.Middlewares, e.Middleware)
		e.Middleware = nil
	}
}

func (e *HTTPEndpoint) GetEndpointDesc() *EndpointDesc {
	return &e.EndpointDesc
}

func (e *HTTPEndpoint) Validate() error {
	e.Normalize()
	if len(e.Path) == 0 && len(e.Children) == 0 {
		return errors.Errorf("endpoint missing path")
	}
	return nil
}

func (e HTTPEndpoint) String() string {
	method := e.Method
	if method == "" {
		method = "ANY"
	}
	return fmt.Sprintf("%v %v | %v", method, e.Path, e.EndpointDesc.String())
}

func HTTPEndpointSortByPathLen(a, b *HTTPEndpoint) bool {
	return len(a.Path) > len(b.Path)
}
