module github.com/webitel/protos/storage

go 1.18

replace (
	github.com/webitel/protos/engine => ./engine
)

require (
	google.golang.org/genproto v0.0.0-20220801145646-83ce21fca29f // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)