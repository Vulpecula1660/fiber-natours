package enum

import (
	"fmt"
	"net/http"
)

// CustomError :
type CustomError struct {
	HTTPStatus int
	Code       int
	Message    string
}

func (c CustomError) Error() string {
	return c.Message
}

// REST :
func REST(HTTPStatus int, Code int, Message string) *CustomError {
	return &CustomError{HTTPStatus: HTTPStatus, Code: Code, Message: Message}
}

var (
	ParameterError = func(err error) error {
		return REST(http.StatusBadRequest, 10000, fmt.Sprint("參數錯誤 ", err.Error()))
	}

	AccountUsed = REST(http.StatusBadRequest, 10001, "帳號已被使用")

	UserNotFound = REST(http.StatusBadRequest, 10002, "找不到會員")

	WrongPassword = REST(http.StatusBadRequest, 10003, "密碼錯誤")

	NonLogin = REST(http.StatusUnauthorized, 10004, "未登入")

	DataNotFound = REST(http.StatusNotFound, 10005, "找不到資料")
)
