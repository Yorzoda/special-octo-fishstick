package controller

import (
	"github.com/labstack/echo/v4"
)

type EchoHandler struct {
	*echo.Echo
}

func NewHandler() *EchoHandler {
	return &EchoHandler{}
}
