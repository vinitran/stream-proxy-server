package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

func responseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{Code: http.StatusOK, Data: data})
}

func responseFailureWithMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, Response{Code: http.StatusBadRequest, Data: true, Message: message})
}

func responseSuccessWithMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, Response{Code: http.StatusOK, Data: true, Message: message})
}

func responseErrUnauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{Code: http.StatusUnauthorized, Message: http.StatusText(http.StatusUnauthorized)})
}

func responseErrValidation(c *gin.Context, message string) {
	c.JSON(http.StatusUnprocessableEntity, Response{Code: http.StatusUnprocessableEntity, Message: message})
}

func responseErrInternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, Response{Code: http.StatusInternalServerError, Message: http.StatusText(http.StatusInternalServerError)})
}

func responseErrInternalServerErrorWithDetail(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, Response{Code: http.StatusInternalServerError, Message: message})
}

func responseSuccessWithData(c *gin.Context, contentType string, data []byte) {
	c.Data(http.StatusOK, contentType, data)
}

func responseErrNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, Response{Code: http.StatusNotFound, Message: http.StatusText(http.StatusNotFound)})
}
