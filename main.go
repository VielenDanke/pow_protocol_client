package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	cliAddress := os.Getenv("CLIENT_ADDRESS")

	if len(cliAddress) == 0 {
		cliAddress = ":8090"
	}
	credentials, credentialsErr := BuildCredentialsFromEnv()
	if credentialsErr != nil {
		log.Fatalln("ERROR: credentials for server cannot be empty")
	}
	cli := NewDefaultClient(credentials)

	http.HandleFunc("/ping", func(rw http.ResponseWriter, r *http.Request) {
		addressTo := r.Header.Get("X-Redirect-To")

		if len(addressTo) == 0 {
			addressTo = "localhost:8080"
		}
		response, err := cli.SendRequest(addressTo)

		if err != nil {
			log.Printf("ERROR: call remote server on address %s is failed by %s error\n", addressTo, err)
			return
		}
		log.Printf("INFO: wisdom words - %s\n", response)
	})
	log.Fatalf("Service terminated: %s\n", http.ListenAndServe(cliAddress, nil))
}
