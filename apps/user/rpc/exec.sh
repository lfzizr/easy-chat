goctl rpc protoc ./user.proto --go_out=. --go-grpc_out=. --zrpc_out=.

goctl model mysql ddl -src="./deploy/sql/user.sql" -dir="./apps/user/models/" -c
goctl api go -api apps/user/api/user.api -dir apps/user/api -style gozero