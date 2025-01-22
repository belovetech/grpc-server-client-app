.PHONY: gen run

gen:
	@echo "Generating Go proto files"
	protoc --go_out=. --go-grpc_out=. hello.proto


run:
	@echo "Running server"
	go run server.go
