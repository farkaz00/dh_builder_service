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

Go service (`github.com/farkaz00/dh_builder_service`) structured in three layers:

```
dhbuilder/                        — service layer (business logic)
├── interfaces.go                 — DHServicer and DHDAO interfaces
├── service.go                    — DHService struct, deps, constructor
├── service_card.go               — card operations (CreateCard, UpdateCard, GetCard, GetCards)
├── service_login.go              — login operation
└── models/
    └── models_card.go            — Card struct and CardRealm type

dhbuilder_dao/                    — DAO layer (data access)
├── constructor.go                — NewDHDAO factory (DAOType-based)
└── csv/                          — CSV implementation of DHDAO
    ├── dao_csv.go                — DHCSV struct, deps, constructor
    └── dao_csv_card.go           — card persistence (SaveCard, GetCard, GetCards)

dhbuilder_server/
└── rest/                         — REST layer (HTTP transport)
    ├── handlers.go               — HTTP handler functions
    ├── endpoints.go              — endpoint wiring
    ├── decoders.go               — request decoders
    └── encoders.go               — response encoders

main.go                           — entrypoint
```

### Layer conventions

- Each domain operation gets its own file suffixed by entity (e.g. `service_card.go`, `dao_csv_card.go`)
- Structs follow the `DH<Impl>` pattern (`DHService`, `DHCSV`) with a matching `DH<Impl>Deps` for constructor dependencies
- Constructors return the interface, not the concrete type
- Errors are logged via `zap` with a method-scoped logger before being returned: `logger := dhs.logger.With(zap.String("method", "..."))`
- Sentinel errors are declared as package-level `var` in the relevant service file
