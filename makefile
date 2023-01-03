run: build
	./jcb -f ./test.db

test:
	go test ./lib/formatter/data
	go test ./lib/formatter/string
	go test ./lib/repeater
	go test ./ui/input-bindings/input-bindings_test.go
	go test ./ui/acceptanceFunction/acceptanceFunction_test.go

release: release/jcb_darwin_amd64 \
         release/jcb_linux_amd64 \
         release/jcb_openbsd_amd64

release/%: export GOOS = $(shell echo $* | cut -d_ -f2)
release/%: export GOARCH = $(shell echo $* | cut -d_ -f3)
release/%: export CGO_ENABLED := 1
release/%:
	mkdir -p release
	go build -o release/jcb_${GOOS}_${GOARCH} ./cmd/main.go

tag_release: version = $(shell cat ./config/config.go | grep VERSION | cut -d\  -f4 | tr -d '"')
tag_release:
	gh release create v$(version) ./release/*

build: test
	go build -o jcb ./cmd/main.go
