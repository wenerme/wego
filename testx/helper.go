package testx

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func MustNonEOF[T any](v T, err error) T {
	if err != nil && !errors.Is(err, io.EOF) {
		panic(err)
	}
	return v
}
