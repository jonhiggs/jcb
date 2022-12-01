.DEFAULT_GOAL = run

test:
	#go test ./lib/ui
	go test ./ui/formatter/data
	go test ./ui/formatter/string

build:
	go build  -o jcb ./cmd/main.go

run: test build
	./jcb
