package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacobintern/meme_coin/pkg/errors"
)

type Res struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func OK(ctx *gin.Context, resp any) {
	ctx.JSON(http.StatusOK, Res{
		Code: 0,
		Msg:  "success",
		Data: resp,
	})
}

func Failed(ctx *gin.Context, err error) {
	customErr := errors.CauseCustomError(err)
	httpCode := customErr.Code()
	httpStatus := customErr.Status().ToHTTPStatus()
	ctx.JSON(httpStatus, gin.H{
		"code": httpCode,
		"msg":  err.Error(),
	})
}
