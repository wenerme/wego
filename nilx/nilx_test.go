package nilx

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestZero(t *testing.T) {
	assert.Nil(t, IsZeroToNil(time.Time{}))
}
