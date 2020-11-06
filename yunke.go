package yunke

import (
	"strings"

	class100 "github.com/class100/core"
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

	client = &httpSignatureClient{
		client: class100.HttpSignatureClient{
			Options: appliedOptions.ClientOptions,
		},
		options: appliedOptions,
	}

	return
}
