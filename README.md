# Proof Of Work challenge-response protocol client

### Environment variables
CLIENT_ADDRESS - address of client server;   
CLIENT_USERNAME - username is using for accessing server and validating request;   
CLIENT_PASSWORD - password is using to accessing server and validating request;   

### Restrictions
Client default address - :8090, to change it - supply header above.   
Client should be provided by credentials using headers above, if not - client won't start.   
Nonce symbols amount should not be less than 18 symbols.   

### Interaction with server
````
curl -v -H 'X-Redirect-To:<SERVER_ADDRESS>' <CLIENT_ADDRESS>/ping
````

### Options
Client is configurable. Read timeout, write timeout, common timeout for read and ride could be supplied.   
Nonce generator number too, but it has a restriction not less than 18 symbols.   
Some examples below:
````go
package main

import (
	"log"
	"time"
)

type ClientOption func(c Client)

func WithNonceGenerator(nonceNumber int) ClientOption {
	return func(c Client) {
		if nonceNumber < 18 {
			log.Println("WARN: cannot set nonceNumber variable less than 18")
		} else {
			c.SetNonceNumber(nonceNumber)
		}
	}
}

func WithReadDeadline(deadline time.Duration) ClientOption {
	return func(c Client) {
		c.SetDeadlineToRead(deadline)
	}
}
````