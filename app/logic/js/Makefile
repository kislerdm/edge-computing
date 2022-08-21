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
	@ docker run --rm -v $(PWD)/logic:/src -w /src node:alpine3.16 \
  		/bin/sh -c "npm install uglify-js -g > '/dev/null' 2>&1 && uglifyjs data/*.js *.js -o logic.test.js && node logic.test.js"
	@ [ -f $(PWD)/logic/logic.test.js ] && rm $(PWD)/logic/logic.test.js


build:  ## Builds logic and js dependencies.
	@ test -d assets || mkdir assets
	@ docker run --rm -v $(PWD):/src -w /src smithmicro/uglifyjs logic/data/*.js $(ls logic/*.js | grep -v *.test.js) -o assets/logic.js
