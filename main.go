package main

import (
    "net/http"
    "github.com/labstack/echo"
)

func main() {
    e := echo.New()

    e.GET("/hello", Hello())
    e.GET("/goodbye", Goodbye())
    e.GET("/api/hello", ApiHelloGet())
    e.GET("/api/goodbye", ApiGoodbyeGet())

    //e.Start(":8080")
    // エラーをチェックしてログに出力する(静的解析のログ対策）
    if err := e.Start(":8080"); err != nil {
        e.Logger.Fatal(err)
    }
}

func Hello() echo.HandlerFunc {
    return func(c echo.Context) error {     
        return c.String(http.StatusOK, "hello, world.")
    }
}

func Goodbye() echo.HandlerFunc {
    return func(c echo.Context) error {     
        return c.String(http.StatusOK, "goodbye.")
    }
}

func ApiHelloGet() echo.HandlerFunc {
    return func(c echo.Context) error {     
        return c.JSON(http.StatusOK, map[string]interface{}{"message": "ハロー"})
    }
}

func ApiGoodbyeGet() echo.HandlerFunc {
    return func(c echo.Context) error {     
        return c.JSON(http.StatusOK, map[string]interface{}{"message": "goodbye"})
    }
}
