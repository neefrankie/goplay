package web

import (
	"testing"

	"github.com/google/uuid"
)

func TestUUID(t *testing.T) {
	t.Logf("%s", uuid.NewString())
}
