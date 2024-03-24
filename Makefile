.PHONY: tidy run build tag wasm

version := $(shell cat VERSION)

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

tag:
	git tag -a v$(version) -m "Release v$(version)"
	git push origin v$(version)
