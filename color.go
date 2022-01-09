package main

type Color struct {
	red   float64
	green float64
	blue  float64
}

var black = Color{0, 0, 0}
var white = Color{1, 1, 1}

func (color Color) Equals(other Color) bool {
	return floatEquals(color.red, other.red) &&
		floatEquals(color.green, other.green) &&
		floatEquals(color.blue, other.blue)
}

func (color Color) Add(other Color) Color {
	return Color{color.red + other.red,
		color.green + other.green,
		color.blue + other.blue}
}

func (color Color) Subtract(other Color) Color {
	return Color{color.red - other.red,
		color.green - other.green,
		color.blue - other.blue}
}

func (color Color) MultiplyScalar(scalar float64) Color {
	return Color{color.red * scalar,
		color.green * scalar,
		color.blue * scalar}
}

func (color Color) Multiply(other Color) Color {
	return Color{color.red * other.red,
		color.green * other.green,
		color.blue * other.blue}
}
