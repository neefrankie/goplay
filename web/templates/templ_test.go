package templates

import (
	"context"
	"os"
	"testing"
)

func TestTempl(t *testing.T) {
	component := hello("John")
	component.Render(context.Background(), os.Stdout)
}
