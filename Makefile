PACKAGES := $(shell go list ./...)
name := $(shell basename ${PWD})

all: help

.PHONY: help
help: Makefile
	@echo
	@echo " Choose a make command to run"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

## init: initialize project (make init module=github.com/user/project)
.PHONY: init
init:
	go mod init ${module}
	go install github.com/cosmtrek/air@latest
	asdf reshim golang

## vet: vet code
.PHONY: vet
vet:
	go vet $(PACKAGES)

## templ: generate templ templates
.PHONY: templ
templ:
	templ generate

## templ: watch generate templ templates
.PHONY: templ-watch
templ-watch:
	templ generate -watch

## test: run unit tests
.PHONY: test
test:
	go test -race -cover $(PACKAGES)

## build: build a binary
.PHONY: build
build:
	go build -v -o ./tmp/main .

## docker-build: build project into a docker container image
.PHONY: docker-build
docker-build: test
	GOPROXY=direct docker build -t ${name} .

## docker-run: run project in a container
.PHONY: docker-run
docker-run:
	docker run -it --rm -p 8080:8080 ${name}

## start: build and run local project
.PHONY: start
start:
	air

## css: build tailwindcss
.PHONY: css
css:
	npx tailwindcss -i assets/styles/input.css -o assets/styles/main.css --minify

## css-watch: watch build tailwindcss
.PHONY: css-watch
css-watch:
	npx tailwindcss -i assets/styles/input.css -o assets/styles/main.css --watch

## new-migration: create new tern migration
.PHONY: new-migration
new-migration:
	TERN_MIGRATIONS=db/migrations tern new migration ${name}

## db-migrate: run tern migrations
.PHONY: db-migrate
db-migrate:
	TERN_MIGRATIONS=db/migrations tern migrate

## fly-db-proxy: start fly.io db proxy
.PHONY: fly-db-proxy
fly-db-proxy:
	fly proxy 5434:5432 -a song-contest-rater-service-db

## deploy: deploy to fly.io
.PHONY: deploy
deploy:
	fly deploy