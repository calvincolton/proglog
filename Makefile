compile:
	protoc api/v1/*.proto \
	--go_out=. \
	--go-grpc_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	--proto_path=.

.PHONY: test
test:
	go test -v ./...

.PHONY: test/race
test/race:
	go test -race ./...