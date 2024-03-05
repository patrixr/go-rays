package core

import (
	. "raytracer/internal/geometry"
)

type Viewport struct {
	Size int `yaml:"size" json:"size"`
}

type Camera struct {
	LookFrom     Vec3    `yaml:"look_from" json:"lookFrom"`
	LookAt       Vec3    `yaml:"look_at" json:"lookAt"`
	Vup          Vec3    `yaml:"vup" json:"vup"`
	FocusDist    float64 `yaml:"focus_dist" json:"focusDist"`
	DefocusAngle float64 `yaml:"defocus_angle" json:"defocusAngle"`
	VFOV         float64 `yaml:"vfov" json:"VFOV"`
}

func (c *Camera) BasisVectors() (Vec3, Vec3, Vec3) {
	w := c.LookFrom.Sub(c.LookAt).Normalize()
	u := c.Vup.Cross(w).Normalize()
	v := w.Cross(u).Normalize()
	return u, v, w
}
