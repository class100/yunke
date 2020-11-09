package yunke

import (
	class100 `github.com/class100/core`
	`github.com/class100/yunke-core`
)

type meeting interface {
	// MeetingJoin 加入会议
	MeetingJoin(req *core.JoinMeetingReq) (rsp *core.JoinMeetingRsp, err error)
	// 离开会议
	MeetingLeave(req *core.LeaveMeetingReq) (err error)
}

func (hsc *httpSignatureClient) MeetingJoin(req *core.JoinMeetingReq) (rsp *core.JoinMeetingRsp, err error) {
	rsp = new(core.JoinMeetingRsp)

	err = hsc.requestApi(
		core.MeetingApiJoin,
		class100.HttpMethodPost,
		req,
		core.ApiVersionDefault,
		rsp,
	)

	return
}

func (hsc *httpSignatureClient) MeetingLeave(req *core.LeaveMeetingReq) (err error) {
	err = hsc.requestApi(
		core.MeetingApiLeave,
		class100.HttpMethodPost,
		req,
		core.ApiVersionDefault,
		nil,
	)

	return
}
