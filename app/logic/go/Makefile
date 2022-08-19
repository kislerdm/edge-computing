.DEFAULT_GOAL := help

help: ## Prints help message.
	@ grep -h -E '^[a-zA-Z_.-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[1m%-30s\033[0m %s\n", $$1, $$2}'

test:  ## Runs unit tests.
	@ go test --tags=unittest -v -coverpkg=./... -coverprofile=coverage.out ./...
	@ go tool cover -func coverage.out
	@ rm coverage.out

compile:
	@ docker run --rm \
        -v $(PWD):/src \
        -w="/src" \
        tinygo/tinygo:0.24.0 tinygo build \
                -target=wasm \
                -gc=leaking -opt=2 -no-debug -panic=trap \
            -o assets/logic.wasm main.go

setup.wasm:
	@ docker run --rm \
              -v $(PWD):/src \
              -w="/src" \
              tinygo/tinygo:0.24.0 \
              /bin/bash -c "cp /usr/local/tinygo/targets/wasm_exec.js assets/wasm_exec.js"

setup.js: setup.wasm
	@ docker run -v $(PWD)/assets:/assets -w /assets smithmicro/uglifyjs wasm_exec.js init.js -o logic.js
	@ rm assets/wasm_exec.js

build: compile setup.js  ## Builds logic and js dependencies.
