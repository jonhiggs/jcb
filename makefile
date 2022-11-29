.DEFAULT_GOAL = run

test:
	go test ./lib/ui

build:
	go build  -o jcb ./cmd/main.go

run: test build
	./jcb
