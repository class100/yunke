package yunke

import (
	`strconv`

	class100 "github.com/class100/core"
	"github.com/class100/yunke-core"
	`github.com/storezhang/gox`
)

type user interface {
	// AddUser 添加用户
	AddUser(req *core.AddUserReq) (user *core.User, err error)
	// DeleteUser 删除用户
	DeleteUser(id int64) (err error)
	// UpdateUser 更新用户
	UpdateUser(id int64, params map[string]interface{}) (course *core.User, err error)
	// GetUser 获取用户
	GetUser(req *core.GetUserReq) (rsp *core.GetUserRsp, err error)
}

func (hsc *httpSignatureClient) AddUser(req *core.AddUserReq) (user *core.User, err error) {
	user = new(core.User)
	err = hsc.requestApi(
		core.UserApiAdd,
		class100.HttpMethodPost,
		req,
		core.ApiVersionDefault,
		user,
	)

	return
}

func (hsc *httpSignatureClient) DeleteUser(id int64) (err error) {
	err = hsc.requestApi(
		core.UserTeacherApiDelete,
		class100.HttpMethodDelete,
		nil,
		core.ApiVersionDefault,
		nil,
		gox.NewHttpPathParameter("id", strconv.FormatInt(id, 10)),
	)

	return
}

func (hsc *httpSignatureClient) UpdateUser(id int64, params map[string]interface{}) (user *core.User, err error) {
	user = new(core.User)
	err = hsc.requestApi(
		core.UserApiUpdate,
		class100.HttpMethodPut,
		params,
		core.ApiVersionDefault,
		user,
		gox.NewHttpPathParameter("id", strconv.FormatInt(id, 10)),
	)

	return
}

func (hsc *httpSignatureClient) GetUser(req *core.GetUserReq) (rsp *core.GetUserRsp, err error) {
	rsp = new(core.GetUserRsp)
	err = hsc.requestApi(
		core.UserApiGet,
		class100.HttpMethodGet,
		req,
		core.ApiVersionDefault,
		rsp,
		gox.NewHttpPathParameter("id", req.Id),
	)

	return
}
