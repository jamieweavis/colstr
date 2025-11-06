.PHONY: build run

run:
	go run ./cmd/colstr/main.go

build:
	go build -o ./bin/colstr cmd/colstr/main.go

package:
	tar -czf ./bin/colstr-mac.tar.gz ./bin/colstr

check-sum:
	shasum -a 256 ./bin/colstr-mac.tar.gz

clean:
	rm -rf ./bin/

.DEFAULT_GOAL := build
