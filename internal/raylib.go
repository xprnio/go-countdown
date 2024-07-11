package internal

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type DrawableDuration time.Duration
type UniformFloat float64

type TextRectangle struct {
	rl.Rectangle
	Text string
}

func (f UniformFloat) String(max int) string {
	return fmt.Sprintf("%02d", int(f)%max)
}

func (d DrawableDuration) duration() time.Duration {
	return time.Duration(d)
}

func (d DrawableDuration) Hours() UniformFloat {
	return UniformFloat(d.duration().Hours())
}

func (d DrawableDuration) Minutes() UniformFloat {
	return UniformFloat(d.duration().Minutes())
}

func (d DrawableDuration) Seconds() UniformFloat {
	return UniformFloat(d.duration().Seconds())
}

func (d DrawableDuration) Rectangles(center rl.Vector2) []TextRectangle {
	var rectangles []TextRectangle

	blocks := []string{
		d.Hours().String(24),
		d.Minutes().String(60),
		d.Seconds().String(60),
	}

	var numChars int
	for i, str := range blocks {
		numChars += len(str)
		if i+1 < len(blocks) {
			// account for divider
			numChars++
		}
	}

	windowHeight := rl.GetScreenHeight()
	blockSize := rl.MeasureTextEx(rl.GetFontDefault(), "0", float32(windowHeight)/1.5, 0)
	bounds := rl.NewRectangle(
		center.X-(blockSize.X*float32(numChars))/2,
		center.Y-(blockSize.Y/2),
		float32(numChars)*blockSize.X,
		blockSize.Y,
	)

	var offset int
	for i, str := range blocks {
		rectangles = append(rectangles, TextRectangle{
			Text: str,
			Rectangle: rl.NewRectangle(
				bounds.X+float32(offset)*blockSize.X,
				bounds.Y,
				blockSize.X*2,
				blockSize.Y,
			),
		})

		if i+1 < len(blocks) {
			// append divider block
			rectangles = append(rectangles, TextRectangle{
				Text: ":",
				Rectangle: rl.NewRectangle(
					bounds.X+blockSize.X*float32(offset+len(str)),
					bounds.Y,
					blockSize.X,
					blockSize.Y,
				),
			})
		}

		offset += len(str) + 1
	}

	// TODO: hours rect
	// TODO: divider rect
	// TODO: minutes rect
	// TODO: divider rect
	// TODO: seconds rect

	return rectangles
}
