package utility

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomBytes(t *testing.T) {
	bytes, err := RandomBytes(128)
	assert.Nil(t, err)
	assert.Len(t, bytes, 128)
}
