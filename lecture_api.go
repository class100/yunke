package yunke

import (
	"fmt"

	class100 "github.com/class100/core"
	"github.com/class100/yunke-core"
)

type lectureHdl interface {
	// LectureAdd 章节添加
	LectureAdd(req *core.AddLectureReq, version core.ApiVersion) (lecture *core.LectureInfo, err error)
	// LectureDelete 章节删除
	LectureDelete(id int64, version core.ApiVersion) (err error)
	// LectureUpdate 章节更新
	LectureUpdate(id int64, req map[string]interface{}, version core.ApiVersion) (lecture *core.LectureInfo, err error)
	// LectureGetById 根据Id获取
	LectureGetById(id int64, version core.ApiVersion) (lecture *core.LectureInfo, err error)
	// LectureGetsByCourseId 根据课程Id获取
	LectureGetsByCourseId(courseId int64, version core.ApiVersion) (chapters []*core.ChapterInfo, err error)
	// LectureSwitchSequence 调整章节顺序
	LectureSwitchSequence(req *core.SwitchSequenceReq, version core.ApiVersion) (lectures []*core.LectureInfo, err error)
	// LectureFirstByCourseId 获取课程第一个章节讲次
	LectureFirstByCourseId(courseId int64, version core.ApiVersion) (lecture *core.LectureInfo, err error)
}

func (hsc *httpSignatureClient) LectureAdd(
	req *core.AddLectureReq,
	version core.ApiVersion,
) (lecture *core.LectureInfo, err error) {
	lecture = new(core.LectureInfo)
	err = hsc.requestApi(core.LectureApiAdd, class100.HttpMethodPost, nil, req, nil, version, lecture)

	return
}

func (hsc *httpSignatureClient) LectureDelete(id int64, version core.ApiVersion) (err error) {
	pathParams := map[string]string{
		"id": fmt.Sprintf("%v", id),
	}
	err = hsc.requestApi(
		core.LectureApiDeleteById,
		class100.HttpMethodDelete,
		nil,
		nil,
		pathParams,
		version,
		nil,
	)

	return
}

func (hsc *httpSignatureClient) LectureUpdate(
	id int64,
	req map[string]interface{},
	version core.ApiVersion,
) (lecture *core.LectureInfo, err error) {
	lecture = new(core.LectureInfo)
	pathParams := map[string]string{
		"id": fmt.Sprintf("%v", id),
	}
	err = hsc.requestApi(
		core.LectureApiUpdate,
		class100.HttpMethodPut,
		nil,
		req,
		pathParams,
		version,
		lecture,
	)

	return
}

func (hsc *httpSignatureClient) LectureGetById(id int64, version core.ApiVersion) (lecture *core.LectureInfo, err error) {
	lecture = new(core.LectureInfo)

	pathParams := map[string]string{
		"id": fmt.Sprintf("%v", id),
	}
	err = hsc.requestApi(
		core.LectureApiGetById,
		class100.HttpMethodGet,
		nil,
		nil,
		pathParams,
		version,
		lecture,
	)

	return
}

func (hsc *httpSignatureClient) LectureGetsByCourseId(
	courseId int64,
	version core.ApiVersion,
) (chapters []*core.ChapterInfo, err error) {
	pathParams := map[string]string{
		"id": fmt.Sprintf("%v", courseId),
	}
	err = hsc.requestApi(
		core.LectureApiGetByCourseId,
		class100.HttpMethodGet,
		nil,
		nil,
		pathParams,
		version,
		&chapters,
	)

	return
}

func (hsc *httpSignatureClient) LectureSwitchSequence(
	req *core.SwitchSequenceReq,
	version core.ApiVersion,
) (lectures []*core.LectureInfo, err error) {
	err = hsc.requestApi(
		core.LectureSwitchSequence,
		class100.HttpMethodPost,
		nil,
		req,
		nil,
		version,
		&lectures,
	)

	return
}

func (hsc *httpSignatureClient) LectureFirstByCourseId(
	courseId int64,
	version core.ApiVersion,
) (lecture *core.LectureInfo, err error) {
	params := map[string]string{
		"courseId": fmt.Sprintf("%v", courseId),
	}
	lecture = new(core.LectureInfo)
	err = hsc.requestApi(
		core.LectureFirstByCourseId,
		class100.HttpMethodGet,
		nil,
		nil,
		params,
		version,
		&lecture,
	)

	return
}
