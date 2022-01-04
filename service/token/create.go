package token

import (
	"context"
	"fmt"
	"time"

	"github.com/Vulpecula1660/fiber-natours/model/redis"
	"github.com/google/uuid"
)

type CreateInput struct {
	UserID string // 會員ID
}

func Create(ctx context.Context, input *CreateInput) error {
	if input == nil || input.UserID == "" {
		return fmt.Errorf("參數錯誤")
	}

	token := uuid.New().String()

	err := redis.Set(
		ctx,
		token,
		input.UserID,
		time.Hour,
	)
	if err != nil {
		return err
	}

	err = redis.Set(
		ctx,
		input.UserID,
		token,
		time.Hour,
	)
	if err != nil {
		return err
	}

	return nil
}
