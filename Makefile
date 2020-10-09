
.PHONY: all

all: engine_swagger engine_proto fs storage_swagger storage_proto swagger_mix


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
      --go_out=plugins=grpc:./engine ./engine/*.proto

.PHONY: fs

fs:
	protoc --go_out=plugins=grpc:./ ./fs/fs.proto

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

.PHONY: swagger_mix

swagger_mix:
	swagger mixin ./swagger/engine.swagger.json ./swagger/storage.swagger.json  > ./swagger/api.json | true


.PHONY: clean

clean:
	rm -rf ./engine/*.go
	rm -rf ./fs/*.go
	rm -rf ./storage/*.go
	rm -rf ./swagger/*