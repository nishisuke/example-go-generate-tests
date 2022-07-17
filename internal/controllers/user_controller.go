package controllers

//go:generate gotests -w -all $GOFILE
//go:generate gomockhandler -config=$SAMPLE_GO_TEST_ROOT/gomockhandler.json -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

import (
	"fmt"
	"net/http"
	"sample/internal/models"
)

type (
	UserController struct {
		userUsecase UserUsecase
	}

	UserUsecase interface {
		GetMe(int) (models.User, error)
	}
)

func (c UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	me, err := c.userUsecase.GetMe(1)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
	} else {
		fmt.Fprintf(w, "%v", me)
	}
}
