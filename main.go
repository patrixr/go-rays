//go:build js && wasm

package main

import (
	"raytracer/internal/utils"
	"raytracer/internal/wasm"
	"syscall/js"
)

func main() {
	c := make(chan bool)

	utils.Log.Info("Raytracer engine started")

	js.Global().Set("RT", map[string]interface{}{
		"renderFromJson": js.FuncOf(wasm.ScheduleRender),
	})

	<-c
}
