package model

import "fmt"

var (
	ResponseCode99 = "99"
)

type ErrorDetail struct {
	Code        string `json:"error_code"`
	Description error  `json:"error_description"`
}

func (e ErrorDetail) Error() string {
	return fmt.Sprintf("error_code: %s, error_description: %v", e.Code, e.Description)
}
