grpc:
	@sh ./scripts/proto-gen.sh
	@gofmt -s -w proto