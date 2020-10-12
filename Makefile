
.PHONY: all

all: engine_swagger engine_proto fs storage_swagger storage_proto workflow chat bot call_center swagger_mix


.PHONY: engine_swagger

engine_swagger:
	protoc -I/usr/local/include -I./engine -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis  \
      -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
      --swagger_out=version=false,json_names_for_fields=false,allow_delete_body=true,include_package_in_tags=false,allow_repeated_fields_in_body=false,fqn_for_swagger_name=false,merge_file_name=engine,allow_merge=true:./swagger \
      ./engine/*.proto

.PHONY: engine_proto

engine_proto:
	protoc -I/usr/local/include -I./engine -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis  \
      -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
      --go_out=plugins=grpc,paths=source_relative:./engine ./engine/*.proto

.PHONY: call_center

call_center:
	protoc -I/usr/local/include -I./cc -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis  \
      -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
      --go_out=plugins=grpc,paths=source_relative:./cc ./cc/*.proto

.PHONY: fs

fs:
	protoc -I/usr/local/include -I./fs --go_out=plugins=grpc,paths=source_relative:./fs ./fs/fs.proto

.PHONY: storage_swagger

storage_swagger:
	protoc -I/usr/local/include -I./storage -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis  \
      -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
      -I. \
      --swagger_out=version=false,json_names_for_fields=false,allow_delete_body=true,include_package_in_tags=false,allow_repeated_fields_in_body=false,fqn_for_swagger_name=false,merge_file_name=storage,allow_merge=true:./swagger \
      ./storage/*.proto

.PHONY: storage_proto

storage_proto:
	protoc -I/usr/local/include -I. -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis  \
      -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
      --go_out=plugins=grpc,paths=source_relative:. ./storage/*.proto

.PHONY: workflow

workflow:
	protoc -I/usr/local/include -I. -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis  \
      -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
      --micro_out=plugins=grpc,paths=source_relative:. \
      --go_out=plugins=grpc,paths=source_relative:. ./workflow/*.proto

.PHONY: chat

chat:
	protoc --go_out=plugins=grpc,paths=source_relative:. --micro_out=plugins=grpc,paths=source_relative:. ./chat/*.proto

.PHONY: bot

bot:
	protoc -Ichat -Ibot --go_out=plugins=grpc,paths=source_relative:bot --micro_out=plugins=grpc,paths=source_relative:bot ./bot/*.proto


.PHONY: swagger_mix

swagger_mix:
	swagger mixin ./swagger/engine.swagger.json ./swagger/storage.swagger.json  > ./swagger/api.json | true


.PHONY: clean

clean:
	rm -rf ./engine/*.go
	rm -rf ./fs/*.go
	rm -rf ./storage/*.go
	rm -rf ./workflow/*.go
	rm -rf ./swagger/*