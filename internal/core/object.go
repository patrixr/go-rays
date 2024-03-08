package core

import (
	"encoding/json"
	. "raytracer/internal/geometry"
	. "raytracer/internal/geometry/shapes"
	"raytracer/internal/utils"

	"github.com/mitchellh/mapstructure"
)

type World struct {
	Objects []Object `yaml:"objects" json:"objects"`
}

type Object struct {
	Shape    Shape    `json:"shape" yaml:"shape"`
	Material Material `json:"material" yaml:"material"`
}

// -----------------------------------------

func (world *World) Intersect(ray Ray, limit utils.Range) (bool, Intersection, *Object) {
	hitAnything := false
	closestSoFar := limit.Max
	intersection := Intersection{}
	var hitObject *Object

	for index, obj := range world.Objects {
		hit, tempIntersection := obj.Shape.Intersect(ray, limit)

		if hit && tempIntersection.Dist < closestSoFar {
			hitAnything = true
			closestSoFar = tempIntersection.Dist
			intersection = tempIntersection
			hitObject = &world.Objects[index]
		}
	}

	return hitAnything, intersection, hitObject
}

// -----------------------------------------
// Marshalling
// -----------------------------------------

func (s *World) UnmarshalYAML(unmarshal func(interface{}) error) error {
	objs := make([]map[string]interface{}, 0, 50)

	if err := unmarshal(&objs); err != nil {
		return err
	}

	for _, o := range objs {
		obj := Object{}
		obj.Shape = decodeShape(o)
		obj.Material = decodeMaterial(o)
		s.Objects = append(s.Objects, obj)
	}

	return nil
}

func (s *World) UnmarshalJSON(data []byte) error {
	objs := make([]map[string]interface{}, 0, 50)

	if err := json.Unmarshal(data, &objs); err != nil {
		return err
	}

	for _, o := range objs {
		obj := Object{}
		obj.Shape = decodeShape(o)
		obj.Material = decodeMaterial(o)
		s.Objects = append(s.Objects, obj)
	}

	return nil
}

// -----------------------------------------
// Helpers
// -----------------------------------------

func decodeShape(data map[string]interface{}) Shape {
	shapeType := data["type"]

	if shapeType == "sphere" {
		sphere := Sphere{}
		utils.Boom(
			mapstructure.Decode(data, &sphere),
		)
		return &sphere
	}

	utils.Log.Fatal("Unknown shape type ", shapeType)
	return nil
}

func decodeMaterial(data map[string]interface{}) Material {
	materialType := data["material"]

	if materialType == "matte" {
		mat := MatteMaterial{}
		utils.Boom(
			mapstructure.Decode(data, &mat),
		)
		return &mat
	} else if materialType == "transparent" {
		transparent := Transparent{}
		utils.Boom(
			mapstructure.Decode(data, &transparent),
		)
		return &transparent
	} else if materialType == "metal" {
		metal := Metal{}
		utils.Boom(
			mapstructure.Decode(data, &metal),
		)
		return &metal
	}

	utils.Log.Fatal("Unknown material ", materialType)
	return nil
}
