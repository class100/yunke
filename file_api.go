package yunke

import (
	class100 "github.com/class100/core"
	"github.com/class100/yunke-core"
	`github.com/storezhang/gox`
)

type filer interface {
	// GetUploadFileInfo 获取文件上传信息
	GetUploadFileInfo(req *core.BaseUploadReq) (rsp *core.BaseUploadReq, err error)
	// GetDownloadFileInfo 获取文件下载信息
	GetDownloadFileInfo(req *core.BaseDownloadReq) (rsp *core.BaseDownloadRsp, err error)
	// DeleteFile 删除上传文件
	DeleteFile(fileId string) (err error)
}

func (hsc *httpSignatureClient) GetUploadFileInfo(req *core.BaseUploadReq) (rsp *core.BaseUploadReq, err error) {
	rsp = new(core.BaseUploadReq)
	params := make(map[string]string)
	if nil != req.FileId && 0 != len(*req.FileId) {
		params["fileId"] = *req.FileId
	}
	err = hsc.requestApi(
		core.FileApiUploadGet,
		class100.HttpMethodGet,
		params,
		core.ApiVersionDefault,
		rsp,
	)

	return
}

func (hsc *httpSignatureClient) GetDownloadFileInfo(req *core.BaseDownloadReq) (rsp *core.BaseDownloadRsp, err error) {
	rsp = new(core.BaseDownloadRsp)
	err = hsc.requestApi(
		core.FileApiDownloadGet,
		class100.HttpMethodGet,
		req,
		core.ApiVersionDefault,
		rsp,
		gox.NewHttpPathParameter("fileId", req.FileId),
	)

	return
}

func (hsc *httpSignatureClient) DeleteFile(fileId string) (err error) {
	err = hsc.requestApi(
		core.FileApiDelete,
		class100.HttpMethodDelete,
		nil,
		core.ApiVersionDefault,
		nil,
		gox.NewHttpPathParameter("fileId", fileId),
	)

	return
}
