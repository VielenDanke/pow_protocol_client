package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var credentials Credentials

func init() {
	credentials, _ = BuildCredentials("username", "password")
}

func TestWithNetworkType(t *testing.T) {
	// given
	networkType := "udp"

	// when
	cli := NewDefaultClient(credentials, WithNetworkType(networkType))

	// then
	assert.Equal(t, networkType, cli.GetNetworkType())
}

func TestWithCommonDeadline(t *testing.T) {
	// given
	commonDeadline := 500 * time.Millisecond

	// when
	cli := NewDefaultClient(credentials, WithCommonDeadline(commonDeadline))

	// then
	assert.Equal(t, commonDeadline, cli.GetCommonDeadline())
}

func TestWithCommonDeadline_DefaultDeadline(t *testing.T) {
	// when
	cli := NewDefaultClient(credentials)

	// then
	assert.Equal(t, defaultCommonDeadline, cli.GetCommonDeadline())
}

func TestWithReadDeadline(t *testing.T) {
	// given
	readDeadline := 500 * time.Millisecond

	// when
	cli := NewDefaultClient(credentials, WithReadDeadline(readDeadline))

	// then
	assert.Equal(t, readDeadline, cli.GetDeadlineToRead())
}

func TestWithWriteDeadline(t *testing.T) {
	// given
	writeDeadline := 500 * time.Millisecond

	// when
	cli := NewDefaultClient(credentials, WithWriteDeadline(writeDeadline))

	// then
	assert.Equal(t, writeDeadline, cli.GetDeadlineToWrite())
}

func TestWithNonceGenerator(t *testing.T) {
	// given
	nonceNum := 100

	// when
	cli := NewDefaultClient(credentials, WithNonceGenerator(nonceNum))

	// then
	assert.Equal(t, nonceNum, cli.GetNonceNumber())
}

func TestWithNonceGenerator_IncorrectNonceNumber(t *testing.T) {
	// given
	nonceNum := 1

	// when
	cli := NewDefaultClient(credentials, WithNonceGenerator(nonceNum))

	// then
	assert.Equal(t, defaultNonceNumber, cli.GetNonceNumber())
}
