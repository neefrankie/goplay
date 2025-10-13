package main

import (
	"gopl/ch8/thumbnail"
)

func main() {
	dir := "~/datasource/django-static/covers/nlp-pytorch"

	thumbnail.MakeThumbnails(dir)
}
