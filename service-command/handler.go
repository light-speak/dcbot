package main

import (
    "errors"
    "github.com/lty5240/dcbot/service-command/proto"
    "golang.org/x/net/context"
)

type service struct {
    repository Repository
}

func (s service) Find(ctx context.Context, req *proto.CommandRequest, res *proto.Response) error {
    command, err := s.repository.FindByCommand(req.Command)
    if err != nil {
        return err
    }
    if command.Status == proto.Command_DISABLE {
        return errors.New("该命令已关闭")
    }
    res.Service = command.Service
    return nil
}

func (s service) Send(ctx context.Context, req *proto.ActionRequest, res *proto.Response) error {
    res.Message = "测试命令是否执行成功"
    return nil
}
