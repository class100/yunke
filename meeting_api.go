package yunke

import (
	class100 `github.com/class100/core`
	`github.com/class100/yunke-core`
)

type meeting interface {
	// JoinMeeting 加入会议
	JoinMeeting(req *core.JoinMeetingReq) (rsp *core.JoinMeetingRsp, err error)
	// LeaveMeeting 离开会议
	LeaveMeeting(req *core.LeaveMeetingReq) (err error)
}

func (hsc *httpSignatureClient) JoinMeeting(req *core.JoinMeetingReq) (rsp *core.JoinMeetingRsp, err error) {
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

func (hsc *httpSignatureClient) LeaveMeeting(req *core.LeaveMeetingReq) (err error) {
	err = hsc.requestApi(
		core.MeetingApiLeave,
		class100.HttpMethodPost,
		req,
		core.ApiVersionDefault,
		nil,
	)

	return
}
