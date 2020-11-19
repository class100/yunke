package yunke

import (
	`fmt`

	class100 `github.com/class100/core`
	`github.com/class100/yunke-core`
	`github.com/storezhang/gox`
)

// client 客户端接口
type client interface {
	ClientGet(clientType core.ClientType) (rsp *core.GetClientRsp, err error)
}

func (hsc *httpSignatureClient) ClientGet(clientType core.ClientType) (rsp *core.GetClientRsp, err error) {
	rsp = new(core.GetClientRsp)
	err = hsc.requestApi(
		core.ClientApiGetByType,
		class100.HttpMethodGet,
		nil,
		core.ApiVersionDefault,
		rsp,
		gox.NewHttpPathParameter("type", fmt.Sprintf("%v", clientType)),
	)

	return
}
