package sessioncontroller

import (
	"generate-auth-workshop/model"
	sessionmodel "generate-auth-workshop/model/sessionModel"
	"generate-auth-workshop/utils"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func RegisterSessionAuthRoutes(e *echo.Group, logger echo.Logger) {
	// Create a new endpoint group:
	api := e.Group("/session")

	// Create a session cookie store:
	var store = sessions.NewCookieStore([]byte(os.Getenv("SECRET_KEY")))
	store.Options = &sessions.Options{
		MaxAge:   int(60 * 60 * 24),
		Path:     "/",
		HttpOnly: true,
	}

	// Initialize backend middleware:
	api.Use(session.Middleware(store))

	// Establish user session:
	api.POST("/authorize", func(c echo.Context) error {
		// Unbind the credentials from the body of the request:
		var credentials model.Credentials
		if err := c.Bind(&credentials); err != nil {
			return c.JSON(utils.CreateErrorResponse(404, err.Error()))
		}

		// Login the user:
		user, loginError := sessionmodel.Login(c, credentials, store)
		if loginError != nil {
			return c.JSON(utils.CreateErrorResponse(loginError.Status, loginError.Message))
		} 

		return c.JSON(http.StatusOK, user)
	})

	// Revoke user session:
	api.POST("/revoke", func(c echo.Context) error {
		// Logout the user:
		if logoutErr := sessionmodel.Logout(c, store); logoutErr != nil {
			return c.JSON(utils.CreateErrorResponse(logoutErr.Status, logoutErr.Message))
		}

		return c.NoContent(http.StatusOK)
	})

	// Retrieve protected data:
	api.GET("/secret-notes/:id", func(c echo.Context) error {
		// Retrieve the query params:
		id := c.Param("id")

		// Retrieve the secret note:
		secretNotes, err := sessionmodel.SessionGetSecretMessage(c, store, id)
		if err != nil {
			return c.JSON(utils.CreateErrorResponse(err.Status, err.Message))
		}

		return c.JSON(http.StatusOK, secretNotes)
	})
}
