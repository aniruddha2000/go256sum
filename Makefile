hello:
	echo "Hello"

build:
	go build -o bin/go256sum cmd/go256sum/go256sum.go

run:
	go run cmd/go256sum/go256sum.go
