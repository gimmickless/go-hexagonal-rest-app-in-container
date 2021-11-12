# go-hexagonal-rest-app-in-container

A project template for Go REST services.

> Prerequisites:
> Change the name of the service to something unique in `./go.mod`, `./api/openapi.yml`, then fix the errors Go will report.

## Linting

[`golangci-lint`](https://golangci-lint.run) is used as an linter aggregator.

It runs as a [Github action](./.github/workflows/lint.yml) on the main branch. However, to run it locally, you'll need to first [install](https://golangci-lint.run/usage/install/#local-installation) and then run `golangci-lint run` in the main directory.

## See also

- [Hexagonal Architecture](https://netflixtechblog.com/ready-for-changes-with-hexagonal-architecture-b315ec967749)

- [Go Standard Project Layout](https://github.com/golang-standards/project-layout)
