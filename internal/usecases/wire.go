//go:build wireinject
// +build wireinject

package usecases

import (
	"sample/internal/repositories"
	"sample/internal/usecases/mock_usecases"

	"github.com/google/wire"
)

var diUserRepo = wire.Bind(
	new(UserRepo),
	new(repositories.UserRepository),
)

var diMockUserRepo = wire.Bind(
	new(UserRepo),
	new(*mock_usecases.MockUserRepo),
)

var userRepoSet = wire.NewSet(
	wire.Struct(new(repositories.UserRepository), "*"),
	diUserRepo,
)

func NewUserUsecase() UserUsecase {
	wire.Build(
		userRepoSet,
		wire.Struct(new(UserUsecase), "*"),
	)
	return UserUsecase{}
}

func NewMockUserUsecase(m *mock_usecases.MockUserRepo) UserUsecase {
	wire.Build(
		wire.Struct(new(UserUsecase), "*"),
		diMockUserRepo,
	)
	return UserUsecase{}
}
