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

var defaultNonceNumber = 18
var defaultNetworkType = "tcp"
var defaultCommonDeadline = 500 * time.Millisecond

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
		commonDeadline:    defaultCommonDeadline,
		networkType:       defaultNetworkType,
		credentials:       credentials,
		nonceGeneratorNum: defaultNonceNumber,
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
	wisdomWords, handshakeErr := dc.DoHandshake(conn)
	if handshakeErr != nil {
		log.Printf("ERROR: handshake failed - %s\n", handshakeErr)
		if connErr = conn.Close(); connErr != nil {
			log.Printf("ERROR: cannot close connection %s\n", connErr)
		}
		return nil, handshakeErr
	}
	return wisdomWords, nil
}

func (dc *DefaultClient) DoHandshake(conn net.Conn) ([]byte, error) {
	var readErr, writeErr error
	var readLen int
	var serverResponseArr []string
	var serverProofNonce []string

	clientNonce := randomStringGenerator(dc.nonceGeneratorNum)

	_, writeErr = conn.Write([]byte(fmt.Sprintf("%s,%s", dc.credentials.username, clientNonce)))
	if writeErr != nil {
		log.Println("ERROR: cannot write to connection - return")
		return nil, writeErr
	}
	buff := make([]byte, 1024)

	readLen, readErr = conn.Read(buff)

	if readErr != nil {
		log.Println("ERROR: cannot read from connection - return")
		return nil, readErr
	}
	serverResponseArr, buff = strings.Split(string(buff[:readLen]), ","), buff[readLen+1:]

	iterationCount, convErr := strconv.Atoi(serverResponseArr[2])

	if convErr != nil {
		log.Printf("ERROR: server responded with incorrect type of iteration count - return")
		return nil, convErr
	}
	serverNonce := serverResponseArr[0]
	salt := serverResponseArr[1]

	hmacGen := hmacGenerator(dc.credentials.password, salt, serverNonce, iterationCount)

	_, writeErr = conn.Write([]byte(fmt.Sprintf("%s,%s", serverNonce, hmacGen)))

	if writeErr != nil {
		log.Println("ERROR: cannot write to connection - return")
		return nil, writeErr
	}
	readLen, readErr = conn.Read(buff)

	if readErr != nil {
		log.Println("ERROR: cannot read from connection - return")
		return nil, readErr
	}
	clientHmacGen := hmacGenerator(dc.credentials.password, salt, clientNonce, iterationCount)

	serverProofNonce, buff = strings.Split(string(buff[:readLen]), ","), buff[readLen+1:]

	serverProof := serverProofNonce[1]
	serverSendClientNonce := serverProofNonce[0]

	if clientHmacGen != serverProof {
		log.Println("ERROR: proof is not correct - return")
		return nil, errors.New("server proof is not correct")
	}
	if clientNonce != serverSendClientNonce {
		log.Println("ERROR: nonce is not correct - return")
		return nil, errors.New("client nonce is not correct")
	}
	_, writeErr = conn.Write([]byte("success"))

	if writeErr != nil {
		log.Println("ERROR: cannot write to connection - return")
		return nil, writeErr
	}
	readLen, readErr = conn.Read(buff)

	if readErr != nil {
		log.Println("ERROR: cannot read from connection - return")
		return nil, readErr
	}
	return buff[:readLen], nil
}

func (dc *DefaultClient) SetNonceNumber(number int) {
	dc.nonceGeneratorNum = number
}

func (dc *DefaultClient) SetDeadlineToRead(deadline time.Duration) {
	dc.deadlineToRead = deadline
}

func (dc *DefaultClient) SetDeadlineToWrite(deadline time.Duration) {
	dc.deadlineToWrite = deadline
}

func (dc *DefaultClient) SetCommonDeadline(deadline time.Duration) {
	dc.commonDeadline = deadline
}

func (dc *DefaultClient) SetNetworkType(networkType string) {
	dc.networkType = networkType
}

func (dc *DefaultClient) GetNetworkType() string {
	return dc.networkType
}

func (dc *DefaultClient) GetNonceNumber() int {
	return dc.nonceGeneratorNum
}

func (dc *DefaultClient) GetDeadlineToRead() time.Duration {
	return dc.deadlineToRead
}

func (dc *DefaultClient) GetDeadlineToWrite() time.Duration {
	return dc.deadlineToWrite
}

func (dc *DefaultClient) GetCommonDeadline() time.Duration {
	return dc.commonDeadline
}
