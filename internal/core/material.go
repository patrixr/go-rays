package core

import (
	"math"
	"raytracer/internal/colors"
	. "raytracer/internal/geometry"
	. "raytracer/internal/geometry/shapes"
	"raytracer/internal/utils"
)

type Material interface {
	Scatter(ray Ray, hit Intersection) (
		success bool,
		scatteredRay Ray,
	)

	AttenuationColor() colors.Color
}

// DiffusedMaterial
// ------------------------------------------------

type MatteMaterial struct {
	Color colors.Color
}

func (m *MatteMaterial) AttenuationColor() colors.Color {
	return m.Color
}

func (m *MatteMaterial) Scatter(ray Ray, hit Intersection) (
	success bool,
	scatteredRay Ray,
) {
	bounceDirection := hit.Normal.Normalize().Add(CreateRandomVector().Normalize())

	if bounceDirection.NearZero() {
		bounceDirection = hit.Normal
	}

	scatteredRay = CreateRay(hit.Point, bounceDirection)
	success = true
	return
}

// Metal
// ------------------------------------------------

type Metal struct {
	Color colors.Color
	Fuzz  float64
}

func (m *Metal) AttenuationColor() colors.Color {
	return m.Color
}

func (m *Metal) Scatter(ray Ray, hit Intersection) (
	success bool,
	scatteredRay Ray,
) {
	fuzz := utils.Clamp(m.Fuzz, 0.0, 1.0)

	reflected := ray.Direction.Normalize().Reflect(hit.Normal)
	scatteredRay = CreateRay(hit.Point, reflected.Add(CreateRandomVector().Mul(fuzz)))
	success = true
	return
}

// Transparent
// ------------------------------------------------

type Transparent struct {
	Color      colors.Color
	Refraction float64
}

func (m *Transparent) AttenuationColor() colors.Color {
	return m.Color
}

func (m *Transparent) Scatter(ray Ray, hit Intersection) (
	success bool,
	scatteredRay Ray,
) {
	var refractionRatio float64

	if hit.OutwardFacing {
		refractionRatio = 1.0 / m.Refraction
	} else {
		refractionRatio = m.Refraction
	}

	dir := ray.Direction.Normalize()

	cosTheta := math.Min(dir.Mul(-1).Dot(hit.Normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)
	cannotRefract := refractionRatio*sinTheta > 1.0

	var outDir Vec3

	if cannotRefract || reflectance(cosTheta, refractionRatio) > utils.Rand(0, 1) {
		outDir = dir.Reflect(hit.Normal)
	} else {
		outDir = dir.Refract(hit.Normal, refractionRatio)
	}

	scatteredRay = CreateRay(hit.Point, outDir)
	success = true
	return
}

func reflectance(cosine float64, refIdx float64) float64 {
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}
