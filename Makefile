all: build

build:
	go build -gcflags="all=-N -l" -o gilo .

build-release:
	go build -o gilo .

run: build
	./gilo

clean:
	rm editor.log

test:
	go test ./... -v
