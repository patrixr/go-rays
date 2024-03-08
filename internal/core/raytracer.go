package core

import (
	"image"
	"math"
	"raytracer/internal/colors"
	. "raytracer/internal/geometry"
	"raytracer/internal/utils"
	"sync"
)

type AspectRatio struct {
	W int `yaml:"w" json:"w"`
	H int `yaml:"h" json:"h"`
}

type Antialiasing struct {
	Samples int `yaml:"samples" json:"samples"`
}

type Raytracer struct {
	Size       int          `yaml:"size" json:"size"`
	Viewport   Viewport     `yaml:"viewport" json:"viewport"`
	Camera     Camera       `yaml:"camera" json:"camera"`
	World      World        `yaml:"objects" json:"objects"`
	Threads    int          `yaml:"threads" json:"threads"`
	Background colors.Color `yaml:"background" json:"background"`

	Aspect AspectRatio `yaml:"aspect_ratio" json:"aspectRatio"`

	Antialiasing Antialiasing `yaml:"antialiasing" json:"antialiasing"`
}

// -----------------------------------------

func (rt *Raytracer) ScreenWidth() int {
	return rt.Size
}

// -----------------------------------------

func (rt *Raytracer) ScreenHeight() int {
	return (rt.Size * rt.Aspect.H) / rt.Aspect.W
}

// -----------------------------------------

func (rt *Raytracer) ViewPortHeight() float64 {
	theta := rt.Camera.VFOV * (math.Pi / 180)
	h := math.Tan(theta / 2)
	return 2 * h * rt.Camera.FocusDist
}

// -----------------------------------------

func (rt *Raytracer) FocalLength() float64 {
	return rt.Camera.LookFrom.Sub(rt.Camera.LookAt).Length()
}

// -----------------------------------------

func (rt *Raytracer) ViewPortWidth() float64 {
	return rt.ViewPortHeight() * (float64(rt.ScreenWidth()) / float64(rt.ScreenHeight()))
}

// -----------------------------------------
// Main method to render the scene
// -----------------------------------------

func (rt *Raytracer) Render() image.Image {
	screenW := rt.ScreenWidth()
	screenH := rt.ScreenHeight()
	img := image.NewNRGBA(image.Rect(0, 0, screenW, screenH))
	pImg := utils.Protect(img)
	sectionHeight := screenH / rt.Threads

	wg := sync.WaitGroup{}
	wg.Add(rt.Threads)
	for i := 0; i < rt.Threads; i++ {
		go rt.RenderSection(pImg, &wg, Rect{0, float64(i * sectionHeight), float64(screenW), float64(sectionHeight)})
	}
	wg.Wait()
	return img
}

