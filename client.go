package yunke

import (
	`fmt`

	`github.com/class100/core`
	`github.com/class100/yunke-core`
)

type httpSignatureClient struct {
	client *core.HttpSignatureClient

	options options
}

func (hsc *httpSignatureClient) requestApi(
	path yunke.ApiPath,
	method core.HttpMethod,
	headers map[string]string,
	params interface{}, pathParams map[string]string,
	version yunke.ApiVersion,
	rsp interface{},
) (err error) {
	var url string
	if yunke.ApiVersionDefault == version {
		url = fmt.Sprintf("%s/api/open/%s", hsc.options.Endpoint, path)
	} else {
		url = fmt.Sprintf("%s/api/open/%s/%s", hsc.options.Endpoint, version, path)
	}

	return hsc.client.RequestApi(
		url,
		method,
		headers,
		params, pathParams,
		rsp,
	)
}
