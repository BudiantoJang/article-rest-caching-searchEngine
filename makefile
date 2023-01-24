run:
	go run cmd/main.go

mock-gen:
	@./scripts/generate-mock.sh port

test:
	go test ./... -short