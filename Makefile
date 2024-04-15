build:
	go build -o bin/weather main.go
run:
	go run main.go
test:
	go test -coverpkg=./... ./...
integ-test:
	go test -run Integration
unit-test:
	go test -short ./...
