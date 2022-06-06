run_proto:
	protoc --go_out=protofile/. --go_opt=module=task1/protofile --go-grpc_out=protofile/. --go-grpc_opt=module=task1/protofile protofile/*proto
run_proto1:
	protoc --go_out=protofile1/. --go_opt=module=task1/protofile1 --go-grpc_out=protofile1/. --go-grpc_opt=module=task1/protofile1 protofile1/*proto
