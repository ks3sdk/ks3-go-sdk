package s3manager

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ks3sdk/ks3-go-sdk/aws"
	"github.com/ks3sdk/ks3-go-sdk/aws/awserr"
	"github.com/ks3sdk/ks3-go-sdk/awstesting/unit"
	"github.com/ks3sdk/ks3-go-sdk/service/s3"
)

func testSetupGetBucketRegionServer(region string, statusCode int, incHeader bool) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if incHeader {
			w.Header().Set(bucketRegionHeader, region)
		}
		w.WriteHeader(statusCode)
	}))
}

var testGetBucketRegionCases = []struct {
	RespRegion string
	StatusCode int
}{
	{"bucket-region", 301},
	{"bucket-region", 403},
	{"bucket-region", 200},
}

func TestGetBucketRegion_Exists(t *testing.T) {
	for i, c := range testGetBucketRegionCases {
		server := testSetupGetBucketRegionServer(c.RespRegion, c.StatusCode, true)

		sess := unit.Session.Copy()
		sess.Config.Endpoint = aws.String(server.URL)
		sess.Config.DisableSSL = aws.Bool(true)

		ctx := aws.BackgroundContext()
		region, err := GetBucketRegion(ctx, sess, "bucket", "region")
		if err != nil {
			t.Fatalf("%d, expect no error, got %v", i, err)
		}
		if e, a := c.RespRegion, region; e != a {
			t.Errorf("%d, expect %q region, got %q", i, e, a)
		}
	}
}

func TestGetBucketRegion_NotExists(t *testing.T) {
	server := testSetupGetBucketRegionServer("ignore-region", 404, false)

	sess := unit.Session.Copy()
	sess.Config.Endpoint = aws.String(server.URL)
	sess.Config.DisableSSL = aws.Bool(true)

	ctx := aws.BackgroundContext()
	region, err := GetBucketRegion(ctx, sess, "bucket", "region")
	if err == nil {
		t.Fatalf("expect error, but did not get one")
	}
	aerr := err.(awserr.Error)
	if e, a := "NotFound", aerr.Code(); e != a {
		t.Errorf("expect %s error code, got %s", e, a)
	}
	if len(region) != 0 {
		t.Errorf("expect region not to be set, got %q", region)
	}
}

func TestGetBucketRegionWithClient(t *testing.T) {
	for i, c := range testGetBucketRegionCases {
		server := testSetupGetBucketRegionServer(c.RespRegion, c.StatusCode, true)

		svc := s3.New(unit.Session, &aws.Config{
			Region:     aws.String("region"),
			Endpoint:   aws.String(server.URL),
			DisableSSL: aws.Bool(true),
		})

		ctx := aws.BackgroundContext()

		region, err := GetBucketRegionWithClient(ctx, svc, "bucket")
		if err != nil {
			t.Fatalf("%d, expect no error, got %v", i, err)
		}
		if e, a := c.RespRegion, region; e != a {
			t.Errorf("%d, expect %q region, got %q", i, e, a)
		}
	}
}
