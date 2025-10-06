package meta_test

import (
	"context"
	"encoding/json"
	"path/filepath"
	"testing"

	"github.com/youkre/firefly/meta"
)

var toc = `
{
    "id": "crs-5bd2c24bd901b5a9",
    "url": "https://calculusmadeeasy.org/",
    "slug": "calculus-made-easy",
    "title": "Calculus Made Easy",
    "description": "",
    "data": [
        {
            "id": "artl-662caa141e1055c3",
            "slug": "1",
            "url": "https://calculusmadeeasy.org/1.html",
            "href": "1.html",
            "title": "I. To Deliver You From The Preliminary Terrors",
            "summary": "",
            "exercises": [
                {
                    "slug": "1",
                    "url": ""
                }
            ],
            "path": "1",
            "children": [
				{
					"id": "artl-25d704fe889cf1c4",
					"slug": "3",
					"url": "https://calculusmadeeasy.org/3.html",
					"href": "3.html",
					"title": "III. On Relative Growings",
					"summary": "",
					"exercises": [
						{
							"slug": "3",
							"url": ""
						}
					],
					"path": "3",
					"children": []
				}
			]
        },
		{
            "id": "artl-69b61e849f4eaca5",
            "slug": "2",
            "url": "https://calculusmadeeasy.org/2.html",
            "href": "2.html",
            "title": "II. On Different Degrees of Smallness",
            "summary": "",
            "exercises": [],
            "path": "2",
            "children": []
        }
	]
}
`

func TestFilePath(t *testing.T) {
	p := "ch5/ch5-03"

	base := filepath.Base(p)
	t.Logf("Base: %s", base)

	dir := filepath.Dir(p)
	t.Logf("Dir: %s", dir)

	ext := filepath.Ext(p)
	t.Logf("Ext: %s", ext)

	joined := filepath.Join(p, ".md")
	t.Logf("Join: %s", joined)
}

func TestPathWithSuffix(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		path   string
		suffix string
		want   string
	}{
		{
			name:   "Replace suffix",
			path:   "ch5/ch5-03.md",
			suffix: ".html",
			want:   "ch5/ch5-03.html",
		},
		{
			name:   "Add suffix",
			path:   "ch5/ch5-03",
			suffix: ".md",
			want:   "ch5/ch5-03.md",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := meta.PathWithSuffix(tt.path, tt.suffix)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("PathWithSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseIndexPage(t *testing.T) {
	var idxPage meta.IndexPage

	err := json.Unmarshal([]byte(toc), &idxPage)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%v", idxPage)
}

func TestIndexEntry_Iter(t *testing.T) {
	var idxPage meta.IndexPage

	err := json.Unmarshal([]byte(toc), &idxPage)
	if err != nil {
		t.Fatal(err)
	}

	var entryCh = make(chan *meta.IndexEntry)

	go func() {
		idxPage.WalkEntry(func(entry *meta.IndexEntry) {
			entryCh <- entry
		})
		close(entryCh)
	}()

	for entry := range entryCh {
		t.Logf("%v", entry)
	}
}

func TestIndexEntry_WalkEntryAsync(t *testing.T) {
	var idxPage meta.IndexPage

	err := json.Unmarshal([]byte(toc), &idxPage)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	entryCh := idxPage.WalkEntryAsync(ctx)

	for entry := range entryCh {
		t.Logf("%v", entry)
	}
}

func TestIndexEntry_Walk(t *testing.T) {
	var idxPage meta.IndexPage

	err := json.Unmarshal([]byte(toc), &idxPage)
	if err != nil {
		t.Fatal(err)
	}

	for _, entry := range idxPage.Data {
		for entry := range entry.Walk {
			t.Logf("%v", entry)
		}
	}
}

func TestIndexEntry_WalkWithDepth(t *testing.T) {
	var idxPage meta.IndexPage

	err := json.Unmarshal([]byte(toc), &idxPage)
	if err != nil {
		t.Fatal(err)
	}

	for _, entry := range idxPage.Data {
		for depth, entry := range entry.WalkWithDepth(0) {
			t.Logf("%d: %v", depth, entry)
		}
	}
}
