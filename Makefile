start:
	go run ./cmd/shortener

flags:
	./shortener -a="localhost:3000" -b="http://localhost:3000" -f="tmp/tmp/file.json"

build:
	rm -f shortener && go build -v ./cmd/shortener

test:
	go test ./...
