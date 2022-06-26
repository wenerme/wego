package testx

import (
	"encoding/json"
	"fmt"

	"github.com/wenerme/wego/stdx"
)

func PrintJSON(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	NoErr(err)
	fmt.Println(string(b))
}

func NoErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Must[T any](v T, err error) T {
	return stdx.Must(v, err)
}

func MustNonEOF[T any](v T, err error) T {
	return stdx.MustNonEOF(v, err)
}
