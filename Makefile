genProto:
	@protoc --go_out=. --go-grpc_out=. --grpc-gateway_out=. --grpc-gateway_opt generate_unbound_methods=true --openapiv2_out . proto/grpc-gateway.proto
