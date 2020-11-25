package yunke

import (
	`strconv`

	class100 "github.com/class100/core"
	"github.com/class100/yunke-core"
	`github.com/storezhang/gox`
)

type lecturer interface {
	// AddLecture 章节添加
	AddLecture(req *core.AddLectureReq) (lecture *core.LectureInfo, err error)
	// DeleteLecture 章节删除
	DeleteLecture(id int64) (err error)
	// LectureUpdate 章节更新
	UpdateLecture(id int64, req map[string]interface{}) (lecture *core.LectureInfo, err error)
	// GetLectureById 根据Id获取
	GetLectureById(id int64) (lecture *core.LectureInfo, err error)
	// LectureGetsByCourseId 根据课程Id获取
	GetsLectureByCourseId(courseId int64) (chapters []*core.ChapterInfo, err error)
	// LectureSwitchSequence 调整章节顺序
	SwitchLectureSequence(req *core.SwitchSequenceReq) (lectures []*core.LectureInfo, err error)
	// LectureFirstByCourseId 获取课程第一个章节讲次
	FirstLectureByCourseId(courseId int64) (lecture *core.LectureInfo, err error)
}

func (hsc *httpSignatureClient) AddLecture(req *core.AddLectureReq) (lecture *core.LectureInfo, err error) {
	lecture = new(core.LectureInfo)
	err = hsc.requestApi(
		core.LectureApiAdd,
		class100.HttpMethodPost,
		req,
		core.ApiVersionDefault,
		lecture,
	)

	return
}

func (hsc *httpSignatureClient) DeleteLecture(id int64) (err error) {
	err = hsc.requestApi(
		core.LectureApiDeleteById,
		class100.HttpMethodDelete,
		nil,
		core.ApiVersionDefault,
		nil,
		gox.NewHttpPathParameter("id", strconv.FormatInt(id, 10)),
	)

	return
}

func (hsc *httpSignatureClient) UpdateLecture(id int64, req map[string]interface{}) (lecture *core.LectureInfo, err error) {
	lecture = new(core.LectureInfo)

	err = hsc.requestApi(
		core.LectureApiUpdate,
		class100.HttpMethodPut,
		req,
		core.ApiVersionDefault,
		lecture,
		gox.NewHttpPathParameter("id", strconv.FormatInt(id, 10)),
	)

	return
}

func (hsc *httpSignatureClient) GetLectureById(id int64) (lecture *core.LectureInfo, err error) {
	lecture = new(core.LectureInfo)

	err = hsc.requestApi(
		core.LectureApiGetById,
		class100.HttpMethodGet,
		nil,
		core.ApiVersionDefault,
		lecture,
		gox.NewHttpPathParameter("id", strconv.FormatInt(id, 10)),
	)

	return
}

func (hsc *httpSignatureClient) GetsLectureByCourseId(courseId int64) (chapters []*core.ChapterInfo, err error) {
	err = hsc.requestApi(
		core.LectureApiGetByCourseId,
		class100.HttpMethodGet,
		nil,
		core.ApiVersionDefault,
		&chapters,
		gox.NewHttpPathParameter("id", strconv.FormatInt(courseId, 10)),
	)

	return
}

func (hsc *httpSignatureClient) SwitchLectureSequence(req *core.SwitchSequenceReq) (lectures []*core.LectureInfo, err error) {
	err = hsc.requestApi(
		core.LectureSwitchSequence,
		class100.HttpMethodPost,
		req,
		core.ApiVersionDefault,
		&lectures,
	)

	return
}

func (hsc *httpSignatureClient) FirstLectureByCourseId(courseId int64) (lecture *core.LectureInfo, err error) {
	req := core.LectureFirstByCourseIdReq{
		CourseId: courseId,
	}
	lecture = new(core.LectureInfo)
	err = hsc.requestApi(
		core.LectureFirstByCourseId,
		class100.HttpMethodGet,
		req,
		core.ApiVersionDefault,
		&lecture,
	)

	return
}
