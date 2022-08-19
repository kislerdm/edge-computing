.DEFAULT_GOAL := help

help: ## Prints help message.
	@ grep -h -E '^[a-zA-Z_.-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[1m%-30s\033[0m %s\n", $$1, $$2}'

train: ## Re-trains the model.
	@ docker-compose run train "python main.py"

APPS_DIR := ./app

TECHNOLOGY := js

test: ## Runs unit tests for the app developed using specified technology.
	@ cd $(APPS_DIR)/logic/$(TECHNOLOGY)/ && make test

build: ## Builds the app using specified technology.
	@ cd $(APPS_DIR)/logic/$(TECHNOLOGY)/ && make build

publish: ## Publishes artifacts.
	@ $(foreach dir, $(wildcard $(APPS_DIR)/logic/*), $(subst $(dir),$1,$(APPS_DIR)/logic);)
