.DEFAULT_GOAL := help

help: ## Prints help message.
	@ grep -h -E '^[a-zA-Z_.-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[1m%-30s\033[0m %s\n", $$1, $$2}'

train: ## Re-trains the model.
	@ docker-compose run train "python main.py"

APPS_DIR := ./app

TECHNOLOGY := go

generate.dependencies: ## Generates required dependencies, i.e. model and lookup table.
	@ cd $(APPS_DIR)/logic/$(TECHNOLOGY)/ && make generate.dependencies

test: ## Runs unit tests for the app developed using specified technology.
	@ cd $(APPS_DIR)/logic/$(TECHNOLOGY)/ && make test

build: ## Builds the app using specified technology.
	@ cd $(APPS_DIR)/logic/$(TECHNOLOGY)/ && make build

TECHNOLOGIES := go

build-pages: ## Builds the artifacts for release.
	@ for t in $(TECHNOLOGIES) ; do \
	    if [ ! -d public/$${t} ]; then mkdir -p public/$${t}; fi &&\
	        cp -r app/common/* public/$${t}/;\
	        cp -r app/logic/$${t}/assets/* public/$${t}/assets/ ;\
    done

localserver: ## Launch a web server for local tests.
	@ PORT=9090 DIR=$(PWD)/public/ go run --tags=localserver server/main.go
