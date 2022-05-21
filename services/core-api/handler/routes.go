package handler

import (
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/labstack/echo/v4"

	"fmt"
)

// Server struct
type Server struct {
	logger log.Logger
}

// Hello return hello string
func (s Server) Hello(c echo.Context) (err error) {
	return c.String(http.StatusOK, fmt.Sprintf("Hello %s", c.Param("name")))
}
