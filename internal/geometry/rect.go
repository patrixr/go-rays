package geometry

type Rect struct {
	X float64 `yaml:"x"`
	Y float64 `yaml:"y"`
	W float64 `yaml:"w"`
	H float64 `yaml:"h"`
}

func CreateRect(x float64, y float64, w float64, h float64) Rect {
	return Rect{X: x, Y: y, W: w, H: h}
}
