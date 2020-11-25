package yunke

import (
	`strconv`

	class100 "github.com/class100/core"
	"github.com/class100/yunke-core"
	`github.com/storezhang/gox`
)

// courseContent 课程内容接口
type courseContent interface {
	// AddCourseContents 添加课程内容
	AddCourseContents(req *core.AddCourseContentReq) (rsp *core.AddCourseContentRsp, err error)
	// GetsCourseContentByCourseId 根据课程编号获取课程内容
	GetsCourseContentByCourseId(req *core.GetCourseContentReq) (rsp *core.GetCourseContentRsp, err error)
	// 根据课程编号或者课程内容编号删除
	DeleteCourseContents(req *core.DeleteCourseContentReq) (err error)
}

func (hsc *httpSignatureClient) AddCourseContents(
	req *core.AddCourseContentReq,
) (rsp *core.AddCourseContentRsp, err error) {
	rsp = new(core.AddCourseContentRsp)
	err = hsc.requestApi(
		core.CourseContentApiAdd,
		class100.HttpMethodPost,
		req,
		core.ApiVersionDefault,
		rsp,
	)

	return
}

func (hsc *httpSignatureClient) GetsCourseContentByCourseId(
	req *core.GetCourseContentReq,
) (rsp *core.GetCourseContentRsp, err error) {
	rsp = new(core.GetCourseContentRsp)
	err = hsc.requestApi(
		core.CourseContentApiGet,
		class100.HttpMethodGet,
		nil,
		core.ApiVersionDefault,
		rsp,
		gox.NewHttpPathParameter("courseId", strconv.FormatInt(req.CourseId, 10)),
	)

	return
}

func (hsc *httpSignatureClient) DeleteCourseContents(req *core.DeleteCourseContentReq) (err error) {
	err = hsc.requestApi(
		core.CourseContentApiDelete,
		class100.HttpMethodDelete,
		req,
		core.ApiVersionDefault,
		nil,
	)

	return
}
