package main

import (
	"fmt"
	"io/ioutil"
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
	e.GET("/weather", weather_get)
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

func weather_get(c echo.Context) error {
	url := "https://www.jma.go.jp/bosai/common/const/area.json"

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(byteArray)) // htmlをstringで取得
	return c.String(http.StatusOK, string(byteArray))
}

// func weather_area_get(c echo.Context) error {
// 	url := "https://www.jma.go.jp/bosai/common/const/area.json"

// 	req, _ := http.NewRequest("GET", url, nil)

// 	client := new(http.Client)
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	byteArray, _ := ioutil.ReadAll(resp.Body)
// 	fmt.Println(string(byteArray))
// 	return c.String(http.StatusOK, string(byteArray))
// }

func add(a, b int) int {
	return a + b
}
