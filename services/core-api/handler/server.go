package handler

import (
	"net/http"
	"os"
	"strings"

	"github.com/go-kit/kit/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


// NewServer creates a new server.
func NewServer(logger log.Logger) http.Handler {
	e := echo.New()
	s := &Server{logger: logger}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	apiGroup := e.Group("/api/v1")
	apiGroup.GET("/hello", s.Hello)

	// CORS middleware
	origins := os.Getenv("ORIGINS")
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     strings.Split(origins, ","),
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderWWWAuthenticate, "Authorization"},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
	}))

	return e
}
