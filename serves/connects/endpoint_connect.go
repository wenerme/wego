package serve

import (
	"fmt"
	connect_go "github.com/bufbuild/connect-go"
	"github.com/go-chi/chi/v5"
	"net/http"
	"reflect"
)

type ConnectServiceDesc[T any] struct {
	Name       string
	Service    T
	NewHandler func(svc T, opts ...connect_go.HandlerOption) (string, http.Handler)
	NewClient  func(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) T
}
type ConnectEndpoint struct {
	Name           string
	Service        any
	Path           string
	Handler        http.Handler
	Middlewares    chi.Middlewares
	HandlerOptions []connect_go.HandlerOption
	ClientOptions  []connect_go.ClientOption
	NewHandler     any //func(svc T, opts ...connect_go.HandlerOption) (string, http.Handler)
	NewClient      any //func(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) T
}

func (ep *ConnectEndpoint) MountConnect(r chi.Router) {
	if ep.Handler != nil || ep.Path == "" {
		args := []reflect.Value{
			reflect.ValueOf(ep.Service),
		}
		for _, v := range ep.HandlerOptions {
			args = append(args, reflect.ValueOf(v))
		}
		ret := reflect.ValueOf(ep.NewHandler).Call(args)
		ep.Path, ep.Handler = ret[0].String(), ret[1].Interface().(http.Handler)
	}
	r.With(ep.Middlewares...).Mount(ep.Path, ep.Handler)
}

func (ep *ConnectEndpoint) String() string {
	return fmt.Sprintf("ConnectEndpoint(%s)", ep.Name)
}
