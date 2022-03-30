package member

import (
	"context"
	"fmt"

	"github.com/Vulpecula1660/fiber-natours/enum"
	"github.com/Vulpecula1660/fiber-natours/model/dao/user"
	"github.com/Vulpecula1660/fiber-natours/model/dto"
	"github.com/Vulpecula1660/fiber-natours/util"

	"github.com/google/uuid"
)

type CreateInput struct {
	Account  string // 帳號
	Password string // 密碼
}

// Create : 註冊
func Create(ctx context.Context, input *CreateInput) error {
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
		return enum.AccountUsed
	}

	// 密碼加密
	password := util.HashAndSalt([]byte(input.Password))

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
