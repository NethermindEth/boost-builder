.PHONY: all test clean

GOBIN = ./build/bin
GO ?= latest
GORUN = env GO111MODULE=on go run

all:
	$(GORUN) build/ci.go install
	@echo "Done building."
	@echo "Run \"$(GOBIN)/builder\" to launch the builder."

test: all
	$(GORUN) build/ci.go test

clean:
	env GO111MODULE=on go clean -cache
	rm -fr build/_workspace/pkg/ $(GOBIN)/*