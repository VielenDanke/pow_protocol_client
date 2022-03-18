package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPowAlgorithm(t *testing.T) {
	// given
	password := "abc"
	salt := "salt"
	nonce := "nonce"
	repeatedNumber := 5

	// when
	str := hmacGenerator(password, salt, nonce, repeatedNumber)

	// then
	assert.NotEmpty(t, str)
}

func TestPowAlgorithm_PasswordIsEmpty(t *testing.T) {
	// given
	salt := "salt"
	nonce := "nonce"
	repeatedNumber := 5

	// when
	str := hmacGenerator("", salt, nonce, repeatedNumber)

	// then
	assert.NotEmpty(t, str)
}

func TestPowAlgorithm_SaltIsEmpty(t *testing.T) {
	// given
	password := "abc"
	nonce := "nonce"
	repeatedNumber := 5

	// when
	str := hmacGenerator(password, "", nonce, repeatedNumber)

	// then
	assert.NotEmpty(t, str)
}

func TestPowAlgorithm_NonceIsEmpty(t *testing.T) {
	// given
	salt := "salt"
	password := "abc"
	repeatedNumber := 5

	// when
	str := hmacGenerator(password, salt, "", repeatedNumber)

	// then
	assert.NotEmpty(t, str)
}

func TestPowAlgorithm_RepeatedNumberIsLessThanZero(t *testing.T) {
	// given
	salt := "salt"
	password := "abc"
	repeatedNumber := -1
	nonce := "nonce"

	// when
	str := hmacGenerator(password, salt, nonce, repeatedNumber)

	// then
	assert.NotEmpty(t, str)
}
