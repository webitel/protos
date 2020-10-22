module github.com/webitel/protos/pkg/bot

go 1.14

replace (
	github.com/coreos/etcd => go.etcd.io/etcd v0.5.0-alpha.5.0.20200425165423-262c93980547
	go.etcd.io/etcd => go.etcd.io/etcd v0.5.0-alpha.5.0.20200425165423-262c93980547
)

require (
	github.com/golang/protobuf v1.4.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/webitel/protos/pkg/chat v0.0.0-20201022103932-569f11b53c46
	google.golang.org/grpc v1.26.0
	google.golang.org/protobuf v1.22.0
	go.etcd.io/etcd v0.5.0-alpha.5.0.20200425165423-262c93980547 // indirect
)
