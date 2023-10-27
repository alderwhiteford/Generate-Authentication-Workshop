package model

import (
	"generate-auth-workshop/database"
)

func RetrieveUser(credentials Credentials) (*database.User, *ErrorResponse) {
	errResponse := &ErrorResponse{}

	// Check if user exists:
	var User *database.User
	for i := 0 ; i < len(database.Users) ; i++ {
		if database.Users[i].UserName == credentials.UserName {
			User = &database.Users[i]
			break
		}
	}
	if User == nil {
		errResponse.Status = 404
		errResponse.Message = "User does not exist!"
		return nil, errResponse
	}

	// Check if passwords match:
	if User.Password != credentials.Password {
		errResponse.Status = 404
		errResponse.Message = "Incorrect password!"
		return nil, errResponse
	}

	// Return user:
	return User, nil
}

func RetrieveSecretNotesByUserId(id string) database.SecretNote {
	secretNote := database.SecretNote{}

	for i := 0 ; i < len(database.SecretNotes) ; i++ {
		if database.SecretNotes[i].Author == id {
			secretNote = database.SecretNotes[i]
		}
	}

	return secretNote
}
