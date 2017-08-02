package s3crypto_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ks3sdk/ks3-go-sdk/service/s3/s3crypto"
)

func TestCryptoReadCloserRead(t *testing.T) {
	expectedStr := "HELLO WORLD "
	str := strings.NewReader(expectedStr)
	rc := &s3crypto.CryptoReadCloser{Body: ioutil.NopCloser(str), Decrypter: str}

	b, err := ioutil.ReadAll(rc)
	assert.NoError(t, err)
	assert.Equal(t, expectedStr, string(b))
}

func TestCryptoReadCloserClose(t *testing.T) {
	data := "HELLO WORLD "
	expectedStr := ""

	str := strings.NewReader(data)
	rc := &s3crypto.CryptoReadCloser{Body: ioutil.NopCloser(str), Decrypter: str}
	rc.Close()

	b, err := ioutil.ReadAll(rc)
	assert.NoError(t, err)
	assert.Equal(t, expectedStr, string(b))
}
