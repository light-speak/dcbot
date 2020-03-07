module github.com/lty5240/dcbot/service-command

go 1.13

replace github.com/golang/lint => golang.org/x/lint v0.0.0-20200130185559-910be7a94367

replace github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.3.0

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1

require (
	github.com/golang/protobuf v1.3.3
	github.com/jinzhu/gorm v1.9.12
	github.com/micro/go-micro/v2 v2.1.2
	github.com/micro/go-plugins/registry/etcdv3/v2 v2.0.3
	golang.org/x/net v0.0.0-20200222125558-5a598a2470a0
)
