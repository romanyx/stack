PKGS := github.com/romanyx/stack
SRCDIRS := $(shell go list -f '{{.Dir}}' $(PKGS))
GO := go
default: test

check: test vet gofmt 

doc:
	@echo GoDoc link: http://localhost:6060/pkg/github.com/romanyx/stack/
	godoc -http=:6060

test:
	$(GO) test $(PKGS)

vet: | test
	$(GO) vet $(PKGS)

bench:
	GOMAXPROCS=1 go test -bench=. -benchmem

errcheck:
	$(GO) get github.com/kisielk/errcheck
	errcheck $(PKGS)

gofmt:
	@echo Checking code is gofmted
	@test -z "$(shell gofmt -s -l -d -e $(SRCDIRS) | tee /dev/stderr)"
