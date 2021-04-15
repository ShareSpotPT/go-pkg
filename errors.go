package pkg

import (
	"fmt"
	"github.com/ShareSpotPT/go-pkg/code"
)

type Error struct {
	Code code.Code `json:"code"`
	Message string `json:"message"`
}

func New(code code.Code, msg string) Error {
	return Error{
		Code: code,
		Message: msg,
	}
}

func (e Error) Error() string {
	return fmt.Sprintf("%d | %s", e.Code, e.Message)
}