package html

import (
	"testing"
)

func TestNewHead(t *testing.T) {
	h := NewHead().
		WithBasePath("/my-site").
		WithMeta(NewMeta().CharSet()).
		WithMeta(NewMeta().ViewPort("")).
		WithMeta(NewMeta().UACompatible("")).
		WithTitle("My Awesome Page")

	t.Logf("%s", h.String())
}
