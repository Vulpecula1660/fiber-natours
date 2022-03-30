package member

import (
	"context"
	"time"

	"github.com/Vulpecula1660/fiber-natours/api/middleware"
	"github.com/Vulpecula1660/fiber-natours/api/protocol"
	"github.com/Vulpecula1660/fiber-natours/enum"
	"github.com/Vulpecula1660/fiber-natours/model/dao/user"
	"github.com/Vulpecula1660/fiber-natours/model/dto"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type (
	DeleteInput struct {
		ID string `json:"id" validate:"required"` // 會員ID
	}

	deleteTask struct {
		APIName string
		Req     *DeleteInput
		Res     *protocol.Response
		Storage *DeleteStorage
	}

	// DeleteStorage : 暫存
	DeleteStorage struct {
		Err error
	}
)

// newDeleteTask : 實例化Task
func newDeleteTask() *deleteTask {
	return &deleteTask{
		APIName: "Delete",
		Req:     &DeleteInput{},
		Res: &protocol.Response{
			Code:   "1",
			Result: struct{}{},
		},
		Storage: &DeleteStorage{},
	}
}

// Delete : 刪除
func Delete(c *fiber.Ctx) error {
	task := newDeleteTask()

	// 解析參數
	if shouldBreak := task.BindRequest(c); shouldBreak {
		return middleware.ErrorHandler(task.Storage.Err)
	}

	ctx, cancel := context.WithTimeout(c.Context(), time.Minute*15)
	defer cancel()

	// 刪除會員
	if shouldBreak := task.DoDelete(ctx); shouldBreak {
		return middleware.ErrorHandler(task.Storage.Err)
	}

	return c.JSON(task.Res)
}

// BindRequest : 解析參數
func (task *deleteTask) BindRequest(c *fiber.Ctx) bool {
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

// DoDelete : 刪除會員
func (task *deleteTask) DoDelete(ctx context.Context) bool {
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
	}

	// 刪除會員
	err = user.Del(
		ctx,
		nil,
		&dto.User{
			ID: task.Req.ID,
		},
	)
	if err != nil {
		task.Storage.Err = err
		return true
	}

	return false
}
