package yunke

import (
	"strconv"

	"github.com/class100/core"
	"github.com/class100/yunke-core"
)

func (hsc *httpSignatureClient) GetClient(id int64, version yunke.ApiVersion) (client *yunke.BaseClient, err error) {
	client = new(yunke.BaseClient)
	err = hsc.requestApi(yunke.OrgApiClientGetById, core.HttpMethodGet, nil, nil, map[string]string{
		"id": strconv.FormatInt(id, 10),
	}, version, client)

	return
}
