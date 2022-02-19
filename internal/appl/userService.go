package appl

import (
	"API-RS-TOUKIO/internal/domain/users"
	"API-RS-TOUKIO/internal/infra/data/pgclient"
)

type userServiceImpl struct{}

func (p *userServiceImpl) CreateUser(e *users.Entity) error {
	rep := pgclient.NewUserRepository()
	return rep.CreateUser(e)
}

func NewUserService() users.Service {
	return &userServiceImpl{}
}
