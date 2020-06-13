package main

import (
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Static("/", "./dist")
	e.GET("/", home)
	e.Logger.Fatal(e.Start(":80"))
}

func home(c echo.Context) error {
	index, _ := ioutil.ReadFile("./dist/index.html")
	return c.Blob(http.StatusOK, "text/html", index)
}
