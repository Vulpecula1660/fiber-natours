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
	LoginInput struct {
		Account  string `json:"account" validate:"required,min=4,max=18"`  // 帳號（長度： 4~18）
		Password string `json:"password" validate:"required,min=6,max=30"` // 密碼（長度： 6~30）
	}

	LoginIOutput struct {
		Token string `json:"token"`
	}

	loginTask struct {
		APIName string
		Req     *LoginInput
		Res     *protocol.Response
		Storage *LoginStorage
	}

	// LoginStorage : 暫存
	LoginStorage struct {
		Err    error
		Output *LoginIOutput
	}
)

// newLoginTask : 實例化Task
func newLoginTask() *loginTask {
	return &loginTask{
		APIName: "Login",
		Req:     &LoginInput{},
		Res: &protocol.Response{
			Code:    "1",
			Message: "",
			Result:  struct{}{},
		},
		Storage: &LoginStorage{
			Output: &LoginIOutput{},
		},
	}
}

// Login : 登入
func Login(c *fiber.Ctx) error {
	task := newLoginTask()

	// 解析參數
	if shouldBreak := task.BindRequest(c); shouldBreak {
		return middleware.ErrorHandler(task.Storage.Err)
	}

	ctx, cancel := context.WithTimeout(c.Context(), time.Minute*15)
	defer cancel()

	// 會員登入
	if shouldBreak := task.DoLogin(ctx); shouldBreak {
		return middleware.ErrorHandler(task.Storage.Err)
	}

	task.Res.Result = task.Storage.Output

	return c.JSON(task.Res)
}

// BindRequest : 解析參數
func (task *loginTask) BindRequest(c *fiber.Ctx) bool {
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

// DoLogin : 會員登入
func (task *loginTask) DoLogin(ctx context.Context) bool {
	token, err := member.Login(
		ctx,
		&member.LoginInput{
			Account:  task.Req.Account,
			Password: task.Req.Password,
		},
	)
	if err != nil {
		task.Storage.Err = err
		return true
	}

	task.Storage.Output.Token = token

	return false
}
