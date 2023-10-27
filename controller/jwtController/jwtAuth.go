package jwtcontroller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterJWTAuthRoutes(e *echo.Group, logger echo.Logger) {
	api := e.Group("/jwt")

	// Establish user session:
	api.POST("/authorize", func(c echo.Context) error {
		logger.Info("JWT authorize endpoint hit!")

		return c.NoContent(http.StatusOK)
	})

	// Revoke user session:
	api.POST("/revoke", func(c echo.Context) error {
		logger.Info("JWT revoke endpoint hit!")

		return c.NoContent(http.StatusOK)
	})

	// Retrieve protected data:
	api.GET("/protected", func(c echo.Context) error {
		logger.Info("JWT protected endpoint hit!")

		return c.NoContent(http.StatusOK)
	})
}