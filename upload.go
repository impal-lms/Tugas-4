package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	baseUrl := os.Getenv("BASE_URL")
	if baseUrl == "" {
		baseUrl = "http://localhost"
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	file := e.Group("/file")
	baseUrl += "/file"
	file.POST("/upload", func(c echo.Context) error {
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}

		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		path := "assets/"
		mode := os.ModePerm
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, mode)
		}

		extension := filepath.Ext(file.Filename)
		filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), extension)
		dst, err := os.Create(path + filename)
		if err != nil {
			return err
		}

		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]string{
			"path": baseUrl + "/" + filename,
		})
	})

	file.Static("/", "assets")

	e.Logger.Fatal(e.Start(":4000"))
}
