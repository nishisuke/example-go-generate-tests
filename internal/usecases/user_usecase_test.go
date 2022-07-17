package usecases

import (
	"errors"
	"reflect"
	"sample/internal/models"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUserUsecase_GetMe(t *testing.T) {
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
		// TODO: Add test cases.
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
