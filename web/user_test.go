package web

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidateStruct(t *testing.T) {
	address := &Address{
		Street: "Evavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
	}

	user := &User{
		FirstName:     "Badger",
		LastName:      "Smith",
		Age:           135,
		Gender:        "male",
		Email:         "Badger.Smith@gmail.com",
		FavoriteColor: "#000-",
		Addresses:     []*Address{address},
	}

	err := validate.Struct(user)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			t.Fatal(err)
		}

		for _, err := range err.(validator.ValidationErrors) {
			t.Logf("%s\n", err.Namespace())
			t.Logf("%s\n", err.Field())
			t.Logf("%s\n", err.StructNamespace())
			t.Logf("%s\n", err.StructField())
			t.Logf("%s\n", err.Tag())
			t.Logf("%s\n", err.ActualTag())
			t.Logf("%s\n", err.Kind())
			t.Logf("%s\n", err.Type())
			t.Logf("%s\n", err.Value())
			t.Logf("%s\n", err.Param())
		}
	}
}
