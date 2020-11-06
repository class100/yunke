package yunke

import (
	"fmt"

	class100 "github.com/class100/core"
	"github.com/class100/yunke-core"
)

type httpSignatureClient struct {
	client *class100.HttpSignatureClient

	options options
}

func (hsc *httpSignatureClient) requestApi(
	path core.ApiPath,
	method class100.HttpMethod,
	headers map[string]string,
	params interface{}, pathParams map[string]string,
	version core.ApiVersion,
	rsp interface{},
) (err error) {
	var url string
	if core.ApiVersionDefault == version {
		url = fmt.Sprintf("%s/api/open/%s", hsc.options.Endpoint, path)
	} else {
		url = fmt.Sprintf("%s/api/open/%s/%s", hsc.options.Endpoint, version, path)
	}

	if rsp == nil {
		rsp = new(interface{})
	}

	return hsc.client.RequestApi(
		url,
		method,
		headers,
		params, pathParams,
		rsp,
	)
}
