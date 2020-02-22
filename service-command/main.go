package main

import (
    "fmt"
    "github.com/lty5240/dcbot/service-command/proto"
    "github.com/micro/go-micro/v2/registry"
    "github.com/micro/go-micro/v2/service"
    "github.com/micro/go-micro/v2/service/grpc"
    "github.com/micro/go-plugins/registry/consul/v2"
    "log"
    "os"
)

func main() {
    db, err := CreateConnection()
    if db != nil {
        defer db.Close()
    }
    if err != nil {
        log.Fatalf("无法连接数据库，异常：%v", err)
    }

    db.AutoMigrate(&proto.Command{})

    repository := &CommandRepository{}
    registryServer := os.Getenv("REGISTRY_SERVER")

    reg := consul.NewRegistry(func(opt *registry.Options) {
        opt.Addrs = []string{
            registryServer,
        }
    })
    srv := grpc.NewService(
        service.Registry(reg),
        service.Name("dcbot.service.command"),
    )
    srv.Init()
    proto.RegisterCommandServiceHandler(srv.Server(), &commandService{repository})

    if err := srv.Run(); err != nil {
        fmt.Printf("服务运行异常：%s \n", err)
    }
}
