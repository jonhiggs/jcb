run: build
	./jcb

test:
	go test ./lib/transaction
	go test ./lib/formatter/data
	go test ./lib/formatter/string
	go test ./lib/repeater

build:
	go build  -o jcb ./cmd/main.go

