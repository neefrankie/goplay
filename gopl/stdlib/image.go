package stdlib

import (
	"image"
	"image/color"
)

func createImage() *image.RGBA {
	width := 200
	height := 100

	upLeft := image.Point{X: 0, Y: 0}
	lowRight := image.Point{X: width, Y: height}

	img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

	cyan := color.RGBA{R: 100, G: 200, B: 200, A: 0xff}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case x < width/2 && y < height/2:
				img.Set(x, y, cyan)
			case x >= width/2 && y >= height/2:
				img.Set(x, y, color.White)
			default:

			}
		}
	}

	return img
}
