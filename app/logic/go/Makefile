.DEFAULT_GOAL := help

help: ## Prints help message.
	@ grep -h -E '^[a-zA-Z_.-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[1m%-30s\033[0m %s\n", $$1, $$2}'

generate.dependencies: ## Generates required dependencies, i.e. model and lookup table.
	@ docker run --rm \
        -v $(PWD):/src \
        -w="/src" \
        -v $(PWD)/../../../data/colorsname.csv:/src/codegen/data/colorsname.csv \
        -v $(PWD)/../../../colortypemodel/model/model.json:/src/codegen/data/model.json \
        golang:1.17.7-alpine3.15 \
        /bin/sh -c "cd codegen && go run --tags=gen main.go"

test:  ## Runs unit tests.
	@ docker run --rm \
	    -v $(PWD):/src \
	    -w="/src" \
	    golang:1.17.7-alpine3.15 \
	    /bin/sh -c \
	    "go mod tidy && go test --tags=unittest -v -coverpkg=./... -coverprofile=coverage.out ./... && go tool cover -func coverage.out && rm coverage.out"

compile:
	@ test -d build || mkdir build
	@ docker run --rm \
        -v $(PWD):/src \
        -w="/src" \
        tinygo/tinygo:0.24.0 tinygo build -target=wasm \
            -gc=leaking -opt=2 -no-debug -panic=trap \
                -o build/logic.wasm main.go

setup.wasm:
	@ docker run --rm \
        -v $(PWD):/src \
        -w="/src" \
        tinygo/tinygo:0.24.0 \
        /bin/bash -c \
        "cp /usr/local/tinygo/targets/wasm_exec.js wasm_exec.js"

setup.js: setup.wasm
	@ test -d build || mkdir build
	@ docker run --rm -v $(PWD):/src -w /src smithmicro/uglifyjs wasm_exec.js logic/logic.js -o build/logic.js
	@ rm wasm_exec.js

build: compile setup.js  ## Builds logic and js dependencies.
