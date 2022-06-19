package serves

import (
	"fmt"
	"strings"
)

type EndpointDesc struct {
	Name                   string
	Description            string
	Selector               string
	Tags                   []string
	Metas                  map[string]string
	Disabled               bool
	DeprecationDescription string
}

func (e EndpointDesc) String() string {
	var s []string
	if e.Name != "" {
		s = append(s, e.Name)
	}
	if e.Selector != "" {
		s = append(s, "@"+e.Selector)
	}
	if e.Disabled {
		s = append(s, "disabled")
	}
	if len(e.Tags) != 0 {
		s = append(s, "tags="+strings.Join(e.Tags, ","))
	}
	if len(e.Metas) != 0 {
		s = append(s, "meta="+fmt.Sprint(e.Metas))
	}
	return fmt.Sprintf("Endpoint(%v)", strings.Join(s, ","))
}
