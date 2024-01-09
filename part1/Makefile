SRC := main.go
MODNAME := github.com/aleks0ps/sprint1-go-testing-1

.PHONY: all
all: build

go.mod:
	go mod init $(MODNAME)

.PHONY: build
build: go.mod
	go build ./
