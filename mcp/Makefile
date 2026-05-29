BIN := bin/traefik-mcp

# -s -w strips the symbol and DWARF debug tables, which otherwise account for
# ~30% of the binary (the bulk of the rest is the github.com/traefik/traefik/v3
# import graph). Fine for this read-only POC; drop the flags to debug with a
# symbolicated stack trace.
LDFLAGS := -s -w

.PHONY: build test lint

build:
	go build -ldflags="$(LDFLAGS)" -o $(BIN) ./cmd/traefik-mcp

test:
	go test ./...

lint:
	go vet ./...
