// +build codegen

package api

import (
	"bytes"
	"fmt"
)

type wafregionalExamplesBuilder struct {
	defaultExamplesBuilder
}

func (builder wafregionalExamplesBuilder) Imports(a *API) string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(`"fmt"
	"strings"
	"time"

	"github.com/ks3sdk/ks3-go-sdk/aws"
	"github.com/ks3sdk/ks3-go-sdk/aws/awserr"
	"github.com/ks3sdk/ks3-go-sdk/aws/session"
	"github.com/ks3sdk/ks3-go-sdk/service/waf"
	`)

	buf.WriteString(fmt.Sprintf("\"%s/%s\"", "github.com/ks3sdk/ks3-go-sdk/service", a.PackageName()))
	return buf.String()
}
