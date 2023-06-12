package recur

import (
	"testing"
)

func TestMax(t *testing.T) {
	bytes := []byte{'T', 'I', 'N', 'Y', 'E', 'X', 'A', 'M', 'P', 'L', 'E'}
	m := Max[byte](bytes, 0, 10)

	t.Logf("%s", string(m))
}
