package awsutil_test

import (
	"testing"

	"github.com/ks3sdk/ks3-go-sdk/aws"
	"github.com/ks3sdk/ks3-go-sdk/aws/awsutil"
	"github.com/stretchr/testify/assert"
)

func TestDeepEqual(t *testing.T) {
	cases := []struct {
		a, b  interface{}
		equal bool
	}{
		{"a", "a", true},
		{"a", "b", false},
		{"a", aws.String(""), false},
		{"a", nil, false},
		{"a", aws.String("a"), true},
		{(*bool)(nil), (*bool)(nil), true},
		{(*bool)(nil), (*string)(nil), false},
		{nil, nil, true},
	}

	for i, c := range cases {
		assert.Equal(t, c.equal, awsutil.DeepEqual(c.a, c.b), "%d, a:%v b:%v, %t", i, c.a, c.b, c.equal)
	}
}
