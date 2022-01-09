package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Canvas struct {
	width  uint
	height uint
	pixel  [][]Color
}

func NewCanvas(width, height uint) Canvas {
	pixel := make([][]Color, width)
	for i := range pixel {
		pixel[i] = make([]Color, height)
	}

	return Canvas{width, height, pixel}
}

func (canvas Canvas) WritePixel(x, y uint, color Color) {
	canvas.pixel[x][y] = color
}

func (canvas Canvas) ToPPM() string {
	const maxCharPerLine = 70
	const maxCharPixel = 5

	data := strings.Builder{}

	header := fmt.Sprintf("P3\n%d %d\n255\n", canvas.width, canvas.height)
	data.WriteString(header)

	for j := uint(0); j < canvas.height; j++ {
		row := ""
		for i := uint(0); i < canvas.width; i++ {
			pixel := []string{
				strconv.Itoa(scaleFloat(canvas.pixel[i][j].red)),
				strconv.Itoa(scaleFloat(canvas.pixel[i][j].green)),
				strconv.Itoa(scaleFloat(canvas.pixel[i][j].blue))}

			for _, color := range pixel {
				if len(row)+maxCharPixel > maxCharPerLine {
					row += "\n"
					data.WriteString(row)
					row = ""
				}

				if row != "" {
					row += " "
				}

				row += color
			}
		}
		row += "\n"
		data.WriteString(row)
	}

	return data.String()
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
