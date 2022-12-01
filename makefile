.DEFAULT_GOAL = run

test:
	#go test ./lib/ui
	go test ./lib/ui/formatter/data
	go test ./lib/ui/formatter/string

build:
	go build  -o jcb ./cmd/main.go

run: test build
	./jcb
