# Go Graphql Example
GraphQL Example Implementation in Go

## Prerequisites
- [https://github.com/pressly/goose](Goose): for database migration
- MySQL: for database engine
## Getting Started
- Run database migration: ` goose -dir ./internal/pkg/db/migrations/mysql mysql "username:password@tcp(localhost:3306)/gql_example?parseTime=true" up`
- Download dependencies: `go mod vendor`
- Run the service: `go run serve.go`
