proto:
	@protoc --go_out=./pkg/protogo --go_opt=paths=source_relative \
		--go-grpc_out=./pkg/protogo --go-grpc_opt=paths=source_relative \
		--proto_path=./../../proto \
		--go_opt=Msession.proto=./protogo --go-grpc_opt=Msession.proto=./protogo \
		session.proto

di:
	@cd ./internal/di && \
	wire gen