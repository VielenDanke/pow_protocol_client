package main

import "log"

func main() {
	defaultClient := NewDefaultClient(BuildCredentials("user", "password"))
	for {
		_, err := defaultClient.SendRequest("localhost:8080")

		if err != nil {
			log.Fatalln(err)
		}
	}
}
