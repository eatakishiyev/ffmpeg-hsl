package main

import (
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	service "hsl_proxy/services"
)

var (
	baseUrl = "/api/v1"
)

func main() {

	server := echo.New()
	baseGroup := server.Group(baseUrl)
	baseGroup.Use(middleware.Logger())
	hls := baseGroup.Group("/hls")
	hls.POST("", service.Hls)

	server.Start(":8080")
}
