package main

import "net"

type Client interface {
	doHandshake(conn net.Conn) error
	SendRequest(address string) ([]byte, error)
}
