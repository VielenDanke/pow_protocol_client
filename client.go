package main

import (
	"net"
	"time"
)

type SetClientOptions interface {
	SetNonceNumber(number int)
	SetDeadlineToRead(deadline time.Duration)
	SetDeadlineToWrite(deadline time.Duration)
	SetCommonDeadline(deadline time.Duration)
	SetNetworkType(networkType string)
}

type Client interface {
	DoHandshake(conn net.Conn) error
	SendRequest(address string) ([]byte, error)
	SetClientOptions
}
