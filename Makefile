LINT = golangci-lint

OBJ = build/algohub

all: install

install: build

build: test 
	go build -o $(OBJ)

test: lint
	go test ./...

lint: fmt
	$(LINT) run .

fmt:
	go fmt ./...

clean:
	rm -rf build
