run: build
	./bin/app

build:
	make tailwind-build
	make templ-generate
	go build -ldflags "-X main.Environment=production" -o ./bin/$(APP_NAME) ./cmd/$(APP_NAME)/main.go

templ-generate:
	templ generate

templ-watch:
	templ generate --watch

tailwind-watch:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch

tailwind-build:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify

dev:
	go build -o ./tmp/$(APP_NAME) ./$(APP_NAME)/main.go && air

vet:
	go vet ./...

.PHONY: run build tailwind-watch tailwind-build templ templ-watch dev