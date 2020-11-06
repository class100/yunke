package yunke

import (
	"fmt"

	class100 "github.com/class100/core"
	"github.com/class100/yunke-core"
)

// courseHdl 课程接口
type courseHdl interface {
	// CourseAdd 添加课程
	CourseAdd(req *core.AddCourseReq, version core.ApiVersion) (course *core.Course, err error)
	// CourseDelete 删除课程
	CourseDelete(id int64, version core.ApiVersion) (err error)
	// CourseUpdate 更新课程
	CourseUpdate(id int64, params map[string]interface{}, version core.ApiVersion) (course *core.Course, err error)
}

func (hsc *httpSignatureClient) CourseAdd(req *core.AddCourseReq, version core.ApiVersion) (course *core.Course, err error) {
	course = new(core.Course)
	err = hsc.requestApi(core.CourseApiAdd, class100.HttpMethodPost, nil, req, nil, version, course)

	return
}

func (hsc *httpSignatureClient) CourseDelete(id int64, version core.ApiVersion) (err error) {
	pathParams := map[string]string{
		"id": fmt.Sprintf("%v", id),
	}
	err = hsc.requestApi(core.CourseApiDelete, class100.HttpMethodDelete, nil, nil, pathParams, version, nil)

	return
}

func (hsc *httpSignatureClient) CourseUpdate(
	id int64,
	params map[string]interface{},
	version core.ApiVersion,
) (course *core.Course, err error) {
	course = new(core.Course)
	pathParams := map[string]string{
		"id": fmt.Sprintf("%v", id),
	}
	err = hsc.requestApi(core.CourseApiUpdate, class100.HttpMethodPut, nil, params, pathParams, version, course)

	return
}
