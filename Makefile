GOCMD=go
GOTEST=$(GOCMD) test

.PHONY: all test

all: test

test:
	$(GOTEST) -v cover -count=1 -mod=vendor
