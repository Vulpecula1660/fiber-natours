package member

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/Vulpecula1660/fiber-natours/api/middleware"
	"github.com/Vulpecula1660/fiber-natours/api/protocol"
)

type (
	RegisterInput struct {
		Account  string `json:"account" validate:"required,min=4,max=18"`  // 帳號（長度： 4~18、格式：英文數字）
		Password string `json:"password" validate:"required,min=6,max=30"` // 密碼（長度： 6~30、格式：任意字元）
	}

	registerTask struct {
		APIName string
		Req     *RegisterInput
		Res     *protocol.Response
		Storage *RegisterStorage
	}

	// RegisterStorage : 暫存
	RegisterStorage struct {
		Err error
	}
)

// newRegisterTask : 實例化Task
func newRegisterTask() *registerTask {
	return &registerTask{
		APIName: "Register",
		Req:     &RegisterInput{},
		Res: &protocol.Response{
			Code:    "1",
			Message: "",
			Result:  struct{}{},
		},
		Storage: &RegisterStorage{},
	}
}

// Register : 註冊
func Register(c *fiber.Ctx) error {
	task := newRegisterTask()

	// 解析參數
	if shouldBreak := task.BindRequest(c); shouldBreak {
		return middleware.ErrorHandler(task.Storage.Err)
	}

	ctx, cancel := context.WithTimeout(c.Context(), time.Second*30)
	defer cancel()

	// 註冊會員
	if shouldBreak := task.DoRegister(ctx); shouldBreak {
		return middleware.ErrorHandler(task.Storage.Err)
	}

	return c.JSON(task.Res)
}

// BindRequest : 解析參數
func (task *registerTask) BindRequest(c *fiber.Ctx) bool {
	if err := c.BodyParser(task.Req); err != nil {
		task.Storage.Err = err
		return true
	}

	err := validator.New().Struct(task.Req)
	if err != nil {
		task.Storage.Err = err
		return true
	}

	return false
}

// DoRegister : 註冊會員
func (task *registerTask) DoRegister(ctx context.Context) bool {

	return false
}
