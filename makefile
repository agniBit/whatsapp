run:
	go run cmd/api/main.go

run_test:
	go run cmd/test/main.go

build:
	go build -o bin/api/main cmd/test/main.go
	go build -o bin/api/main cmd/api/main.go