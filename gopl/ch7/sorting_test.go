package ch7

import (
	"sort"
	"testing"
)

var tracks = []*Track{
	{
		Title:  "Go",
		Artist: "Delilah",
		Album:  "From the Roots Up",
		Year:   2012,
		Length: length("3m38s"),
	},
	{
		Title:  "Go",
		Artist: "Moby",
		Album:  "Moby",
		Year:   1992,
		Length: length("3m37s"),
	},
	{
		Title:  "Go Ahead",
		Artist: "Alicia Keys",
		Album:  "As I Am",
		Year:   2007,
		Length: length("4m36s"),
	},
	{
		Title:  "Ready 2 Go",
		Artist: "Martin Solveig",
		Album:  "Smash",
		Year:   2011,
		Length: length("4m24s"),
	},
}

func TestSortTrack(t *testing.T) {
	sort.Sort(byArtist(tracks))
	printTracks(tracks)

	sort.Sort(sort.Reverse(byArtist(tracks)))
	printTracks(tracks)

	sort.Sort(byYear(tracks))
	printTracks(tracks)

	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}

		if x.Year != y.Year {
			return x.Year < y.Year
		}

		if x.Length != y.Length {
			return x.Length < y.Length
		}

		return false
	}})
	printTracks(tracks)
}
