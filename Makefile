build:
	go build -o bin/main main.go
run:
	go run main.go
test:
	go test -coverpkg=./... ./...
integ-test:
	go test -run Integration
unit-test:
	go test -v -short ./...
test-build: build
	bin/main