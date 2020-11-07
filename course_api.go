package yunke

import (
	"fmt"

	class100 "github.com/class100/core"
	"github.com/class100/yunke-core"
)

// courser 课程接口
type courser interface {
	// CourseAdd 添加课程
	CourseAdd(req *core.AddCourseReq) (course *core.Course, err error)
	// CourseDelete 删除课程
	CourseDelete(id int64) (err error)
	// CourseUpdate 更新课程
	CourseUpdate(id int64, params map[string]interface{}) (course *core.Course, err error)
}

func (hsc *httpSignatureClient) CourseAdd(req *core.AddCourseReq) (course *core.Course, err error) {
	course = new(core.Course)
	err = hsc.requestApi(
		core.CourseApiAdd,
		class100.HttpMethodPost,
		req, nil,
		core.ApiVersionDefault,
		course,
	)

	return
}

func (hsc *httpSignatureClient) CourseDelete(id int64) (err error) {
	pathParams := map[string]string{
		"id": fmt.Sprintf("%v", id),
	}
	err = hsc.requestApi(
		core.CourseApiDelete,
		class100.HttpMethodDelete,
		nil, pathParams,
		core.ApiVersionDefault,
		nil,
	)

	return
}

func (hsc *httpSignatureClient) CourseUpdate(id int64, params map[string]interface{}) (course *core.Course, err error) {
	course = new(core.Course)
	pathParams := map[string]string{
		"id": fmt.Sprintf("%v", id),
	}
	err = hsc.requestApi(
		core.CourseApiUpdate,
		class100.HttpMethodPut,
		params, pathParams,
		core.ApiVersionDefault,
		course,
	)

	return
}
