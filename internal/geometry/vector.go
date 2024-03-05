package geometry

import (
	"math"
	"raytracer/internal/utils"
)

type Vec3 struct {
	X float64 `yaml:"x"`
	Y float64 `yaml:"y"`
	Z float64 `yaml:"z"`
}

func CreateVector(x float64, y float64, z float64) Vec3 {
	return Vec3{X: x, Y: y, Z: z}
}

func CreateNullVector() Vec3 {
	return CreateVector(0, 0, 0)
}

func CreateRandomVector() Vec3 {
	return CreateVector(
		utils.Rand(-1, 1),
		utils.Rand(-1, 1),
		utils.Rand(-1, 1),
	)
}

func CreateRandomVectorInUnitSphere() Vec3 {
	for {
		v := CreateRandomVector().Mul(2).Sub(CreateVector(1, 1, 1))
		if v.LengthSquared() < 1 {
			return v.Normalize()
		}
	}
}

// This function creates a random vector on the hemisphere
// that is facing the normal vector.
func CreateRandomVectorOnHemisphere(normal Vec3) Vec3 {
	inUnitSphere := CreateRandomVectorInUnitSphere()
	if inUnitSphere.Dot(normal) > 0.0 {
		return inUnitSphere
	}
	return inUnitSphere.Mul(-1)
}

func (v Vec3) Add(v2 Vec3) Vec3 {
	return CreateVector(v.X+v2.X, v.Y+v2.Y, v.Z+v2.Z)
}

func (v Vec3) AddX(v2 Vec3) Vec3 {
	return CreateVector(v.X+v2.X, v.Y, v.Z)
}

func (v Vec3) AddY(v2 Vec3) Vec3 {
	return CreateVector(v.X, v.Y+v2.Y, v.Z)
}

func (v Vec3) AddZ(v2 Vec3) Vec3 {
	return CreateVector(v.X, v.Y, v.Z+v2.Z)
}

func (v Vec3) Sub(v2 Vec3) Vec3 {
	return CreateVector(v.X-v2.X, v.Y-v2.Y, v.Z-v2.Z)
}

func (v Vec3) SubX(x float64) Vec3 {
	return CreateVector(v.X-x, v.Y, v.Z)
}

func (v Vec3) SubY(y float64) Vec3 {
	return CreateVector(v.X, v.Y-y, v.Z)
}

func (v Vec3) SubZ(z float64) Vec3 {
	return CreateVector(v.X, v.Y, v.Z-z)
}

func (v Vec3) Mul(scalar float64) Vec3 {
	return CreateVector(v.X*scalar, v.Y*scalar, v.Z*scalar)
}

func (v Vec3) Div(scalar float64) Vec3 {
	return CreateVector(v.X/scalar, v.Y/scalar, v.Z/scalar)
}

func (v Vec3) Dot(v2 Vec3) float64 {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

func (v Vec3) Cross(v2 Vec3) Vec3 {
	return CreateVector(v.Y*v2.Z-v.Z*v2.Y, v.Z*v2.X-v.X*v2.Z, v.X*v2.Y-v.Y*v2.X)
}

func (v Vec3) LengthSquared() float64 {
	return v.Dot(v)
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v Vec3) Normalize() Vec3 {
	return v.Div(v.Length())
}

func (v Vec3) NearZero() bool {
	s := 1e-8
	return math.Abs(v.X) < s && math.Abs(v.Y) < s && math.Abs(v.Z) < s
}

func (v Vec3) Reflect(normal Vec3) Vec3 {
	return v.Sub(normal.Mul(2 * v.Dot(normal)))
}

func (v Vec3) Refract(normal Vec3, etaiOverEtat float64) Vec3 {
	cosTheta := math.Min(v.Mul(-1).Dot(normal), 1.0)
	rOutPerpendicular := v.Add(normal.Mul(cosTheta)).Mul(etaiOverEtat)
	rOutParallel := normal.Mul(-math.Sqrt(math.Abs(1.0 - rOutPerpendicular.LengthSquared())))
	return rOutPerpendicular.Add(rOutParallel)
}

func CreateRandomVectorInUnitDisk() Vec3 {
	for {
		p := CreateVector(utils.Rand(-1, 1), utils.Rand(-1, 1), 0)
		if p.LengthSquared() < 1 {
			return p
		}
	}
}
