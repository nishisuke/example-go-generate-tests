package usecases

import (
	"errors"
	"reflect"
	"sample/internal/models"
	"sample/internal/usecases/mock_usecases"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUserUsecase_GetMe(t *testing.T) {
	validUser := models.User{Name: "Valid"}
	type fields struct {
		userRepo UserRepo
	}
	type args struct {
		id int
	}
	tests := []struct {
		name       string
		initialize func(ctrl *gomock.Controller) UserUsecase
		args       args
		want       models.User
		wantErr    error
	}{
		{
			"case1",
			func(ctrl *gomock.Controller) UserUsecase {
				m := mock_usecases.NewMockUserRepo(ctrl)
				m.EXPECT().Find(validUser.ID).Return(validUser)
				return NewMockUserUsecase(m)
			},
			args{validUser.ID},
			validUser,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			u := tt.initialize(ctrl)
			got, err := u.GetMe(tt.args.id)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("UserUsecase.GetMe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserUsecase.GetMe() = %v, want %v", got, tt.want)
			}
		})
	}
}
