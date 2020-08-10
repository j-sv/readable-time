PREFIX ?= /usr/local/bin
BINARY ?= readable-time

all: $(BINARY)

.PHONY: $(BINARY)
$(BINARY):
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags netgo -ldflags '-w -extldflags "-static"' -o $@ main.go

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: test
test:
	go test -short -race ./...

clean:
	@rm readable-time || true

.PHONY: install
install:
	install $(BINARY) $(PREFIX)/$(BINARY)
