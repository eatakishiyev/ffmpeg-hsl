package main

import (
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	service "hsl_proxy/services"
	"net/http"
)

var (
	baseUrl = "/api/v1"
)

func main() {

	server := echo.New()
	server.Use(middleware.Logger())

	baseGroup := server.Group(baseUrl)
	baseGroup.Use(middleware.CORS())
	baseGroup.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	hls := baseGroup.Group("/hls")

	hls.POST("", service.Hls)
	server.Start(":8081")
}
