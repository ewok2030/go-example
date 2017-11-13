# go-example

## Description

This example builds a hello world application.

## Build

The base `golang:alpine` Docker image is several hundred MBs in size, to run a tiny application. To make the smallest possible image, we'll build the application first and then copy it into a `scratch` base image.

Compile the application before building the image:

```bash
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/go-example .
```

```bash
docker build -t go-example .
```

## Run

Run the following to run a container that will be removed when the command exits:

```bash
docker run -it --rm -p 8080:8080 go-example
```

Browse to `http://localhost:8080` to test.