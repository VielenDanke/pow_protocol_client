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

func WithWriteDeadline(deadline time.Duration) ClientOption {
	return func(c Client) {
		c.SetDeadlineToWrite(deadline)
	}
}

func WithCommonDeadline(commonDeadline time.Duration) ClientOption {
	return func(c Client) {
		c.SetCommonDeadline(commonDeadline)
	}
}

func WithNetworkType(networkType string) ClientOption {
	return func(c Client) {
		c.SetNetworkType(networkType)
	}
}
