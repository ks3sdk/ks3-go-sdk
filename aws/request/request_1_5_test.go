// +build !go1.6

package request_test

import (
	"errors"

	"github.com/ks3sdk/ks3-go-sdk/aws/awserr"
)

var errTimeout = awserr.New("foo", "bar", errors.New("net/http: request canceled Timeout"))
