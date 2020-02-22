module github.com/lty5240/dcbot/service-command

go 1.13

replace github.com/golang/lint => golang.org/x/lint v0.0.0-20200130185559-910be7a94367

replace github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.3.0

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1

replace github.com/micro/go-micro/cmd => github.com/micro/go-micro/v2 v2.1.0

replace github.com/micro/go-micro/selector => github.com/micro/go-micro/v2 v2.1.0

replace github.com/micro/go-micro/server/debug => github.com/micro/go-micro/v2 v2.1.0

require (
	github.com/golang/protobuf v1.3.3
	github.com/jinzhu/gorm v1.9.12
	github.com/micro/go-grpc v1.0.1
	github.com/micro/go-micro/cmd v0.0.0-00010101000000-000000000000 // indirect
	github.com/micro/go-micro/selector v0.0.0-00010101000000-000000000000 // indirect
	github.com/micro/go-micro/server/debug v0.0.0-00010101000000-000000000000 // indirect
	github.com/micro/go-micro/v2 v2.1.0
	github.com/micro/go-plugins/registry/consul/v2 v2.0.2
	golang.org/x/net v0.0.0-20200219183655-46282727080f
)
