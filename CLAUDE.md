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

### REST handler pattern

All HTTP handlers are built via `HandlerWrapper` in `handlers.go`, which composes five typed functions:

```
HandlerWrapper(srv, serviceCallerMakerFunc, requestDecoderFunc, responseEncoderFunc, errorHandlingFunc)
```

- `serviceCallerMakerFunction` — takes `DHServicer`, returns a `serviceCallerFunction` (the actual service call)
- `requestDecoderFunction` — parses `*http.Request` → `any`
- `requestEncoderFunction` — writes response to `http.ResponseWriter`
- `errorHandlingFunction` — all errors currently map to HTTP 500 via `ServerErrorEncoder`

New endpoints need a decoder in `decoders.go`, a handler maker in `handlers.go`, and wiring in `endpoints.go`.

### CSV DAO specifics

- Cards are persisted to `cards.csv` in the working directory (hardcoded in `dhbuilder_dao/constructor.go`)
- `SaveCard` is an upsert: updates the matching record if `card.ID` exists, appends otherwise
- IDs are generated in the DAO on `SaveCard` when `card.ID == ""`: 16 random bytes encoded as a 32-char hex string via `crypto/rand`
- `GetCard` returns `(nil, nil)` when not found — not an error; the service layer translates this to `ErrCardNotFound`
- The CSV DAO has no logger (unlike the service layer)

### Known issues

- All REST errors return HTTP 500; no differentiated status codes yet

### Runtime

- Server listens on `:8080` (hardcoded in `main.go`)
- Logger is `zap.NewDevelopment()` (debug level); production logger config is a TODO