func (rt *Raytracer) RenderSection(pImg utils.Protected[*image.NRGBA], wg *sync.WaitGroup, area Rect) {
	defer wg.Done()

	utils.Log.Info("RenderSection", "area", area)

	u, v, w := rt.Camera.BasisVectors()

	// Calculate the Vec3s across the horizontal and down the vertical viewport edges.
	viewportU := u.Mul(rt.ViewPortWidth())
	viewportV := v.Mul(-1).Mul(rt.ViewPortHeight())

	// Calculate the horizontal and vertical delta Vec3s from pixel to pixel.
	pixelDeltaU := viewportU.Div(float64(rt.ScreenWidth()))
	pixelDeltaV := viewportV.Div(float64(rt.ScreenHeight()))

	viewportUpperLeft := rt.Camera.LookFrom.
		Sub(w.Mul(rt.Camera.FocusDist)). // distance between the camera and the viewport
		Sub(viewportU.Div(2)).           // move to the left
		Sub(viewportV.Div(2))            // move up

	rootPixel := viewportUpperLeft.Add(pixelDeltaU.Add(pixelDeltaV).Div(2))

	defocusRadius := rt.Camera.FocusDist * math.Tan(utils.DegreesToRadians(rt.Camera.DefocusAngle)/2.0)
	defocusDiskU := u.Mul(defocusRadius)
	defocusDiskV := v.Mul(defocusRadius)

	farRight := int(area.X + area.W)
	bottom := int(area.Y + area.H)

	for y := int(area.Y); y < bottom; y++ {
		for x := int(area.X); x < farRight; x++ {
			avg := colors.CreateColorAverage()

			realPixelLocation := rootPixel.
				Add(pixelDeltaU.Mul(float64(x))).
				Add(pixelDeltaV.Mul(float64(y)))

			for s := 0; s < rt.Antialiasing.Samples; s++ {
				// Generate randomized variance to the pixel location per sample
				variance := CreateVector(
					utils.Rand(-0.5, 0.5)*pixelDeltaU.X,
					utils.Rand(-0.5, 0.5)*pixelDeltaV.Y,
					0,
				)

				// Add the variance to the real pixel location
				pixelLocation := realPixelLocation.Add(variance)

				var rayOrigin Vec3

				if rt.Camera.DefocusAngle <= 0 {
					rayOrigin = rt.Camera.LookFrom
				} else {
					// Returns a random point in the camera defocus disk
					p := CreateRandomVectorInUnitDisk()
					rayOrigin = rt.Camera.LookFrom.
						Add(defocusDiskU.Mul(p.X)).
						Add(defocusDiskV.Mul(p.Y))
				}

				ray := CreateRay(rayOrigin, pixelLocation.Sub(rt.Camera.LookFrom))

				avg.Add(rt.RayColor(ray, 10))
			}

			img := pImg.Take()
			img.Set(x, y, avg.Color().WithGammaCorrection())
			pImg.Release()
		}
	}
}

// -----------------------------------------

func (rt *Raytracer) RayColor(ray Ray, maxDepth int) colors.Color {

	if maxDepth <= 0 {
		return colors.Black()
	}

	limit := utils.CreateRange(0.0001, math.Inf(1))
	if intersects, hitpoint, obj := rt.World.Intersect(ray, limit); intersects {
		mat := obj.Material

		if scattered, bouncedRay := mat.Scatter(ray, hitpoint); scattered {
			//return rt.RayColor(bouncedRay, maxDepth-1).MultiplyScalar(0.5)
			return rt.RayColor(bouncedRay, maxDepth-1).Multiply(mat.AttenuationColor())
		}
		return colors.Black()
	}

	return rt.defaultRayColor(ray)
}

// -----------------------------------------

func (rt *Raytracer) defaultRayColor(ray Ray) colors.Color {
	unitDirection := ray.Direction.Normalize()
	t := 0.5 * (unitDirection.Y + 1.0)
	return colors.White().Lerp(rt.Background, t)
}

// -----------------------------------------
// Debug
// -----------------------------------------

func (rt *Raytracer) PrintDetails() {
	utils.Log.Info("Dimensions", "w", rt.ScreenWidth(), "h", rt.ScreenHeight())
	utils.Log.Info("Viewport", "w", rt.ViewPortWidth(), "h", rt.ViewPortHeight())
	utils.Log.Info("Focal Length", "size", rt.FocalLength)
	utils.Log.Info("Camera",
		"origin", rt.Camera.LookFrom,
		"up", rt.Camera.Vup,
		"look_at", rt.Camera.LookAt,
		"vfov", rt.Camera.VFOV)
	utils.Log.Info("Aspect Ratio", "w", rt.Aspect.W, "h", rt.Aspect.H)
	utils.Log.Info("Objects", "count", len(rt.World.Objects))
	utils.Log.Info("Antialiasing", "samples", rt.Antialiasing.Samples)
}

// -----------------------------------------
// Configuration
// -----------------------------------------

func (rt Raytracer) WithSamples(samples int) *Raytracer {
	rt.Antialiasing.Samples = samples
	return &rt
}

func (rt Raytracer) WithObjects(objects ...Object) *Raytracer {
	rt.World.Objects = objects
	return &rt
}
