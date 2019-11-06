GOCMD=go
GOTEST=$(GOCMD) test
GOTOOL=$(GOCMD) tool
GOPHERBADGER=$(HOME)/go/bin/gopherbadger

.PHONY: all test coverage funcoverage

all: test coverage funcoverage

test:
	$(GOTEST) -v -cover -count=1 -mod=vendor

coverage:
	$(GOPHERBADGER) -md="README.md" -png=false

funcoverage:
	$(GOTEST) -mod=vendor -coverprofile=coverage.out && $(GOTOOL) cover -func=coverage.out
