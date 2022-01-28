run: build
	./sample

build:
	go build -o sample ./cmd/web
