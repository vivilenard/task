package main

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/a-h/templ"
	"io"
)


func main() {
	e := echo.New()
	e.Static("/static/", "static")
	
	e.GET("/random", randomStringHandler)
	e.GET("/", echo.WrapHandler(templ.Handler(mainView())))
	e.Logger.Fatal(e.Start(":3000"))
}

func randomStringHandler(c echo.Context) error {
	url := "https://osob.de/r"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making GET request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: Received status code %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v", err)
	}
	return c.String(http.StatusOK, string(body))
}
