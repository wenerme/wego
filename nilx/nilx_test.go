package nilx

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestZero(t *testing.T) {
	assert.Nil(t, IsZeroToNil(time.Time{}))
}
