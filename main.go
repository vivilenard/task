package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/a-h/templ"
	"math/rand"
)


func main() {
	e := echo.New()
	e.Static("/static/", "static")
	
	e.GET("/random", randomStringHandler)
	e.GET("/", echo.WrapHandler(templ.Handler(mainView())))
	e.Logger.Fatal(e.Start(":3000"))
}

func randomStringHandler(c echo.Context) error{
	charSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	randomString := ""
	stringLength := 100
	
	for i := 0; i < stringLength; i++ {
		randomInt := rand.Intn(len(charSet) - 1)
		randomString += string(charSet[randomInt])
	}
	return c.String(http.StatusOK, randomString)
}
