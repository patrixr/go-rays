.PHONY: tidy run build

tidy:
	go mod tidy

demo:
	go run ./cli.go render

render:
	go run ./cli.go render -s $(scene)

help:
	go run ./cli.go render --help

wasm:
	GOARCH=wasm GOOS=js go build -o dist/raytracer.wasm ./main.go

build: wasm
