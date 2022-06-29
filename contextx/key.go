package contextx

import (
	"context"
	"crypto/rand"
	"fmt"
	"reflect"
)

// Key is a context key
type Key[T any] struct {
	name string
	x    [8]byte
}

// FromContext returns the value stored in the context for the given key.
func (key Key[T]) FromContext(ctx context.Context) (o T, ok bool) {
	if ctx == nil {
		ok = false
		return
	}
	o, ok = ctx.Value(key).(T)
	return o, ok
}

// NewContext returns a new context with the given key and value set.
func (key Key[T]) NewContext(ctx context.Context, val T) context.Context {
	return context.WithValue(ctx, key, val)
}

// WithDefaultProvider is a helper function to create a new key with a default value provided by a provider
func (key Key[T]) WithDefaultProvider(ctx context.Context, provider func(context.Context) (context.Context, T)) context.Context {
	if key.Exists(ctx) {
		return ctx
	}
	next, val := provider(ctx)
	return key.NewContext(next, val)
}

// WithDefaultValue is a helper function to create a new key with a default value
func (key Key[T]) WithDefaultValue(ctx context.Context, val T) context.Context {
	if key.Exists(ctx) {
		return ctx
	}
	return key.NewContext(ctx, val)
}

func (key Key[T]) String() string {
	name := key.name
	if name != "" {
		name = "@" + name
	}
	return fmt.Sprintf("contextx.Key(%s%s)", reflect.TypeOf(new(T)).Elem().String(), name)
}

// Get returns the value stored in the context for the given key.
func (key Key[T]) Get(ctx context.Context) T {
	o, _ := key.FromContext(ctx)
	return o
}

// GetWithDefault is a helper function to get a value from the context with a default value
func (key Key[T]) GetWithDefault(ctx context.Context, def T) T {
	v, found := key.FromContext(ctx)
	if !found {
		return def
	}
	return v
}

// Exists return true if the context contains a value for the given key.
func (key Key[T]) Exists(ctx context.Context) bool {
	_, ok := key.FromContext(ctx)
	return ok
}

// Must return the value stored in the context for the given key.
func (key Key[T]) Must(ctx context.Context) T {
	o, ok := key.FromContext(ctx)
	if !ok {
		panic(fmt.Errorf("%s: not found in context", key.String()))
	}
	return o
}

// KeyOptions is a struct to configure a key
type KeyOptions struct {
	Name    string // the name of the key
	Private bool   // if true, the key is private and should not be shared
}

// NewKey returns a new Key
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
