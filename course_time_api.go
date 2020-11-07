package yunke

import (
	class100 "github.com/class100/core"
	"github.com/class100/yunke-core"
)

// courseTimer 课程时刻接口
type courseTimer interface {
	// CourseTimeAdds 添加课程
	CourseTimeAdds(req *core.BatchAddCourseTimeReq) (err error)
	// CourseTimeDeletes 删除课程
	CourseTimeDeletes(req *core.BatchDeleteCourseTimeReq) (err error)
}

func (hsc *httpSignatureClient) CourseTimeAdds(req *core.BatchAddCourseTimeReq) (err error) {
	err = hsc.requestApi(
		core.CourseTimeApiAdd,
		class100.HttpMethodPost,
		req, nil,
		core.ApiVersionDefault,
		nil,
	)

	return
}

func (hsc *httpSignatureClient) CourseTimeDeletes(req *core.BatchDeleteCourseTimeReq) (err error) {
	err = hsc.requestApi(
		core.CourseTimeApiDelete,
		class100.HttpMethodDelete,
		req, nil,
		core.ApiVersionDefault,
		nil,
	)

	return
}
