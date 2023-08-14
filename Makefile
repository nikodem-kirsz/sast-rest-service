include .env
export

.PHONY: openapi_http
openapi_http:
	@./scripts/openapi-http.sh sast internal/sast/ports ports

.PHONY: lint
lint:
	@go-cleanarch
	@./scripts/lint.sh common
	@./scripts/lint.sh sast

.PHONY: fmt
fmt:
	goimports -l -w internal/