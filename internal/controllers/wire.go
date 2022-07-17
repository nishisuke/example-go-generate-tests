//go:build wireinject
// +build wireinject

package controllers

import (
	"sample/internal/usecases"

	"github.com/google/wire"
)

var diUserUsecase = wire.Bind(
	new(UserUsecase),
	new(usecases.UserUsecase),
)
var userUsecaseSet = wire.NewSet(
	usecases.NewUserUsecase,
	diUserUsecase,
)

func NewUserController() UserController {
	wire.Build(
		userUsecaseSet,
		wire.Struct(new(UserController), "*"),
	)
	return UserController{}
}
