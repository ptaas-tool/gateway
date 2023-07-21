package request

import (
	"fmt"

	"github.com/automated-pen-testing/api/pkg/models/user"

	"github.com/automated-pen-testing/api/pkg/enum"
)

type UserRegisterRequest struct {
	Name string `json:"username"`
	Pass string `json:"password"`
}

type UserUpdateRequest struct {
	UserID uint      `json:"user_id"`
	Role   enum.Role `json:"role"`
}

func (u UserRegisterRequest) Validate() error {
	if len(u.Name) == 0 {
		return fmt.Errorf("username cannot be empty")
	}

	if len(u.Pass) == 0 {
		return fmt.Errorf("password cannot be empty")
	}

	return nil
}

func (u UserRegisterRequest) ToModel() *user.User {
	return &user.User{
		Username: u.Name,
		Password: u.Pass,
	}
}
