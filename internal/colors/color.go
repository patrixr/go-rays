package colors

import (
	"math"
	"raytracer/internal/utils"
)

type Color struct {
	R uint8 `yaml:"r"`
	G uint8 `yaml:"g"`
	B uint8 `yaml:"b"`
	A uint8 `yaml:"a"`
}

// ---------------------------------------------------

func White() Color {
	return Color{255, 255, 255, 255}
}

// ---------------------------------------------------

func MetalGray() Color {
	return Color{204, 204, 204, 255}
}

// ---------------------------------------------------

func TeaGreen() Color {
	return Color{204, 213, 174, 255}
}

// ---------------------------------------------------

func Black() Color {
	return Color{0, 0, 0, 255}
}

// ---------------------------------------------------

func (c Color) Multiply(other Color) Color {
	return Color{
		R: uint8(float64(c.R) * float64(other.R) / 255),
		G: uint8(float64(c.G) * float64(other.G) / 255),
		B: uint8(float64(c.B) * float64(other.B) / 255),
		A: c.A,
	}
}

// ---------------------------------------------------

func (c Color) MultiplyScalar(scalar float64) Color {
	return Color{
		R: uint8(float64(c.R) * scalar),
		G: uint8(float64(c.G) * scalar),
		B: uint8(float64(c.B) * scalar),
		A: c.A,
	}
}

// ---------------------------------------------------

func (c Color) WithGammaCorrection() Color {
	return Color{
		R: uint8(256 * utils.Clamp(math.Sqrt(float64(c.R)/256), 0, 1)),
		G: uint8(256 * utils.Clamp(math.Sqrt(float64(c.G)/256), 0, 1)),
		B: uint8(256 * utils.Clamp(math.Sqrt(float64(c.B)/256), 0, 1)),
		A: c.A,
	}
}

// ---------------------------------------------------

func (c Color) Lerp(other Color, t float64) Color {
	return Color{
		R: uint8(float64(c.R)*(1-t) + float64(other.R)*t),
		G: uint8(float64(c.G)*(1-t) + float64(other.G)*t),
		B: uint8(float64(c.B)*(1-t) + float64(other.B)*t),
		A: uint8(float64(c.A)*(1-t) + float64(other.A)*t),
	}
}

// ---------------------------------------------------

func (c Color) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8
	g = uint32(c.G)
	g |= g << 8
	b = uint32(c.B)
	b |= b << 8
	a = uint32(c.A)
	a |= a << 8
	return
}

// ---------------------------------------------------

type ColorAverage struct {
	R, G, B, A float64
	count      int
}

func CreateColorAverage() *ColorAverage {
	return &ColorAverage{}
}

func (c *ColorAverage) Add(color Color) {
	c.R += float64(color.R)
	c.G += float64(color.G)
	c.B += float64(color.B)
	c.A += float64(color.A)
	c.count++
}

func (c *ColorAverage) Color() Color {
	return Color{
		R: uint8(c.R / float64(c.count)),
		G: uint8(c.G / float64(c.count)),
		B: uint8(c.B / float64(c.count)),
		A: uint8(c.A / float64(c.count)),
	}
}
