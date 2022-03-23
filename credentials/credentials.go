package credentials

import (
	"errors"
	"os"
)

var ErrCredentialsAreEmpty = errors.New("password or username cannot be empty")

type Credentials struct {
	username, password string
}

func BuildCredentials(username, password string) (Credentials, error) {
	if len(username) == 0 || len(password) == 0 {
		return Credentials{}, ErrCredentialsAreEmpty
	}
	return Credentials{username: username, password: password}, nil
}

func BuildCredentialsFromEnv() (Credentials, error) {
	username := os.Getenv("CLIENT_USERNAME")
	password := os.Getenv("CLIENT_PASSWORD")
	if len(username) == 0 || len(password) == 0 {
		return Credentials{}, ErrCredentialsAreEmpty
	}
	return Credentials{username: username, password: password}, nil
}

func (c Credentials) GetUsername() string {
	return c.username
}

func (c Credentials) GetPassword() string {
	return c.password
}
