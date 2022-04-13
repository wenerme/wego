package stdx

import (
	"errors"
	"io"
)

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
