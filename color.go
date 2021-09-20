package main

type Color struct {
	red   float64
	green float64
	blue  float64
}

func (color Color) equals(other Color) bool {
	return equal(color.red, other.red) &&
		equal(color.green, other.green) &&
		equal(color.blue, other.blue)
}

func (color Color) add(other Color) Color {
	return Color{color.red + other.red,
		color.green + other.green,
		color.blue + other.blue}
}

func (color Color) subtract(other Color) Color {
	return Color{color.red - other.red,
		color.green - other.green,
		color.blue - other.blue}
}

func (color Color) multiplyScalar(scalar float64) Color {
	return Color{color.red * scalar,
		color.green * scalar,
		color.blue * scalar}
}

func (color Color) multiply(other Color) Color {
	return Color{color.red * other.red,
		color.green * other.green,
		color.blue * other.blue}
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
