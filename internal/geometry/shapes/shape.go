package shapes

import (
	. "raytracer/internal/geometry"
	"raytracer/internal/utils"
)

type Intersection struct {
	Dist          float64 `json:"dist" yaml:"dist"`
	Point         Vec3    `json:"point" yaml:"point"`
	Normal        Vec3    `json:"normal" yaml:"normal"`
	OutwardFacing bool    `json:"outward_facing" yaml:"outwardFacing"`
}

type Shape interface {
	Intersect(ray Ray, limit utils.Range) (bool, Intersection)
}
