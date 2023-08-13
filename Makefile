include .env
export

.PHONY: openapi_http
openapi_http:
	@./scripts/openapi-http.sh sast internal/sast/ports ports