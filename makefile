run: test build
	./jcb

test:
	go test ./lib/transaction
	go test ./ui/formatter/data
	go test ./ui/formatter/string
	go test ./ui/repeater

build:
	go build  -o jcb ./cmd/main.go

