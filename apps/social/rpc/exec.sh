goctl rpc protoc ./social.proto --go_out=. --go-grpc_out=. --zrpc_out=.

goctl model mysql ddl -src="./deploy/sql/social.sql" -dir="./apps/social/socialmodels/" -c
goctl api go -api apps/social/api/social.api -dir apps/social/api -style gozero
