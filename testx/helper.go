package testx

import (
	"encoding/json"
	"fmt"
	"github.com/wenerme/wego/stdx"
)

func PrintJson(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	NoErr(err)
	fmt.Println(string(b))
}

func NoErr(err error) {
	if err != nil {
		panic(err)
	}
}

var Must = stdx.Must
var MustNonEOF = stdx.MustNonEOF
