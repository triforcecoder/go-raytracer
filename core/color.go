package core

type Color struct {
	Red   float64
	Green float64
	Blue  float64
}

var Black = Color{0, 0, 0}
var White = Color{1, 1, 1}
var Red = Color{1, 0, 0}
var Green = Color{0, 1, 0}
var Blue = Color{0, 0, 1}

func NewColor(red, green, blue float64) Color {
	return Color{red, green, blue}
}

func (color Color) Equals(other Color) bool {
	return FloatEquals(color.Red, other.Red) &&
		FloatEquals(color.Green, other.Green) &&
		FloatEquals(color.Blue, other.Blue)
}

func (color Color) Add(other Color) Color {
	return Color{color.Red + other.Red,
		color.Green + other.Green,
		color.Blue + other.Blue}
}

func (color Color) Subtract(other Color) Color {
	return Color{color.Red - other.Red,
		color.Green - other.Green,
		color.Blue - other.Blue}
}

func (color Color) MultiplyScalar(scalar float64) Color {
	return Color{color.Red * scalar,
		color.Green * scalar,
		color.Blue * scalar}
}

func (color Color) Multiply(other Color) Color {
	return Color{color.Red * other.Red,
		color.Green * other.Green,
		color.Blue * other.Blue}
}
