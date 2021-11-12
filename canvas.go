package main

import (
	"fmt"
	"strconv"
)

type Canvas struct {
	width  int
	height int
	pixel  [][]Color
}

func NewCanvas(width, height int) Canvas {
	pixel := make([][]Color, width)
	for i := range pixel {
		pixel[i] = make([]Color, height)
	}

	return Canvas{width, height, pixel}
}

func (canvas Canvas) WritePixel(x, y int, color Color) {
	canvas.pixel[x][y] = color
}

func (canvas Canvas) ToPPM() string {
	const maxCharPerLine = 70
	const maxCharPixel = 5

	header := fmt.Sprintf("P3\n%d %d\n255\n", canvas.width, canvas.height)

	data := ""
	for j := 0; j < canvas.height; j++ {
		row := ""
		for i := 0; i < canvas.width; i++ {
			pixel := []string{
				strconv.Itoa(scaleFloat(canvas.pixel[i][j].red)),
				strconv.Itoa(scaleFloat(canvas.pixel[i][j].green)),
				strconv.Itoa(scaleFloat(canvas.pixel[i][j].blue))}

			for _, color := range pixel {
				if len(row)+maxCharPixel > maxCharPerLine {
					row += "\n"
					data += row
					row = ""
				}

				if row != "" {
					row += " "
				}

				row += color
			}
		}
		row += "\n"
		data += row
	}

	return header + data
}

// scale from float 0:1 to int 0:255
func scaleFloat(x float64) int {
	if x <= 0 {
		return 0
	} else if x >= 1 {
		return 255
	} else {
		return int(x * 256)
	}
}
