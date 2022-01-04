package member

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/Vulpecula1660/fiber-natours/enum"
	"github.com/Vulpecula1660/fiber-natours/model/dao/user"
	"github.com/Vulpecula1660/fiber-natours/model/dto"
)

type RegisterInput struct {
	Account  string // 帳號
	Password string // 密碼
}

// Register : 註冊
func Register(ctx context.Context, input *RegisterInput) error {
	if input == nil {
		return fmt.Errorf("參數錯誤")
	}

	// 查看帳號是否使用過
	total, err := user.GetCount(ctx,
		&user.GetInput{
			Account: input.Account,
		},
		nil,
	)
	if err != nil {
		return err
	}

	if total > 0 {
		return &fiber.Error{
			Code:    enum.AccountUsed,
			Message: "帳號已被使用",
		}
	}

	// 密碼加密
	password := hashAndSalt([]byte(input.Password))

	// 產生UUID
	uuid := uuid.New().String()

	// 新增資料
	err = user.Ins(
		ctx,
		nil,
		&dto.User{
			ID:       uuid,
			Account:  input.Account,
			Password: password,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func hashAndSalt(pwd []byte) string {
	hash, _ := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	return string(hash)
}
