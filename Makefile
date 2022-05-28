run:
	@go run cmd/cli/main.go

gen:
	@go generate ./...

openapi:
	@oapi-codegen -package rippling -generate "types,client,chi-server,spec" ./OASv3/rippling/RipplingOpenAPI.v1.yaml > ./api/rippling/rippling.gen.go

oas3server:
	@go-oas3 -package rippling -path ./server/rippling -swagger-addr "https://stoplight.io/api/v1/projects/rippling/rippling-api/nodes/RipplingOpenAPI.v1.yaml"

tools:
	@go install github.com/mikekonan/go-oas3@latest # install OpenAPIv3 Go Server Generator
