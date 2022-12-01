.DEFAULT_GOAL = run

test:
	#go test ./lib/ui
	go test ./ui/formatter/data
	go test ./ui/formatter/string
	go test ./ui/repeater

build:
	go build  -o jcb ./cmd/main.go

run: test build
	./jcb
