package yunke

import (
	`strconv`

	class100 "github.com/class100/core"
	"github.com/class100/yunke-core"
	`github.com/storezhang/gox`
)

// courser 课程接口
type courser interface {
	// AddCourse 添加课程
	AddCourse(req *core.AddCourseReq) (course *core.Course, err error)
	// DeleteCourse 删除课程
	DeleteCourse(id int64) (err error)
	// UpdateCourse 更新课程
	UpdateCourse(id int64, params map[string]interface{}) (course *core.Course, err error)
}

func (hsc *httpSignatureClient) AddCourse(req *core.AddCourseReq) (course *core.Course, err error) {
	course = new(core.Course)
	err = hsc.requestApi(
		core.CourseApiAdd,
		class100.HttpMethodPost,
		req,
		core.ApiVersionDefault,
		course,
	)

	return
}

func (hsc *httpSignatureClient) DeleteCourse(id int64) (err error) {
	err = hsc.requestApi(
		core.CourseApiDelete,
		class100.HttpMethodDelete,
		nil,
		core.ApiVersionDefault,
		nil,
		gox.NewHttpPathParameter("id", strconv.FormatInt(id, 10)),
	)

	return
}

func (hsc *httpSignatureClient) UpdateCourse(id int64, params map[string]interface{}) (course *core.Course, err error) {
	course = new(core.Course)
	err = hsc.requestApi(
		core.CourseApiUpdate,
		class100.HttpMethodPut,
		params,
		core.ApiVersionDefault,
		course,
		gox.NewHttpPathParameter("id", strconv.FormatInt(id, 10)),
	)

	return
}
