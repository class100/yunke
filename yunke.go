package yunke

import (
	"strings"

	"github.com/class100/core"
	"github.com/class100/yunke-core"
)

// Client 云视课堂客户端
type Client interface {
	// GetClient 按编号获得客户端
	GetClient(id int64, version yunke.ApiVersion) (client *yunke.BaseClient, err error)
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
