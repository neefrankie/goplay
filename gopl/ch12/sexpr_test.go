package ch12

import (
	"testing"
)

func TestMarshal(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}

	var strangelove = Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Actor: map[string]string{
			"Dr. Strangelove":           "Peter Sellers",
			"Grp. Capt. Lion Mandrake":  "Peter Sellers",
			"Pres. Merkin Muffly":       "Peter Sellers",
			"Gen. Buck Turgidson":       "George C. Scott",
			"Brig. Gen. Jack D. Ripper": "Sterling Hayden",
			`Maj. T.J. "King Kong"`:     "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
		Sequel: nil,
	}

	b, err := Marshal(strangelove)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%s", b)
}

func TestBits(t *testing.T) {
	t.Logf("%b", 1<<2)
	t.Logf("%b", 4>>2)
}
