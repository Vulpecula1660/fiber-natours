package member

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/Vulpecula1660/fiber-natours/api/middleware"
	"github.com/Vulpecula1660/fiber-natours/api/protocol"
	"github.com/Vulpecula1660/fiber-natours/enum"
	"github.com/Vulpecula1660/fiber-natours/service/member"
)

type (
	CreateInput struct {
		Account  string `json:"account" validate:"required,min=4,max=18"`  // 帳號（長度： 4~18）
		Password string `json:"password" validate:"required,min=6,max=30"` // 密碼（長度： 6~30）
	}

	createTask struct {
		APIName string
		Req     *CreateInput
		Res     *protocol.Response
		Storage *CreateStorage
	}

	// CreateStorage : 暫存
	CreateStorage struct {
		Err error
	}
)

// newCreateTask : 實例化Task
func newCreateTask() *createTask {
	return &createTask{
		APIName: "Create",
		Req:     &CreateInput{},
		Res: &protocol.Response{
			Code:    "1",
			Message: "OK",
			Result:  struct{}{},
		},
		Storage: &CreateStorage{},
	}
}

// Create : 新增
func Create(c *fiber.Ctx) error {
	task := newCreateTask()

	// 解析參數
	if shouldBreak := task.BindRequest(c); shouldBreak {
		return middleware.ErrorHandler(task.Storage.Err)
	}

	ctx, cancel := context.WithTimeout(c.Context(), time.Minute*15)
	defer cancel()

	// 新增會員
	if shouldBreak := task.DoCreate(ctx); shouldBreak {
		return middleware.ErrorHandler(task.Storage.Err)
	}

	return c.JSON(task.Res)
}

// BindRequest : 解析參數
func (task *createTask) BindRequest(c *fiber.Ctx) bool {
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

// DoCreate : 新增會員
func (task *createTask) DoCreate(ctx context.Context) bool {
	err := member.Create(
		ctx,
		&member.CreateInput{
			Account:  task.Req.Account,
			Password: task.Req.Password,
		},
	)
	if err != nil {
		task.Storage.Err = err
		return true
	}

	return false
}
