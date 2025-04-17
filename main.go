package main

import (
    "net/http"
    "github.com/labstack/echo"
)

func main() {
    e := echo.New()

    e.GET("/hello", Hello())
    e.GET("/goodbye", Goodbye())
    e.GET("/thankyou", Thankyou())
    e.GET("/api/hello", ApiHelloGet())
    e.GET("/api/goodbye", ApiGoodbyeGet())
    e.GET("/api/thankyou", ApiThankyouGet())

    //e.Start(":8080")
    // エラーをチェックしてログに出力する(静的解析のログ対策）
    if err := e.Start(":8080"); err != nil {
        e.Logger.Fatal(err)
    }
}

func Hello() echo.HandlerFunc {
    return func(c echo.Context) error {     
        //return c.String(http.StatusOK, "hello, world.")
        html := `<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta http-equiv="refresh" content="30">
  <title>hello, world!</title>
</head>
<body>
  <h1>hello, world!</h1>
</body>
</html>`
        return c.HTML(http.StatusOK, html)
    }
}

func Goodbye() echo.HandlerFunc {
    return func(c echo.Context) error {     
        return c.String(http.StatusOK, "goodbye.")
    }
}

func Thankyou() echo.HandlerFunc {
    return func(c echo.Context) error {     
        return c.String(http.StatusOK, "thank you.")
    }
}

func ApiHelloGet() echo.HandlerFunc {
    return func(c echo.Context) error {     
        return c.JSON(http.StatusOK, map[string]interface{}{"message": "hello, world."})
    }
}

func ApiGoodbyeGet() echo.HandlerFunc {
    return func(c echo.Context) error {     
        return c.JSON(http.StatusOK, map[string]interface{}{"message": "goodbye."})
    }
}

func ApiThankyouGet() echo.HandlerFunc {
    return func(c echo.Context) error {     
        return c.JSON(http.StatusOK, map[string]interface{}{"message": "Thank you."})
    }
}
