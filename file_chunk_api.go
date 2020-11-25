package yunke

import (
	class100 `github.com/class100/core`
	`github.com/class100/yunke-core`
)

type fileChunk interface {
	// InitiateMultipartUpload 初始化分块上传信息
	InitiateMultipartUpload(req *core.InitiateMultipartUploadReq) (rsp *core.InitiateMultipartUploadRsp, err error)
	// CompleteMultipartUpload 完成化分块上传信息
	CompleteMultipartUpload(req *core.CompleteMultipartUploadReq) (err error)
	// AbortMultipartUpload 取消分块上传
	AbortMultipartUpload(req *core.AbortMultipartUploadReq) (err error)
}

func (hsc *httpSignatureClient) InitiateMultipartUpload(req *core.InitiateMultipartUploadReq) (rsp *core.InitiateMultipartUploadRsp, err error) {
	rsp = new(core.InitiateMultipartUploadRsp)
	err = hsc.requestApi(
		core.FileChunkApiAbortMultipartUploadReq,
		class100.HttpMethodPost,
		req,
		core.ApiVersionDefault,
		rsp,
	)

	return
}

func (hsc *httpSignatureClient) CompleteMultipartUpload(req *core.CompleteMultipartUploadReq) (err error) {
	err = hsc.requestApi(
		core.FileChunkApiCompleteMultipartUpload,
		class100.HttpMethodPut,
		req,
		core.ApiVersionDefault,
		nil,
	)

	return
}

func (hsc *httpSignatureClient) AbortMultipartUpload(req *core.AbortMultipartUploadReq) (err error) {
	err = hsc.requestApi(
		core.FileApiDelete,
		class100.HttpMethodDelete,
		req,
		core.ApiVersionDefault,
		nil,
	)

	return
}
