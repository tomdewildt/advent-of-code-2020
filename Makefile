.PHONY: init clean run test build lint
.DEFAULT_GOAL := help

NAMESPACE := tomdewildt
NAME := advent-of-code-2020

help: ## Show this help
	@echo "${NAMESPACE}/${NAME}"
	@echo
	@fgrep -h "##" $(MAKEFILE_LIST) | \
	fgrep -v fgrep | sed -e 's/## */##/' | column -t -s##

##

init: ## Initialize the environment
	go mod download

clean: ## Clean the environment
	go mod tidy

##

run: ## Run the tool
	go run \
		-ldflags "-X main.Name=${NAME} -X main.Version=0.0.0" \
		cmd/advent-of-code-2020/advent-of-code-2020.go ${day} \
		-f ./assets/day10/day${day}.txt

test: ## Run tests
	go test ./... 

##

build: ## Build the tool
	go build \
		-o ${NAME} \
		-ldflags "-X main.Name=${NAME} -X main.Version=0.0.0" \
		cmd/advent-of-code-2020/advent-of-code-2020.go

##

lint: ## Run lint & syntax check
	go vet ./...
	gofmt -s -w .
