package shapes

import (
	"math"
	. "raytracer/internal/geometry"
	"raytracer/internal/utils"
)

type Sphere struct {
	Center Vec3    `yaml:"center"`
	Radius float64 `yaml:"radius"`
}

func (s *Sphere) Intersect(ray Ray, limit utils.Range) (bool, Intersection) {
	oc := ray.Origin.Sub(s.Center)
	a := ray.Direction.Dot(ray.Direction)
	b := 2.0 * oc.Dot(ray.Direction)
	c := oc.Dot(oc) - s.Radius*s.Radius
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return false, Intersection{}
	}

	point := CreateNullVector()
	t := 0.0

	if t = (-b - math.Sqrt(discriminant)) / (2.0 * a); utils.Surrounded(t, limit.Min, limit.Max) {
		// Closest point fo the sphere
		point = ray.PointAt(t)
	} else if t = (-b + math.Sqrt(discriminant)) / (2.0 * a); utils.Surrounded(t, limit.Min, limit.Max) {
		// Furthest point of the sphere, we might be inside of it
		point = ray.PointAt(t)
	} else {
		return false, Intersection{}
	}

	n := s.OutwardNormal(point)
	outFacing := ray.Direction.Dot(n) < 0

	if !outFacing {
		n = n.Mul(-1)
	}

	return true, Intersection{
		Point:         point,
		Normal:        n,
		OutwardFacing: outFacing,
		Dist:          t,
	}
}

func (s *Sphere) OutwardNormal(contactPoint Vec3) Vec3 {
	return contactPoint.Sub(s.Center).Div(s.Radius)
}
