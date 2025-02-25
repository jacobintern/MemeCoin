swag:
	swag init -g ./cmd/api/main.go --parseDependency --parseInternal

gen-env:
	cp ./config/.env.default ./config/.env

lint:
	golangci-lint run ./... --timeout=10m