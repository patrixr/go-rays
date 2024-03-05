package core

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"raytracer/internal/colors"
	"raytracer/internal/geometry"
	"raytracer/internal/geometry/shapes"
	"raytracer/internal/utils"

	"gopkg.in/yaml.v2"
)

func InitRaytracerFromScene(path string) Raytracer {
	var rt Raytracer

	data, err := os.ReadFile(path)

	utils.PanicOnError(err)

	// check if json or yaml file
	ext := filepath.Ext(path)

	switch ext {
	case ".yaml", ".yml":
		err = yaml.Unmarshal(data, &rt)
	case ".json":
		err = json.Unmarshal(data, &rt)
	default:
		utils.PanicOnError(errors.New("Invalid file format: " + ext))
	}

	utils.PanicOnError(err)

	return rt
}

func InitRaytracerFromJSON(data []byte) Raytracer {
	var rt Raytracer

	err := json.Unmarshal(data, &rt)
	utils.PanicOnError(err)

	return rt
}

func InitRaytracer() *Raytracer {
	return &Raytracer{
		Size: 800,
		Viewport: Viewport{
			Size: 2.0,
		},
		Camera: Camera{
			LookFrom:     geometry.CreateVector(0, 0, 1),
			LookAt:       geometry.CreateVector(0, 0, -1),
			Vup:          geometry.CreateVector(0, 1, 0),
			FocusDist:    1,
			DefocusAngle: 0,
			VFOV:         70,
		},
		World: World{
			Objects: []Object{
				{
					Material: &Metal{
						Color: colors.MetalGray(),
						Fuzz:  0,
					},
					Shape: &shapes.Sphere{
						Center: geometry.CreateVector(0, 0, -1),
						Radius: 0.5,
					},
				},
				{
					Material: &MatteMaterial{
						Color: colors.TeaGreen(),
					},
					Shape: &shapes.Sphere{
						Center: geometry.CreateVector(0, -100.5, -1),
						Radius: 100,
					},
				},
			},
		},
		Threads: 4,
		Aspect: AspectRatio{
			W: 16,
			H: 9,
		},
		Antialiasing: Antialiasing{
			Samples: 50,
		},
	}
}
