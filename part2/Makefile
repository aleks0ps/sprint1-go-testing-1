SRC := main.go
MODNAME := github.com/aleks0ps/sprint1-go-testing-1

.PHONY: all
all: build

go.mod:
	go mod init $(MODNAME)

.PHONY: build
build: go.mod
	go build ./

.PHONY: test
test:
	go get github.com/stretchr/testify/assert
	go test

.PHONY: clean
clean:
	find ./ -mindepth 1 -maxdepth 1 -executable -delete
