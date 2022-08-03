
.PHONY: all

all: engine_swagger engine_proto fs storage_swagger storage_proto workflow chat bot call_center swagger_mix


.PHONY: engine_swagger

engine_swagger:
	protoc -I/usr/local/include -I./engine -I${GRPC_GATEWAY}/third_party/googleapis  \
      -I${GRPC_GATEWAY} \
      --swagger_out=version=false,json_names_for_fields=false,allow_delete_body=true,include_package_in_tags=false,allow_repeated_fields_in_body=false,fqn_for_swagger_name=false,merge_file_name=engine,allow_merge=true:./swagger \
      ./engine/*.proto

.PHONY: engine_proto

engine_proto:
	protoc -I/usr/local/include -I./engine -I${GRPC_GATEWAY}/third_party/googleapis  \
      -I${GRPC_GATEWAY}  \
      --go-grpc_out=./engine --go-grpc_opt=paths=source_relative \
      --go_opt=paths=source_relative \
      --go_out=./engine ./engine/*.proto

engine_chat_proto:
	protoc -I/usr/local/include -I./engine/chat -I${GRPC_GATEWAY}/third_party/googleapis  \
      -I${GRPC_GATEWAY} \
      --go-grpc_out=./engine/chat --go-grpc_opt=paths=source_relative \
      --go_opt=paths=source_relative \
      --go_out=./engine/chat ./engine/chat/*.proto

engine: engine_proto engine_chat_proto engine_swagger


.PHONY: call_center

call_center:
	protoc -I/usr/local/include -I./cc -I${GRPC_GATEWAY}/third_party/googleapis  \
      -I${GRPC_GATEWAY} \
      --go_opt=paths=source_relative \
      --go-grpc_out=./cc --go-grpc_opt=paths=source_relative \
      --go_out=./cc ./cc/*.proto

.PHONY: fs

fs:
	protoc -I/usr/local/include -I./fs \
	--go_opt=paths=source_relative \
	--go-grpc_out=./fs --go-grpc_opt=paths=source_relative \
	--go_out=./fs ./fs/fs.proto

.PHONY: storage_swagger

storage_swagger:
	protoc -I/usr/local/include -I./storage -I${GRPC_GATEWAY}/third_party/googleapis  \
      -I${GRPC_GATEWAY} \
      -I. \
      --swagger_out=version=false,json_names_for_fields=false,allow_delete_body=true,include_package_in_tags=false,allow_repeated_fields_in_body=false,fqn_for_swagger_name=false,merge_file_name=storage,allow_merge=true:./swagger \
      ./storage/*.proto

.PHONY: storage_proto

storage_proto:
	protoc -I/usr/local/include -I. -I${GRPC_GATEWAY}/third_party/googleapis  \
	  -I${GRPC_GATEWAY}  \
	  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	  --go_opt=paths=source_relative \
      --go_out=. ./storage/*.proto

.PHONY: workflow

workflow:
	protoc -I/usr/local/include -I. -I${GRPC_GATEWAY}/third_party/googleapis  \
      -I${GRPC_GATEWAY}  \
      -I./engine/chat \
      --go-grpc_out=. --go-grpc_opt=paths=source_relative \
      --go_opt=paths=source_relative \
      --go_out=. ./workflow/*.proto

.PHONY: chat

chat:
	protoc \
	-I ./chat \
	--go_opt=paths=source_relative --go_out=chat \
	--go-grpc_out=paths=source_relative:chat \
	\
	./chat/*.proto

	# export GO111MODULE=on  # Enable module mode
	##### require google.golang.org/protobuf v1.25.0
	# go get google.golang.org/protobuf/cmd/protoc-gen-go@v1.25.0 # *.pb.go
	##### require google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.0.1
	# go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.0.1 # *_grpc.pb.go
	##### require github.com/micro/micro/v2 v2.9.1
	# go get github.com/micro/micro/v2/cmd/protoc-gen-micro@v2.9.1 # *.pb.micro.go

.PHONY: bot

bot:
	protoc \
	-I ./bot \
	-I ./chat \
	--go_opt=paths=source_relative --go_out=bot \
	--micro_out=plugins=grpc,paths=source_relative:bot \
	 \
	./bot/*.proto


.PHONY: swagger_mix

swagger_mix:
	swagger-mixin ./swagger/engine.swagger.json ./swagger/storage.swagger.json  > ./swagger/api.json | true


.PHONY: clean

clean:
	rm -rf ./engine/*.go
	rm -rf ./fs/*.go
	rm -rf ./storage/*.go
	rm -rf ./workflow/*.go
	rm -rf ./swagger/*