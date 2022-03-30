package member

import (
	"context"
	"time"

	"github.com/Vulpecula1660/fiber-natours/api/middleware"
	"github.com/Vulpecula1660/fiber-natours/api/protocol"
	"github.com/Vulpecula1660/fiber-natours/enum"
	"github.com/Vulpecula1660/fiber-natours/model/dao/user"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type (
	InfoInput struct {
		ID string `json:"id" validate:"required"` // 會員ID
	}

	InfoIOutput struct {
		ID      string `json:"id"`
		Account string `json:"account"`
	}

	infoTask struct {
		APIName string
		Req     *InfoInput
		Res     *protocol.Response
		Storage *InfoStorage
	}

	// InfoStorage : 暫存
	InfoStorage struct {
		Err    error
		Output *InfoIOutput
	}
)

// newInfoTask : 實例化Task
func newInfoTask() *infoTask {
	return &infoTask{
		APIName: "Info",
		Req:     &InfoInput{},
		Res: &protocol.Response{
			Code:    "1",
			Message: "OK",
			Result:  struct{}{},
		},
		Storage: &InfoStorage{
			Output: &InfoIOutput{},
		},
	}
}

// Info : 登入
func Info(c *fiber.Ctx) error {
	task := newInfoTask()

	// 解析參數
	if shouldBreak := task.BindRequest(c); shouldBreak {
		return middleware.ErrorHandler(task.Storage.Err)
	}

	ctx, cancel := context.WithTimeout(c.Context(), time.Minute*15)
	defer cancel()

	// 取得資料
	if shouldBreak := task.GetInfo(ctx); shouldBreak {
		return middleware.ErrorHandler(task.Storage.Err)
	}

	task.Res.Result = task.Storage.Output

	return c.JSON(task.Res)
}

// BindRequest : 解析參數
func (task *infoTask) BindRequest(c *fiber.Ctx) bool {
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

// GetInfo : 取得資料
func (task *infoTask) GetInfo(ctx context.Context) bool {
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

	task.Storage.Output.ID = member[0].ID
	task.Storage.Output.Account = member[0].Account

	return false
}
