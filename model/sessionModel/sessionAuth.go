package sessionmodel

import (
	"generate-auth-workshop/database"
	"generate-auth-workshop/model"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context, credentials model.Credentials, store *sessions.CookieStore) (*database.User, *model.ErrorResponse) {	
	// Retrieve the user from the database:
	user, err := model.RetrieveUser(credentials)
	if (err != nil) {
		return nil, err
	}
	if err = SessionCreate(c, store); err != nil {
		return nil, err
	}

	return user, nil
}

func Logout(c echo.Context, store *sessions.CookieStore) *model.ErrorResponse {
	if err := SessionAuthorize(c, store); err != nil {
		return err
	}
	if err := SessionRevoke(c, store); err != nil {
		return err
	}
	return nil
}

func SessionCreate(c echo.Context, store *sessions.CookieStore) *model.ErrorResponse {
	errResponse := &model.ErrorResponse{}
	
	// Create a new session:
	session, err := store.Get(c.Request(), "user_session")
	if err != nil {
		errResponse.Status = 500
		errResponse.Message = "Failed creating new session!"
		return errResponse
	}

	// Confirm that the session is new:
	if !session.IsNew {
		errResponse.Status = 404
		errResponse.Message = "You are already authenticated!"
		return errResponse
	}

	// Save the session:
	err = session.Save(c.Request(), c.Response())
	if err != nil {
		errResponse.Status = 500
		errResponse.Message = "Failed to save new session!"
	}

	return nil
}

func SessionAuthorize(c echo.Context, store *sessions.CookieStore) *model.ErrorResponse {
	errResponse := &model.ErrorResponse{}

	session, err := store.Get(c.Request(), "user_session")
	if err != nil {
		errResponse.Status = 500
		errResponse.Message = "Failed retrieving the session!"
		return errResponse
	}

	if session.IsNew {
		errResponse.Status = 401
		errResponse.Message = "Unauthorized Request!"
		return errResponse
	}

	if session.Options.MaxAge < 0 {
		errResponse.Status = 401
		errResponse.Message = "Unauthorized Request!"
		return errResponse
	}

	return nil
}

func SessionRevoke(c echo.Context, store *sessions.CookieStore) *model.ErrorResponse {
	errResponse := &model.ErrorResponse{}

	// Retrieve the session:
	session, err := store.Get(c.Request(), "user_session")
	if err != nil {
		errResponse.Status = 500
		errResponse.Message = "Failed retrieving the session!"
		return errResponse
	}

	// Invalidate the session cookie:
	session.Options.MaxAge = -1

	// Save the session:
	err = session.Save(c.Request(), c.Response())
	if err != nil {
		errResponse.Status = 500
		errResponse.Message = "Failed to revoke the session!"
	}

	return nil
}

func SessionGetSecretMessage(c echo.Context, store *sessions.CookieStore, id string) (*database.SecretNote, *model.ErrorResponse) {
	// Authenticate the session:
	if err := SessionAuthorize(c, store); err != nil {
		return nil, err
	}

	// Retrieve the secret note:
	secretNote := model.RetrieveSecretNotesByUserId(id);

	return &secretNote, nil
}
