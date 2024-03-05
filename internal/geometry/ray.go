package geometry

type Ray struct {
	Origin    Vec3
	Direction Vec3
}

func CreateRay(origin Vec3, direction Vec3) Ray {
	return Ray{Origin: origin, Direction: direction}
}

/**
 * This function returns the point at a given distance along the ray.
 */
func (r Ray) PointAt(t float64) Vec3 {
	return r.Origin.Add(r.Direction.Mul(t))
}
