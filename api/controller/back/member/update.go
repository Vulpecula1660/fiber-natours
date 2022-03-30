package member

import (
	"context"
	"time"

	"github.com/Vulpecula1660/fiber-natours/api/middleware"
	"github.com/Vulpecula1660/fiber-natours/api/protocol"
	"github.com/Vulpecula1660/fiber-natours/enum"
	"github.com/Vulpecula1660/fiber-natours/model/dao/user"
	"github.com/Vulpecula1660/fiber-natours/model/dto"
	"github.com/Vulpecula1660/fiber-natours/util"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type (
	UpdateInput struct {
		ID          string `json:"id" validate:"required"`                       // 會員ID
		NewPassword string `json:"newPassword" validate:"required,min=6,max=30"` // 新密碼（長度： 6~30）
	}

	updateTask struct {
		APIName string
		Req     *UpdateInput
		Res     *protocol.Response
		Storage *UpdateStorage
	}

	// UpdateStorage : 暫存
	UpdateStorage struct {
		Err error
	}
)

// newUpdateTask : 實例化Task
func newUpdateTask() *updateTask {
	return &updateTask{
		APIName: "Update",
		Req:     &UpdateInput{},
		Res: &protocol.Response{
			Code:    "1",
			Message: "OK",
			Result:  struct{}{},
		},
		Storage: &UpdateStorage{},
	}
}

// Update : 更新
func Update(c *fiber.Ctx) error {
	task := newUpdateTask()

	// 解析參數
	if shouldBreak := task.BindRequest(c); shouldBreak {
		return middleware.ErrorHandler(task.Storage.Err)
	}

	ctx, cancel := context.WithTimeout(c.Context(), time.Minute*15)
	defer cancel()

	// 更新密碼
	if shouldBreak := task.DoUpdate(ctx); shouldBreak {
		return middleware.ErrorHandler(task.Storage.Err)
	}

	return c.JSON(task.Res)
}

// BindRequest : 解析參數
func (task *updateTask) BindRequest(c *fiber.Ctx) bool {
	if err := c.BodyParser(task.Req); err != nil {
		task.Storage.Err = enum.ParameterError(err)
		return true
	}

	err := validator.New().Struct(task.Req)
	if err != nil {
		task.Storage.Err = enum.ParameterError(err)
		return true
	}

	return false
}

// DoUpdate : 更新密碼
func (task *updateTask) DoUpdate(ctx context.Context) bool {
	// 查看會員是否存在
	member, err := user.Get(
		ctx,
		&user.GetInput{
			UserID: task.Req.ID,
		},
		nil)
	if err != nil {
		task.Storage.Err = err
		return true
	}

	if len(member) == 0 {
		task.Storage.Err = enum.UserNotFound
		return true
	}

	// 加密密碼
	password := util.HashAndSalt([]byte(task.Req.NewPassword))

	// 更新密碼
	err = user.Upd(
		ctx,
		nil,
		&dto.User{
			ID:       task.Req.ID,
			Password: password,
		})
	if err != nil {
		task.Storage.Err = err
		return true
	}

	return false
}
