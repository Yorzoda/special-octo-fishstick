package controller

import (
	"github.com/labstack/echo/v4"
)

type echoHandler struct {
	*echo.Echo
}

func NewHandler() *echoHandler {
	return &echoHandler{}
}
