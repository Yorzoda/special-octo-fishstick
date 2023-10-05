package main

import (
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const (
	_addr = ":1234"
)

func main() {
	srv := echo.New()

	srv.GET("/", singleParam)
	srv.GET("/users/:name", singleParam2)
	srv.GET("/multi", multiParams)
	srv.POST("/userinfo", bindUserInfo)
	srv.POST("/upload", uploadFile)

	srv.Logger.Fatal(srv.Start(_addr))

}

func singleParam(c echo.Context) error {
	name := c.QueryParam("val1")
	return c.String(http.StatusOK, name)
}

func singleParam2(c echo.Context) error {
	param := c.Param("name")
	return c.String(http.StatusOK, "value ="+param)
}

func multiParams(c echo.Context) error {
	params := c.QueryParams()
	v1 := params.Get("val1")
	v2 := params.Get("val2")
	return c.String(http.StatusOK, v1+v2)
}

type user struct {
	Name     string `json:"name"`
	Alifovec bool   `json:"alifovec"`
}

func bindUserInfo(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}

	if !u.Alifovec {
		return c.String(200, "why are you not alivec?"+u.Name+"?")
	}
	return c.String(http.StatusOK, "Hi "+u.Name+" kucho obed ravem?")
}

func uploadFile(c echo.Context) error {

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(filepath.Join("./echo/files/", filepath.Base(file.Filename)))
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.String(http.StatusOK, file.Filename+" file uploaded successfully")

}
