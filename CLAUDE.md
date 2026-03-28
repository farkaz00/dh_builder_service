# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Build
go build ./...

# Run
go run main.go

# Test
go test ./...

# Single test
go test ./path/to/package -run TestName
```

## Architecture

This is an early-stage Go service (`github.com/farkaz00/dh_builder_service`) with the following intended structure:

- `dh_builder/` — core builder logic
- `dhbuilder_server/rest/` — REST API handlers
- `main.go` — entrypoint
