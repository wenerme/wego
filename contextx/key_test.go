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
	assert.Equal(t, "contextx.Key(int@X)", Key[int]{name: "X"}.String())
	b := k
	assert.Equal(t, b, k)
	assert.Equal(t, b.Must(ctx), 1)
	assert.Equal(t, Key[int]{}.Must(ctx), 1)
	assert.Equal(t, (&Key[int]{}).Must(ctx), 1)
	assert.Equal(t, Key[int]{name: "X"}.Get(ctx), 0)

	k2 := NewKey[int](&KeyOptions{Private: true})
	assert.Equal(t, k2.Get(ctx), 0)
	assert.NotEqual(t, k2, k)
}
