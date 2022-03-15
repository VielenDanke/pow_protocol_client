package main

import "os"

type Credentials struct {
	username, password string
}

func BuildCredentials(username, password string) Credentials {
	return Credentials{username: username, password: password}
}

func BuildCredentialsFromEnv() Credentials {
	return Credentials{username: os.Getenv("CLIENT_USERNAME"), password: os.Getenv("CLIENT_PASSWORD")}
}
