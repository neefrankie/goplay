package web

import (
	"context"
	"goplay/web/config"
	"goplay/web/ent"
	"testing"
)

func TestEnt(t *testing.T) {

	client, err := ent.Open("mysql", config.MustGetMyDSN("entdemo"))

	if err != nil {
		t.Fatal(err)
	}

	defer client.Close()

	// Run the auto migration tool
	err = client.Schema.Create(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}
