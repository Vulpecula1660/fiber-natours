package member

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"github.com/Vulpecula1660/fiber-natours/enum"
	"github.com/Vulpecula1660/fiber-natours/model/dao/user"
	"github.com/Vulpecula1660/fiber-natours/service/token"
)

type LoginInput struct {
	Account  string // 帳號
	Password string // 密碼
}

// Login : 會員登入
func Login(ctx context.Context, input *LoginInput) (apiToken string, err error) {
	if input == nil {
		return "", fmt.Errorf("參數錯誤")
	}

	// 用帳號搜尋
	res, err := user.Get(ctx,
		&user.GetInput{
			Account: input.Account,
		},
		nil,
	)
	if err != nil {
		return "", err
	}

	if len(res) == 0 {
		return "", &fiber.Error{
			Code:    enum.UserNotFound,
			Message: "找不到使用者",
		}
	}

	// 比對密碼
	if !comparePasswords(res[0].Password, []byte(input.Password)) {
		return "", &fiber.Error{
			Code:    enum.WrongPassword,
			Message: "密碼錯誤",
		}
	}

	// token create
	apiToken, err = token.Create(
		ctx,
		&token.CreateInput{
			UserID: res[0].ID,
		},
	)
	if err != nil {
		return "", err
	}

	return apiToken, nil
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	return err == nil
}
