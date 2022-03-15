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

type DefaultClient struct {
	deadlineToRead    time.Duration
	credentials       Credentials
	deadlineToWrite   time.Duration
	commonDeadline    time.Duration
	nonceGeneratorNum int
	networkType       string
}

func NewDefaultClient(credentials Credentials, opts ...ClientOption) Client {
	defaultClient := &DefaultClient{
		commonDeadline:    500 * time.Millisecond,
		networkType:       "tcp",
		credentials:       credentials,
		nonceGeneratorNum: 18,
	}

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

	clientNonce := randomStringGenerator(dc.nonceGeneratorNum)

	_, writeErr = conn.Write([]byte(fmt.Sprintf("%s,%s", dc.credentials.username, clientNonce)))
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

	hmacGen := hmacGenerator(dc.credentials.password, salt, serverNonce, iterationCount)

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
	clientHmacGen := hmacGenerator(dc.credentials.password, salt, clientNonce, iterationCount)

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
