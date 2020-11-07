package yunke

import (
	class100 "github.com/class100/core"
	"github.com/class100/yunke-core"
)

type filer interface {
	// FileUploadInfo 获取文件上传信息
	FileUploadInfo(req *core.UploadFileReq) (rsp *core.FileUploadRsp, err error)
	// GetDownloadInfo 获取文件下载信息
	FileDownloadInfo(req *core.GetDownloadReq) (rsp *core.GetDownloadRsp, err error)
	// FileDelete 删除上传文件
	FileDelete(fileId string) (err error)
}

func (hsc *httpSignatureClient) FileUploadInfo(req *core.UploadFileReq) (rsp *core.FileUploadRsp, err error) {
	rsp = new(core.FileUploadRsp)
	params := make(map[string]string)
	if 0 != len(req.FileId) {
		params["fileId"] = req.FileId
	}
	err = hsc.requestApi(
		core.FileApiUploadGet,
		class100.HttpMethodGet,
		params, nil,
		core.ApiVersionDefault,
		rsp,
	)

	return
}

func (hsc *httpSignatureClient) FileDownloadInfo(req *core.GetDownloadReq) (rsp *core.GetDownloadRsp, err error) {
	rsp = new(core.GetDownloadRsp)
	err = hsc.requestApi(
		core.FileApiDownloadGet,
		class100.HttpMethodGet,
		req, map[string]string{
			"fileId": req.FileId,
		},
		core.ApiVersionDefault,
		rsp,
	)

	return
}

func (hsc *httpSignatureClient) FileDelete(fileId string) (err error) {
	err = hsc.requestApi(
		core.FileApiDelete,
		class100.HttpMethodDelete,
		nil, map[string]string{
			"fileId": fileId,
		},
		core.ApiVersionDefault,
		nil,
	)

	return
}
