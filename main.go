package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", index_get)
	e.POST("/echo/:id", echo_post)
	e.GET("/echo/:id", echo_get)
	e.Logger.Fatal(e.Start(":1323"))

}

func index_get(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
func echo_post(c echo.Context) error {
	return c.String(http.StatusOK, c.Param("id"))
}
func echo_get(c echo.Context) error {
	return c.String(http.StatusOK, c.Param("id"))
}

func add(a, b int) int {
	return a + b
}
