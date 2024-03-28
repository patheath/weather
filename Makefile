build:
	go build -o bin/main main.go
run:
	go run main.go
test:
	go test -coverpkg=./... ./...
test-build: build
	bin/main