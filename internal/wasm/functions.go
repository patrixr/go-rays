//go:build js && wasm

package wasm

import (
	"encoding/json"
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
	"raytracer/internal/core"
	"raytracer/internal/utils"
	"syscall/js"
)

type RenderJob struct {
	Id         int
	Config     *core.Raytracer
	jsCallback *js.Value
	Status     string
	Bytes      js.Value
	Width      int
	Height     int
}

var jobId int

var renderChannel chan RenderJob

func init() {
	renderChannel = make(chan RenderJob)

	go func() {
		for job := range renderChannel {
			job.Status = "rendering"
			img := job.Config.Render()
			bytes := utils.ImageToBytes(img)
			dst := js.Global().Get("Uint8Array").New(len(bytes))

			if got, want := js.CopyBytesToJS(dst, bytes), len(bytes); got != want {
				utils.Log.Error("Copied %d bytes, wanted %d", got, want)
				job.Status = "error"
			} else {
				job.Status = "done"
				job.Bytes = dst
			}
			job.jsCallback.Invoke(structs.Map(&job))
		}
	}()
}

func ScheduleRender(this js.Value, inputs []js.Value) interface{} {
	configAsString := inputs[0].String()

	// Convert the string to a map
	configMap := make(map[string]interface{})
	json.Unmarshal([]byte(configAsString), &configMap)

	rt := core.InitRaytracer()

	// Convert the map to a RaytracerConfig
	utils.PanicOnError(mapstructure.Decode(configMap, &rt))

	jobId++
	id := jobId

	job := RenderJob{
		Id:         id,
		Config:     rt,
		jsCallback: &inputs[1],
		Status:     "pending",
		Width:      rt.ScreenWidth(),
		Height:     rt.ScreenHeight(),
	}

	renderChannel <- job

	return structs.Map(&job)
}
