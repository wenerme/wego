package serves

import (
	"github.com/rs/zerolog/log"
	"reflect"
)

type Endpoint interface {
	String() string
}

var Endpoints []Endpoint

func RegisterEndpoints(eps ...Endpoint) {
	for _, v := range eps {
		log.Trace().
			Str("endpoint", v.String()).
			Str("type", reflect.TypeOf(v).String()).
			Msg("registering endpoint")
	}
	Endpoints = append(Endpoints, eps...)
}

func FilterEndpoints[T Endpoint](filter func(e T) bool) []T {
	var eps []T
	for _, v := range Endpoints {
		vv, ok := v.(T)
		if !ok {
			continue
		}
		if filter == nil || filter(vv) {
			eps = append(eps, vv)
		}
	}
	return eps
}
