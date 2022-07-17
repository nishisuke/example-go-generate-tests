package controllers

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUserController_ServeHTTP(t *testing.T) {
	type fields struct {
		userUsecase UserUsecase
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name       string
		initialize func(ctrl *gomock.Controller) UserController
		args       args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			c := tt.initialize(ctrl)
			c.ServeHTTP(tt.args.w, tt.args.r)
		})
	}
}
