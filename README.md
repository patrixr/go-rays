# GoRays

GoRays is a simple ray tracer written in Go.
It is based on the book "Ray Tracing in One Weekend" by Peter Shirley.

```Makefile

tidy:
	go mod tidy

demo:
	go run ./cli.go demo

run:
	go run ./cli.go

wasm:
	GOARCH=wasm GOOS=js go build -o dist/raytracer.wasm ./main.go

build: wasm

```

## Getting Started

To run the ray tracer with a basic preconfigured scene simply run the following command:

```bash
make demo
```

This will generate a file called `output.png` in the root directory of the project.

## Building for WASM

To build the ray tracer for the web run the following command:

```bash
make wasm
```