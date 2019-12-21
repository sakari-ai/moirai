grpc:
	@sh ./scripts/proto-gen.sh
	@go run ./tools/json2const/main.go proto
	@gofmt -s -w proto