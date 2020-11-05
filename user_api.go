package yunke

import (
	"fmt"

	class100 "github.com/class100/core"
	"github.com/class100/yunke-core"
)

type userHdl interface {
	// UserAdd 添加用户
	UserAdd(req *core.AddUserReq, version core.ApiVersion) (user *core.User, err error)
	UserDelete(id int64, version core.ApiVersion) (err error)
	UserUpdate(
		id int64,
		params map[string]interface{},
		version core.ApiVersion,
	) (course *core.User, err error)
}

func (hsc httpSignatureClient) UserAdd(req *core.AddUserReq, version core.ApiVersion) (user *core.User, err error) {
	user = new(core.User)
	err = hsc.requestApi(core.UserApiAdd, class100.HttpMethodPost, nil, req, nil, version, user)

	return
}

func (hsc httpSignatureClient) UserDelete(id int64, version core.ApiVersion) (err error) {
	pathParams := map[string]string{
		"id": fmt.Sprintf("%v", id),
	}
	err = hsc.requestApi(core.UserTeacherApiDelete, class100.HttpMethodDelete, nil, nil, pathParams, version, nil)

	return
}

func (hsc httpSignatureClient) UserUpdate(
	id int64,
	params map[string]interface{},
	version core.ApiVersion,
) (user *core.User, err error) {
	user = new(core.User)

	pathParams := map[string]string{
		"id": fmt.Sprintf("%v", id),
	}
	err = hsc.requestApi(core.CourseApiUpdate, class100.HttpMethodPut, nil, params, pathParams, version, user)

	return
}
