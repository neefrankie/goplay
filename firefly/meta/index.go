package meta

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
)

func PathWithSuffix(path, suffix string) string {
	if path == "" {
		return suffix
	}
	ext := filepath.Ext(path)
	if ext == "" {
		return fmt.Sprintf("%s%s", path, suffix)
	}

	path = strings.TrimSuffix(path, ext)
	return fmt.Sprintf("%s%s", path, suffix)
}

type BasePage struct {
	Slug string `json:"slug"`
	URL  string `json:"url"`
}

type IndexEntry struct {
	ID        string       `json:"id"`
	Slug      string       `json:"slug"`
	URL       string       `json:"url"`
	Href      string       `json:"href"`
	Title     string       `json:"title"`
	Summary   string       `json:"summary"`
	Exercises []BasePage   `json:"exercises"`
	Path      string       `json:"path"`
	Children  []IndexEntry `json:"children"`
}

func (e *IndexEntry) BuildFilename(suffix string, append bool) (string, error) {
	if e.Path == "" {
		return "", fmt.Errorf("path is empty")
	}

	if suffix == "" {
		return e.Path, nil
	}

	if append {
		return fmt.Sprintf("%s%s", e.Path, suffix), nil
	}

	return PathWithSuffix(e.Path, suffix), nil
}

func (e *IndexEntry) MDFilename(append bool) (string, error) {
	return e.BuildFilename(".md", append)
}

func (e *IndexEntry) JSONFilename(append bool) (string, error) {
	return e.BuildFilename(".json", append)
}

func (e *IndexEntry) HTMLFilename(append bool) (string, error) {
	return e.BuildFilename(".html", append)
}

func (e *IndexEntry) HexIDKey() string {
	return e.Path
}

func (e *IndexEntry) WalkCb(depth int, fn func(entry *IndexEntry, depth int)) {
	fn(e, depth)
	for _, child := range e.Children {
		child.WalkCb(depth+1, fn)
	}
}

func (e *IndexEntry) Walk(
	yield func(entry *IndexEntry) bool,
) {
	if !yield(e) {
		return
	}

	for _, child := range e.Children {
		child.Walk(yield)
	}
}

func (e *IndexEntry) WalkWithDepth(
	startDepth int,
) func(yield func(int, *IndexEntry) bool) {
	return func(yield func(int, *IndexEntry) bool) {
		if !yield(startDepth, e) {
			return
		}
		for _, child := range e.Children {
			child.WalkWithDepth(startDepth + 1)(yield)
		}
	}
}

func WalkEntryAsync(ctx context.Context, entries []IndexEntry) <-chan *IndexEntry {
	ch := make(chan *IndexEntry)
	go func() {
		defer close(ch)
		for _, entry := range entries {
			select {
			case <-ctx.Done():
				return
			default:

			}

			if !walkEntry(ctx, &entry, 0, func(e *IndexEntry, depth int) bool {
				select {
				case ch <- e:
					return true
				case <-ctx.Done():
					return false
				}
			}) {
				return
			}
		}
		close(ch)
	}()
	return ch
}

func walkEntry(
	ctx context.Context,
	entry *IndexEntry,
	depth int,
	fn func(entry *IndexEntry, depth int) bool,
) bool {
	if !fn(entry, depth) {
		return false
	}

	for _, child := range entry.Children {
		if !walkEntry(ctx, &child, depth+1, fn) {
			return false
		}
	}

	return true
}

type IndexPage struct {
	ID          string       `json:"id"`
	URL         string       `json:"url"`
	Slug        string       `json:"slug"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Data        []IndexEntry `json:"data"`
}

func (p *IndexPage) WalkEntry(fn func(entry *IndexEntry)) {
	for _, entry := range p.Data {
		entry.WalkCb(0, func(entry *IndexEntry, depth int) {
			fn(entry)
		})
	}
}

func (p *IndexPage) WalkEntryAsync(ctx context.Context) <-chan *IndexEntry {
	return WalkEntryAsync(ctx, p.Data)
}
