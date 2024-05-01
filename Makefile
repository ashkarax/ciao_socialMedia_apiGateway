protoc:
	protoc --go_out=. --go-grpc_out=. ./pkg/auth_svc/infrastructure/pb/*.proto
server:
	go run cmd/main.go