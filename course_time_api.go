package yunke

import (
	class100 "github.com/class100/core"
	"github.com/class100/yunke-core"
)

// courseTimeHdl 课程时刻接口
type courseTimeHdl interface {
	// CourseTimeAdds 添加课程
	CourseTimeAdds(req *core.BatchAddCourseTimeReq, version core.ApiVersion) (err error)
	// CourseTimeDeletes 删除课程
	CourseTimeDeletes(req *core.BatchDeleteCourseTimeReq, version core.ApiVersion) (err error)
}

func (hsc *httpSignatureClient) CourseTimeAdds(req *core.BatchAddCourseTimeReq, version core.ApiVersion) (err error) {
	err = hsc.requestApi(core.CourseTimeApiAdd, class100.HttpMethodPost, nil, req, nil, version, nil)

	return
}

func (hsc *httpSignatureClient) CourseTimeDeletes(req *core.BatchDeleteCourseTimeReq, version core.ApiVersion) (err error) {
	err = hsc.requestApi(
		core.CourseTimeApiDelete,
		class100.HttpMethodDelete,
		nil,
		req,
		nil,
		version,
		nil,
	)

	return
}
