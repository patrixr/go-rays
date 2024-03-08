package utils

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func ExportImage(image image.Image, path string) {
	ext := strings.ToLower(filepath.Ext(path))

	if ext != ".png" && ext != ".jpg" && ext != ".jpeg" {
		Log.Fatal("Unsupported file format " + ext)
	}

	os.MkdirAll(filepath.Dir(path), os.ModePerm)

	file, err := os.Create(path)

	Boom(err)

	defer file.Close()

	if ext == ".png" {
		Boom(png.Encode(file, image))
	} else {
		Boom(jpeg.Encode(file, image, nil))
	}

	Log.Info("Image exported to " + path)
}

func ImageToBytes(img image.Image) []byte {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	rgba := make([]byte, 0, width*height*4)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			color := img.At(x, y)
			r, g, b, a := color.RGBA()
			// The RGBA() method returns colors in the range [0, 65535].
			// Convert it to [0, 255].
			rgba = append(rgba, byte(r>>8), byte(g>>8), byte(b>>8), byte(a>>8))
		}
	}

	return rgba
}
