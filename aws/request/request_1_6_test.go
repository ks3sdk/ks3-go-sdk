// +build go1.6

package request_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ks3sdk/ks3-go-sdk/aws"
	"github.com/ks3sdk/ks3-go-sdk/aws/awserr"
	"github.com/ks3sdk/ks3-go-sdk/aws/client"
	"github.com/ks3sdk/ks3-go-sdk/aws/client/metadata"
	"github.com/ks3sdk/ks3-go-sdk/aws/defaults"
	"github.com/ks3sdk/ks3-go-sdk/aws/request"
)

// go version 1.4 and 1.5 do not return an error. Version 1.5 will url encode
// the uri while 1.4 will not
func TestRequestInvalidEndpoint(t *testing.T) {
	endpoint := "http://localhost:90 "

	r := request.New(
		aws.Config{},
		metadata.ClientInfo{Endpoint: endpoint},
		defaults.Handlers(),
		client.DefaultRetryer{},
		&request.Operation{},
		nil,
		nil,
	)

	assert.Error(t, r.Error)
}

type timeoutErr struct {
	error
}

var errTimeout = awserr.New("foo", "bar", &timeoutErr{
	errors.New("net/http: request canceled"),
})

func (e *timeoutErr) Timeout() bool {
	return true
}

func (e *timeoutErr) Temporary() bool {
	return true
}
