package repositories

import (
	"reflect"
	"sample/internal/models"
	"testing"
)

func TestUserRepository_Find(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		r    UserRepository
		args args
		want models.User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := UserRepository{}
			if got := r.Find(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
