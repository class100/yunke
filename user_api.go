package yunke

import (
	`strconv`

	class100 "github.com/class100/core"
	"github.com/class100/yunke-core"
)

type user interface {
	// UserAdd 添加用户
	UserAdd(req *core.AddUserReq) (user *core.User, err error)
	UserDelete(id int64) (err error)
	UserUpdate(id int64, params map[string]interface{}) (course *core.User, err error)
}

func (hsc *httpSignatureClient) UserAdd(req *core.AddUserReq) (user *core.User, err error) {
	user = new(core.User)
	err = hsc.requestApi(
		core.UserApiAdd,
		class100.HttpMethodPost,
		req, nil,
		core.ApiVersionDefault,
		user,
	)

	return
}

func (hsc *httpSignatureClient) UserDelete(id int64) (err error) {
	err = hsc.requestApi(
		core.UserTeacherApiDelete,
		class100.HttpMethodDelete,
		nil, map[string]string{
			"id": strconv.FormatInt(id, 10),
		},
		core.ApiVersionDefault,
		nil,
	)

	return
}

func (hsc *httpSignatureClient) UserUpdate(id int64, params map[string]interface{}) (user *core.User, err error) {
	user = new(core.User)
	err = hsc.requestApi(
		core.CourseApiUpdate,
		class100.HttpMethodPut,
		params, map[string]string{
			"id": strconv.FormatInt(id, 10),
		},
		core.ApiVersionDefault,
		user,
	)

	return
}
