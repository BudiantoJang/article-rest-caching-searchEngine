package helper

import (
	"jang-article/internal/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Status       string      `json:"status"`
	ResponseCode string      `json:"responseCode"`
	Description  string      `json:"description"`
	Data         interface{} `json:"data"`
	List         interface{} `json:"list"`
}

func ResponseError(c echo.Context, status int, err error) error {
	res := Response{
		Status:       "ERROR",
		ResponseCode: model.ResponseCode99,
		Description:  err.Error(),
	}

	data, ok := err.(model.ErrorDetail)
	if ok {
		res.ResponseCode = data.Code
		res.Description = data.Description.Error()
	}

	return c.JSON(status, res)
}

func ResponseSuccess(c echo.Context, r interface{}) error {
	res := Response{
		Status:       "OK",
		ResponseCode: "00",
		Description:  "SUCCESS",
		Data:         r,
	}

	return c.JSON(http.StatusOK, res)
}
