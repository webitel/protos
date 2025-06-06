
.PHONY: all

all: engine_swagger engine_proto fs storage_swagger storage_proto logger_swagger logger_proto workflow chat bot call_center swagger_mix


.PHONY: engine_swagger

engine_swagger:
	protoc -I/usr/local/include -I./engine -I./ -I./annotations \
      --swagger_out=simple_operation_ids=true,disable_default_errors=true,version=false,json_names_for_fields=false,allow_delete_body=true,include_package_in_tags=false,allow_repeated_fields_in_body=false,fqn_for_swagger_name=false,merge_file_name=engine,allow_merge=true:./swagger \
      ./engine/*.proto

.PHONY: engine_proto

engine_proto:
	protoc -I/usr/local/include -I./engine -I./annotations  \
      --go-grpc_out=./engine --go-grpc_opt=paths=source_relative \
      --go_opt=paths=source_relative \
      --go_out=./engine ./engine/*.proto

engine_chat_proto:
	protoc -I/usr/local/include -I./engine/chat -I./  \
      --go-grpc_out=./engine/chat --go-grpc_opt=paths=source_relative \
      --go_opt=paths=source_relative \
      --go_out=./engine/chat ./engine/chat/*.proto

engine: engine_proto engine_chat_proto engine_swagger


.PHONY: call_center

call_center:
	protoc -I/usr/local/include -I./cc -I./ -I./engine \
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
	protoc -I/usr/local/include -I./storage -I./engine -I./annotations \
      -I. \
      --swagger_out=version=false,json_names_for_fields=false,allow_delete_body=true,include_package_in_tags=false,allow_repeated_fields_in_body=false,fqn_for_swagger_name=true,merge_file_name=storage,allow_merge=true:./swagger \
      ./storage/*.proto

.PHONY: storage_proto

storage_proto:
	protoc -I/usr/local/include -I./  -I./storage \
	  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	  --go_opt=paths=source_relative \
      --go_out=. ./storage/*.proto

.PHONY: workflow

workflow:
	protoc -I/usr/local/include -I./  \
      -I./engine/chat \
      --go-grpc_out=. --go-grpc_opt=paths=source_relative \
      --go_opt=paths=source_relative \
      --go_out=. ./workflow/*.proto

.PHONY: logger_swagger

logger_swagger:
	protoc -I/usr/local/include -I./logger -I./  \
      --swagger_out=version=false,json_names_for_fields=false,allow_delete_body=true,include_package_in_tags=false,allow_repeated_fields_in_body=false,fqn_for_swagger_name=false,merge_file_name=logger,allow_merge=true:./swagger \
      ./logger/*.proto

.PHONY: logger_proto

logger_proto:
	protoc -I/usr/local/include -I./logger -I./  \
      --go-grpc_out=./logger --go-grpc_opt=paths=source_relative \
      --go_opt=paths=source_relative \
      --go_out=./logger ./logger/*.proto

.PHONY: chat

chat:
	protoc \
	-I chat -I . \
	--go_opt=paths=source_relative --go_out=chat \
	--go-grpc_out=paths=source_relative:chat \
	\
	chat/*.proto

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
	--go-grpc_out=paths=source_relative:bot \
	 \
	./bot/*.proto

.PHONY: messages

messages:
	protoc -I . \
	--go_opt=paths=source_relative --go_out=. \
	--go-grpc_out=paths=source_relative:. \
	\
	chat/messages/*.proto

.PHONY: messages_swagger

messages_swagger:
	protoc -I . \
	--openapiv2_out=merge_file_name=messages,allow_merge=true,openapi_naming_strategy=fqn,json_names_for_fields=false,disable_default_errors=true,repeated_path_param_separator=csv,logtostderr=true:swagger \
	chat/messages/openapiv2.proto \
	chat/messages/catalog.proto


.PHONY: contacts_proto

contacts_proto:
	protoc -I . \
      --go_opt=paths=source_relative --go_out=. \
      --go-grpc_out=. --go-grpc_opt=paths=source_relative \
      --micro_out=plugins=grpc,paths=source_relative:. \
      gateway/contacts/*.proto

.PHONY: swagger_mix

swagger_mix:
	go run github.com/msample/swagger-mixin@latest ./swagger/engine.swagger.json ./swagger/storage.swagger.json \
 	 ./swagger/messages.swagger.json ./swagger/logger.swagger.json ./swagger/contacts.swagger.json ./swagger/webitel-go.swagger.json \
 	 ./swagger/wfm.swagger.json ./swagger/knowledgebase.swagger.json ./swagger/cases.swagger.json ./swagger/custom.swagger.json ./swagger/fts.swagger.json > ./swagger/api.json | true


.PHONY: clean

clean:
	rm -rf ./engine/*.go
	rm -rf ./fs/*.go
	rm -rf ./storage/*.go
	rm -rf ./workflow/*.go
	rm -rf ./swagger/*