package contextx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKey(t *testing.T) {
	k := NewKey[int](nil)
	ctx := context.Background()
	ctx = k.NewContext(ctx, 1)
	assert.Equal(t, k.Get(ctx), 1)
	assert.Equal(t, k.Must(ctx), 1)
	assert.Equal(t, k.Get(context.Background()), 0)
	assert.False(t, k.Exists(context.Background()))
	assert.True(t, k.Exists(ctx))
	assert.Equal(t, "contextx.Key(int)", k.String())
	assert.Equal(t, "contextx.Key(int@X)", NewKey[int](&KeyOptions{Name: "X"}).String())
	b := k
	c := NewKey[int](nil)

	assert.Equal(t, b, k)
	assert.Equal(t, k, c)
	assert.Equal(t, b.Must(ctx), 1)
	assert.Equal(t, c.Must(ctx), 1)
	assert.Equal(t, (&Key[int]{name: "X"}).Get(ctx), 0)
	{
		k2 := NewKey[int](&KeyOptions{Private: true})
		k3 := NewKey[int](&KeyOptions{Private: true})

		assert.NotEqual(t, k2, k3)
		assert.Equal(t, k2.Get(ctx), 0)
		assert.Equal(t, k2.Get(nil), 0)
		assert.Equal(t, k2.GetWithDefault(nil, 10), 10)
		assert.NotEqual(t, k2, k)

		ctx = k3.WithDefaultProvider(ctx, func(ctx context.Context) (context.Context, int) {
			ctx = k2.WithDefaultValue(ctx, 2)
			return ctx, 3
		})

		assert.Equal(t, k2.Get(ctx), 2)
		assert.Equal(t, k3.Get(ctx), 3)
	}
}
