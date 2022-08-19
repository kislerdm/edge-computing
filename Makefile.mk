.DEFAULT_GOAL := help

help: ## Prints help message.
	@ grep -h -E '^[a-zA-Z_.-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[1m%-30s\033[0m %s\n", $$1, $$2}'

TECHNOLOGY := js

build: ## Builds the app using specified technology.
	@ cd ./app/$(TECHNOLOGY)/ && make build

train: ## Re-trains the model.
	@ docker-compose run train "python main.py"
