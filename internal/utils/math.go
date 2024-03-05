package utils

import (
	"math"
	"math/rand"
)

func DegreesToRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180.0)
}

func Rand(min, max float64) float64 {
	return min + (max-min)*rand.Float64()
}

func Surrounded(value, min, max float64) bool {
	return value >= min && value <= max
}

func Contained(value, min, max float64) bool {
	return value >= min && value <= max
}

func Clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
