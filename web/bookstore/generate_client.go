// Run go generate ./... from the root of the repository to generate API client
// OpenAPI 2.0 Specification.
package bookstore

//go:generate npx openapi-typescript@5.4.1 ../../swagger.yaml --output ./api.ts
