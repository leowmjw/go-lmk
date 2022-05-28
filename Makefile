run:
	@go run cmd/cli/main.go

gen:
	@go generate ./...

openapi:
	@oapi-codegen -package rippling -generate "types,client,chi-server,spec" ./OASv3/rippling/RipplingOpenAPI.v1.yaml > ./api/rippling/rippling.gen.go
