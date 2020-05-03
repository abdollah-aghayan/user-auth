package logic

import (
	"context"
	"user-auth/repository"
)

const (
	MOGODB = "mongodb"
	SQL    = "sql"
)

type factory interface {
	GetUserRepo(ctx context.Context) repository.UserRepositoryInterface
}

//Factory production factory
type Factory struct {
	userRepo repository.UserRepositoryInterface
}

//NewFactory new factory
func NewFactory() *Factory {
	return &Factory{}
}

//GetUserRepo create new user repo
func (p *Factory) GetUserRepo(ctx context.Context) repository.UserRepositoryInterface {
	return repository.GetDB()
}
