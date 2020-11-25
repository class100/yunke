package yunke

import (
	class100 "github.com/class100/core"
	"github.com/class100/yunke-core"
)

// courseTimer 课程时刻接口
type courseTimer interface {
	// AddCourseTimes 添加课程
	AddCourseTimes(req *core.BatchAddCourseTimeReq) (err error)
	// CourseTimeDeletes 删除课程
	DeleteCourseTimes(req *core.BatchDeleteCourseTimeReq) (err error)
}

func (hsc *httpSignatureClient) AddCourseTimes(req *core.BatchAddCourseTimeReq) (err error) {
	err = hsc.requestApi(
		core.CourseTimeApiAdd,
		class100.HttpMethodPost,
		req,
		core.ApiVersionDefault,
		nil,
	)

	return
}

func (hsc *httpSignatureClient) DeleteCourseTimes(req *core.BatchDeleteCourseTimeReq) (err error) {
	err = hsc.requestApi(
		core.CourseTimeApiDelete,
		class100.HttpMethodDelete,
		req,
		core.ApiVersionDefault,
		nil,
	)

	return
}
