# Bookstore

TODO: general description, goals, etc.

## Architecture

TODO: architecture goes here

## Client/UI

TODO: document here what we are doing for the client side

## API

We implement API using [OpenAPI 2.0 (Swagger)](https://swagger.io/specification/v2/) specification.
This implementation has been chosen because of the following:

- it is design for the given domain and the problem
- initial development team has more experience with this API type amongst viable alternatives

Alternatives considered:

- GraphQL
- GRPC

### Code Generation from API Specification

Both Client and Server are auto-generated from `swagger.yaml` OpenAPI specification in the root
of the project. To make changes to the API, edit API schema file and then execute command
from the root of the project:

```sh
go generate ./...
```

### Local Development & Debugging

To run API server, run and/or debug `api/cmd/bookstore-server/main.go`:

```sh
go run ./api/cmd/bookstore-server/main.go
```

Follow this documentation to implement and configure API: https://goswagger.io/go-swagger/generate/server/. Actual API implementations have to be hooked up in `api/restapi/configure_bookstore.go`.

## Open Questions

- Should we implement a more general invntory management solution instead? it seems
that we would also need to manage candles, supplies, etc.
- Hosting where?


