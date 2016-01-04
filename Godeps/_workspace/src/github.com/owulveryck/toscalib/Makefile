GOPATH=$(HOME)/GOPROJECTS
GO=go
GOFMT=gofmt -w=true
GOBINDATA=$(HOME)/GOPROJECTS/bin/go-bindata

all: test build

build: *.go format
	$(GO) build
	
format: 
	$(GOFMT) *.go

test: *.go
	$(GO) test -coverprofile=coverage.out 
clean:
	rm tosca

normative_definitions.go: NormativeTypes/*
	$(GOBINDATA) -pkg=toscalib -prefix=NormativeTypes/ -o normative_definitions.go NormativeTypes/

