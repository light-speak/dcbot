package main

import (
    "github.com/jinzhu/gorm"
    "github.com/lty5240/dcbot/service-command/proto"
)

type Repository interface {
    FindByCommand(command string) (*proto.Command, error)
}

type CommandRepository struct {
    db *gorm.DB
}

func (repository CommandRepository) FindByCommand(command string) (*proto.Command, error) {
    cmd := &proto.Command{}
    cmd.Command = command
    if err := repository.db.First(cmd).Error; err != nil {
        return nil, err
    }
    return cmd, nil
}
