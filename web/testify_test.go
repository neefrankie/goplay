package web

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssertTrue(t *testing.T) {
	assert.True(t, true, "True is true")
}
