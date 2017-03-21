ALL_SOURCES := $(shell find . -name '*.go' -not -name '*_test.go' -not -path './vendor/*')
SOURCES := $(shell find . -maxdepth 1 -name '*.go' -not -name '*_test.go')
PACKAGES := $(shell go list ./... | grep -v /vendor/)

all: exe

exe: $(SOURCES)
	go build

dist: $(SOURCES)
	go build -ldflags="-s -w"
	upx-ucl --best kahn

run: $(SOURCES)
	go run -race $(SOURCES)

test: $(SOURCES)
	go test -race -v $(PACKAGES)

integration: $(SOURCES)
	go test -race -tags=integration -v $(PACKAGES)

lint:
	# gometalinter.v1 --install
	gometalinter.v1 $(SOURCES); exit 0

cloc: $(ALL_SOURCES)
	@cloc $(ALL_SOURCES)

fmt: $(ALL_SOURCES)
	@gofmt -w $(ALL_SOURCES)

todo:
	@grep --color -Ion '@\?\(TODO\|XXX\).*' -r . --exclude=Makefile \
		--exclude-dir=vendor

# vim: syntax=make ts=4 sw=4 sts=4 sr noet
