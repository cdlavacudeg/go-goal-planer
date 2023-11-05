BYNARY_NAME= goal-planner

generate_docs:
	swag init -g cmd/app/main.go

build:
	go mod tidy
	go build -o bin/$(BYNARY_NAME) cmd/app/main.go
