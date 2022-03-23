package powalgo

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
	str := HMACGenerator(password, salt, nonce, repeatedNumber)

	// then
	assert.NotEmpty(t, str)
}

func TestPowAlgorithm_PasswordIsEmpty(t *testing.T) {
	// given
	salt := "salt"
	nonce := "nonce"
	repeatedNumber := 5

	// when
	str := HMACGenerator("", salt, nonce, repeatedNumber)

	// then
	assert.NotEmpty(t, str)
}

func TestPowAlgorithm_SaltIsEmpty(t *testing.T) {
	// given
	password := "abc"
	nonce := "nonce"
	repeatedNumber := 5

	// when
	str := HMACGenerator(password, "", nonce, repeatedNumber)

	// then
	assert.NotEmpty(t, str)
}

func TestPowAlgorithm_NonceIsEmpty(t *testing.T) {
	// given
	salt := "salt"
	password := "abc"
	repeatedNumber := 5

	// when
	str := HMACGenerator(password, salt, "", repeatedNumber)

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
	str := HMACGenerator(password, salt, nonce, repeatedNumber)

	// then
	assert.NotEmpty(t, str)
}
