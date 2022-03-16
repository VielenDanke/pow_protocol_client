package main

import (
	"net"
	"time"
)

type ClientOptions interface {
	SetNonceNumber(number int)
	SetDeadlineToRead(deadline time.Duration)
	SetDeadlineToWrite(deadline time.Duration)
	SetCommonDeadline(deadline time.Duration)
	SetNetworkType(networkType string)
}

type Client interface {
	DoHandshake(conn net.Conn) ([]byte, error)
	SendRequest(address string) ([]byte, error)
	ClientOptions
}
