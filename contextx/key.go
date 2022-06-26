package contextx

import (
	"context"
	"crypto/rand"
	"fmt"
	"reflect"
)

type Key[T any] struct {
	name string
	x    [8]byte
}

func (key Key[T]) FromContext(ctx context.Context) (T, bool) {
	o, ok := ctx.Value(key).(T)
	return o, ok
}

func (key Key[T]) NewContext(ctx context.Context, val T) context.Context {
	return context.WithValue(ctx, key, val)
}

func (key Key[T]) String() string {
	name := key.name
	if name != "" {
		name = "@" + name
	}
	return fmt.Sprintf("contextx.Key(%s%s)", reflect.TypeOf(new(T)).Elem().String(), name)
}

func (key Key[T]) Get(ctx context.Context) T {
	o, _ := ctx.Value(key).(T)
	return o
}

func (key Key[T]) Exists(ctx context.Context) bool {
	_, ok := ctx.Value(key).(T)
	return ok
}

func (key Key[T]) Must(ctx context.Context) T {
	o, ok := ctx.Value(key).(T)
	if !ok {
		panic(fmt.Errorf("%s not found in context", key.String()))
	}
	return o
}

type KeyOptions struct {
	Private bool
	Name    string
}

func NewKey[T any](o *KeyOptions) *Key[T] {
	if o == nil {
		o = &KeyOptions{}
	}
	k := &Key[T]{
		name: o.Name,
	}
	if o.Private {
		_, _ = rand.Read(k.x[:])
	}
	return k
}
