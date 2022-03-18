package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestBuildCredentials(t *testing.T) {
	// given
	password := "password"
	username := "username"

	// when
	creds, err := BuildCredentials(username, password)

	// then
	assert.Nil(t, err)
	assert.Equal(t, password, creds.password)
	assert.Equal(t, username, creds.username)
}

func TestBuildCredentialsFromEnv(t *testing.T) {
	// given
	username := "username"
	password := "password"
	os.Setenv("CLIENT_USERNAME", username)
	os.Setenv("CLIENT_PASSWORD", password)

	defer os.Clearenv()

	// when
	creds, err := BuildCredentialsFromEnv()

	// then
	assert.Nil(t, err)
	assert.Equal(t, username, creds.username)
	assert.Equal(t, password, creds.password)
}

func TestBuildCredentialsFromEnv_UsernameIsNotExists(t *testing.T) {
	// given
	os.Setenv("CLIENT_PASSWORD", "password")

	defer os.Clearenv()

	// when
	creds, err := BuildCredentialsFromEnv()

	// then
	assert.NotNil(t, err)
	assert.Empty(t, creds.password)
	assert.Empty(t, creds.username)
}

func TestBuildCredentialsFromEnv_PasswordIsNotExists(t *testing.T) {
	// given
	os.Setenv("CLIENT_USERNAME", "username")

	defer os.Clearenv()

	// when
	creds, err := BuildCredentialsFromEnv()

	// then
	assert.NotNil(t, err)
	assert.Empty(t, creds.username)
	assert.Empty(t, creds.password)
}

func TestBuildCredentials_NoUsername(t *testing.T) {
	// when
	creds, err := BuildCredentials("", "password")

	// then
	assert.NotNil(t, err)
	assert.Empty(t, creds.password)
	assert.Empty(t, creds.username)
}

func TestBuildCredentials_NoPassword(t *testing.T) {
	// when
	creds, err := BuildCredentials("username", "")

	// then
	assert.NotNil(t, err)
	assert.Empty(t, creds.password)
	assert.Empty(t, creds.username)
}
