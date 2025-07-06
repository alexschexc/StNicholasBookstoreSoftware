# Bookstore

TODO: general description, goals, etc.

## Architecture

TODO: architecture goes here

## Client/UI

We implement client as a Next.js application, and it is implemented in `web/bookstore`.
Tutorial: https://nextjs.org/learn/dashboard-app/getting-started.

### Prerequisites

- (Node.js)[https://nodejs.org/en/download/current]
- Package manager:

```sh
npm install -g pnpm
```

Light example generating and consuming API: https://medium.com/@youry.stancatte/generating-typescript-interfaces-from-swagger-1910cc7a726a.

## API

We implement API using [OpenAPI 2.0 (Swagger)](https://swagger.io/specification/v2/) specification. This approach has been chosen because:

- code generation for both: client and server
- RESTful is a widely used API protocol
- development team has more experience with this API type amongst viable alternatives

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


