.PHONY: run
run:
	go run -race ./cmd/main.go

.PHONY: test
test:
	go test ./...

.PHONY: build
build:
	make proto
	go build -o go-messenger ./cmd/main.go

.PHONY: proto
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	./proto/protos.proto
