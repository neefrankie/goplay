package store_test

import (
	"testing"

	"github.com/youkre/firefly/config"
	"github.com/youkre/firefly/store"
)

func TestStore_LoadIndex(t *testing.T) {
	store := store.NewStore(config.NewDataPath(), config.CME)

	idxPage, err := store.LoadIndex()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(idxPage)
}
