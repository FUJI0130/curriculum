package response

import (
	"net/http"

	"github.com/cockroachdb/errors"
)

type ErrorResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func NewErrorResponse(err error) *ErrorResponse {
	rootErr := errors.Cause(err) // エラーの原因を取得
	return &ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: rootErr.Error(),
		Details: err.Error(), // ラップされたエラーメッセージも提供する
	}
}
