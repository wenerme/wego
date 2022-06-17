package nilx

import "golang.org/x/exp/constraints"

type isZero interface {
	IsZero() bool
}

func IsZeroToNil[T isZero](v T) *T {
	if v.IsZero() {
		return nil
	}
	return &v
}

func ZeroToNil[T constraints.Integer](v T) *T {
	if v == 0 {
		return nil
	}
	return &v
}
