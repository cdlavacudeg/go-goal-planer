BYNARY_NAME= goal-planer

generate_docs:
	swag init -g cmd/app/main.go

build:
	go mod tidy
	go build -o bin/$(BYNARY_NAME) cmd/app/main.go
