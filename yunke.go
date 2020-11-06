package yunke

import (
	"strings"

	"github.com/class100/core"
)

// Client 云视课堂客户端
type Client interface {
	courseHdl
	courseTimeHdl
	fileHdl
	lectureHdl
	userHdl
}

// NewClient 创建云视课堂客户端
func NewClient(options ...Option) (client Client, err error) {
	appliedOptions := defaultOptions()
	for _, apply := range options {
		apply(&appliedOptions)
	}

	if "" == strings.TrimSpace(appliedOptions.Endpoint) {
		err = ErrMustSetEndpoint

		return
	}

	hsc := &httpSignatureClient{
		options: appliedOptions,
	}
	if hsc.client, err = core.NewHttpSignatureClient(appliedOptions.options...); nil != err {
		return
	}

	return
}
