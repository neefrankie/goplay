package web

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
)

func TestFakeSimple(t *testing.T) {
	t.Logf("Name: %s\n", gofakeit.Name())
	t.Logf("Email: %s\n", gofakeit.Email())
	t.Logf("Phone: %s\n", gofakeit.Phone())
	t.Logf("BS: %s\n", gofakeit.BS())
	t.Logf("BeerName: %s\n", gofakeit.BeerName())
	t.Logf("Color: %s\n", gofakeit.Color())
	t.Logf("Company: %s\n", gofakeit.Company())
	t.Logf("CreditCardNumber: %s\n", gofakeit.CreditCard().Number)
	t.Logf("HackerPhrase: %s\n", gofakeit.HackerPhrase())
	t.Logf("JobTitle: %s\n", gofakeit.JobTitle())
	t.Logf("CurrencyShort: %s\n", gofakeit.CurrencyShort())
}
