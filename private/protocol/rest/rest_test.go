package rest_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/ks3sdk/ks3-go-sdk/aws"
	"github.com/ks3sdk/ks3-go-sdk/aws/client"
	"github.com/ks3sdk/ks3-go-sdk/aws/client/metadata"
	"github.com/ks3sdk/ks3-go-sdk/aws/request"
	"github.com/ks3sdk/ks3-go-sdk/aws/signer/v4"
	"github.com/ks3sdk/ks3-go-sdk/awstesting/unit"
	"github.com/ks3sdk/ks3-go-sdk/private/protocol/rest"
)

func TestUnsetHeaders(t *testing.T) {
	cfg := &aws.Config{Region: aws.String("us-west-2")}
	c := unit.Session.ClientConfig("testService", cfg)
	svc := client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName:   "testService",
			SigningName:   c.SigningName,
			SigningRegion: c.SigningRegion,
			Endpoint:      c.Endpoint,
			APIVersion:    "",
		},
		c.Handlers,
	)

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(rest.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(rest.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(rest.UnmarshalMetaHandler)
	op := &request.Operation{
		Name:     "test-operation",
		HTTPPath: "/",
	}

	input := &struct {
		Foo aws.JSONValue `location:"header" locationName:"x-amz-foo" type:"jsonvalue"`
		Bar aws.JSONValue `location:"header" locationName:"x-amz-bar" type:"jsonvalue"`
	}{}

	output := &struct {
		Foo aws.JSONValue `location:"header" locationName:"x-amz-foo" type:"jsonvalue"`
		Bar aws.JSONValue `location:"header" locationName:"x-amz-bar" type:"jsonvalue"`
	}{}

	req := svc.NewRequest(op, input, output)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewBuffer(nil)), Header: http.Header{}}
	req.HTTPResponse.Header.Set("X-Amz-Foo", "e30=")

	// unmarshal response
	rest.UnmarshalMeta(req)
	rest.Unmarshal(req)
	if req.Error != nil {
		t.Fatal(req.Error)
	}
}
