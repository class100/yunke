package yunke

import (
	"strings"

	"github.com/class100/core"
)

// Client 云视课堂客户端
type Client interface {
	courser
	courseTimer
	filer
	lecturer
	user
	meeting
	fileChunk
}

// NewClient 创建云视课堂客户端
func NewClient(options ...Option) (client Client, err error) {
	var hsc *core.HttpSignatureClient

	appliedOptions := defaultOptions()
	for _, apply := range options {
		apply(&appliedOptions)
	}

	if "" == strings.TrimSpace(appliedOptions.Endpoint) {
		err = ErrMustSetEndpoint

		return
	}

	if hsc, err = core.NewHttpSignatureClient(appliedOptions.clientOptions...); nil != err {
		return
	}
	client = &httpSignatureClient{
		client:  hsc,
		options: appliedOptions,
	}

	return
}
