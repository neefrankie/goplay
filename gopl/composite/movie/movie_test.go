package movie

import (
	"encoding/json"
	"testing"
)

var movies = []Movie{
	{
		Title:  "Casablanca",
		Year:   1942,
		Color:  false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"},
	},
	{
		Title:  "Cool Hand Luke",
		Year:   1967,
		Color:  true,
		Actors: []string{"Paul Newman"},
	},
	{
		Title:  "Bullitt",
		Year:   1968,
		Color:  true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"},
	},
}

func TestMovie(t *testing.T) {
	data, err := json.Marshal(movies)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s\n", data)
}

func TestMovie_indent(t *testing.T) {
	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s\n", data)
}
