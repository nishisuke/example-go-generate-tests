package usecases

//go:generate gotests -w -all $GOFILE
//go:generate gomockhandler -config=$SAMPLE_GO_TEST_ROOT/gomockhandler.json -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

import "sample/internal/models"

type (
	UserUsecase struct {
		userRepo UserRepo
	}

	UserRepo interface {
		Find(id int) models.User
	}
)

func (u UserUsecase) GetMe(id int) (models.User, error) {
	return u.userRepo.Find(id), nil
}
