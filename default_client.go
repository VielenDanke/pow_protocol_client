package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

type ClientOption func(c Client)

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

type DefaultClient struct {
	deadlineToRead  time.Duration
	deadlineToWrite time.Duration
	commonDeadline  time.Duration
	networkType     string
}

func NewDefaultClient(opts ...ClientOption) Client {
	defaultClient := &DefaultClient{commonDeadline: 500 * time.Millisecond, networkType: "tcp"}

	for _, v := range opts {
		v(defaultClient)
	}
	return defaultClient
}

func (dc *DefaultClient) SendRequest(address string) ([]byte, error) {
	conn, connErr := net.Dial(dc.networkType, address)
	if connErr != nil {
		return nil, connErr
	}
	handshakeErr := dc.doHandshake(conn)
	if handshakeErr != nil {
		log.Printf("ERROR: handshake failed - %s\n", handshakeErr)
		if connErr = conn.Close(); connErr != nil {
			log.Printf("ERROR: cannot close connection %s\n", connErr)
		}
		return nil, handshakeErr
	}
	return []byte{}, nil
}

func (dc *DefaultClient) doHandshake(conn net.Conn) error {
	var readErr, writeErr error
	var readLen int
	var serverResponseArr []string
	var serverProof string
	// TODO: how to add user to request
	clientNonce := randomStringGenerator(18)
	_, writeErr = conn.Write([]byte(fmt.Sprintf("%s,%s", "user", clientNonce)))
	if writeErr != nil {
		log.Println("ERROR: cannot write to connection - return")
		return writeErr
	}
	buff := make([]byte, 1024)

	readLen, readErr = conn.Read(buff)

	if readErr != nil {
		log.Println("ERROR: cannot read from connection - return")
		return readErr
	}
	serverResponseArr, buff = strings.Split(string(buff[:readLen]), ","), buff[readLen+1:]

	iterationCount, convErr := strconv.Atoi(serverResponseArr[2])

	if convErr != nil {
		log.Printf("ERROR: server responded with incorrect type of iteration count - return")
		return convErr
	}
	serverNonce := serverResponseArr[0]
	salt := serverResponseArr[1]

	// TODO: how to add password to request
	hmacGen := hmacGenerator("password", salt, serverNonce, iterationCount)

	_, writeErr = conn.Write([]byte(fmt.Sprintf("%s,%s", serverNonce, hmacGen)))

	if writeErr != nil {
		log.Println("ERROR: cannot write to connection - return")
		return writeErr
	}
	readLen, readErr = conn.Read(buff)

	if readErr != nil {
		log.Println("ERROR: cannot read from connection - return")
		return readErr
	}
	clientHmacGen := hmacGenerator("password", salt, clientNonce, iterationCount)

	serverProof, buff = string(buff[:readLen]), buff[readLen+1:]

	if clientHmacGen != serverProof {
		log.Println("ERROR: proof is not correct - return")
		return errors.New("server proof is not correct")
	}
	_, writeErr = conn.Write([]byte("success"))

	if writeErr != nil {
		log.Println("ERROR: cannot write to connection - return")
		return writeErr
	}
	readLen, readErr = conn.Read(buff)

	if readErr != nil {
		log.Println("ERROR: cannot read from connection - return")
		return readErr
	}
	log.Printf("INFO: successful handshake, wisdom words - %s\n", buff[:readLen])
	return nil
}
