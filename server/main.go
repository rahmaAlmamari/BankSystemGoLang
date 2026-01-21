package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "BankSystem is running 🚀")
	})

	fmt.Println("Server running on :1323")
	e.Logger.Fatal(e.Start(":1323"))
}
