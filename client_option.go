package main

import (
	"log"
	"time"
)

type ClientOption func(c Client)

func WithNonceGenerator(nonceNumber int) ClientOption {
	return func(c Client) {
		client := c.(*DefaultClient)
		if nonceNumber < 18 {
			log.Println("WARN: cannot set nonceNumber variable less than 18")
		} else {
			client.nonceGeneratorNum = nonceNumber
		}
	}
}

func WithReadDeadline(deadline time.Duration) ClientOption {
	return func(c Client) {
		client := c.(*DefaultClient)
		client.deadlineToRead = deadline
	}
}

func WithWriteDeadline(deadline time.Duration) ClientOption {
	return func(c Client) {
		client := c.(*DefaultClient)
		client.deadlineToWrite = deadline
	}
}

func WithCommonDeadline(commonDeadline time.Duration) ClientOption {
	return func(c Client) {
		client := c.(*DefaultClient)
		client.commonDeadline = commonDeadline
	}
}

func WithNetworkType(networkType string) ClientOption {
	return func(c Client) {
		client := c.(*DefaultClient)
		client.networkType = networkType
	}
}
