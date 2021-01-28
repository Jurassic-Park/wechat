package kf

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/silenceper/wechat/v2/officialaccount/context"
	"github.com/silenceper/wechat/v2/util"
)

// https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Customer_Service_Management.html
const (
	// 所有客服
	kfListURL = "https://api.weixin.qq.com/cgi-bin/customservice/getkflist"
	// 在线客服
	kfOnlineListURL = "https://api.weixin.qq.com/cgi-bin/customservice/getonlinekflist"
	// 添加客服
	kfAddURL = "https://api.weixin.qq.com/customservice/kfaccount/add"
	// 邀请绑定客服帐号
	kfInviteWorkerURL = "https://api.weixin.qq.com/customservice/kfaccount/inviteworker"
	// 更新客服信息
	kfUpdateURL = "https://api.weixin.qq.com/customservice/kfaccount/update"
	// 上传客服头像
	kfUploadHeadimgURL = "https://api.weixin.qq.com/customservice/kfaccount/uploadheadimg"
	// 删除客服
	kfDeleteURL = "https://api.weixin.qq.com/customservice/kfaccount/del"
)

//Kf struct
type Kf struct {
	*context.Context
}

// ResBaseInformation 基本信息返回
type ResBaseInformation struct {
	KfList []ResBaseInformationSingle `json:"kf_list"`
}
type ResBaseInformationSingle struct {
	KfAccount        string `json:"kf_account"`
	KfHeadimgurl     string `json:"kf_headimgurl"`
	KfID             string `json:"kf_id"`
	KfNick           string `json:"kf_nick"`
	KfWx             string `json:"kf_wx,omitempty"`
	InviteWx         string `json:"invite_wx,omitempty"`
	InviteExpireTime int    `json:"invite_expire_time,omitempty"`
	InviteStatus     string `json:"invite_status,omitempty"`
}

// 在线信息返回
type ResOnlineInformation struct {
	KfOnlineList []ResOnlineInformationSingle `json:"kf_online_list"`
}
type ResOnlineInformationSingle struct {
	KfAccount    string `json:"kf_account"`
	Status       int    `json:"status"`
	KfID         string `json:"kf_id"`
	AcceptedCase int    `json:"accepted_case"`
}

// 添加/更新客服
type ReqAddUpdate struct {
	KfAccount string `json:"kf_account"`
	Nickname  string `json:"nickname"`
}

// 邀请绑定
type ReqInviteWorker struct {
	KfAccount string `json:"kf_account"`
	InviteWx  string `json:"invite_wx"`
}

// 上传头像
type ReqUploadimg struct {
	KfAccount string `json:"kf_account"`
	Media     string `json:"media"` // base64
}

// 删除
type ReqDelete struct {
	KfAccount string `json:"kf_account"`
}

//NewKf 实例
func NewKf(context *context.Context) *Kf {
	kf := new(Kf)
	kf.Context = context
	return kf
}

// 获取所有客服
func (kf *Kf) GetKfList() (*ResBaseInformation, error) {
	accessToken, err := kf.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?access_token=%s", kfListURL, accessToken)
	response, err := util.HTTPGet(uri)

	var res *ResBaseInformation
	err = json.Unmarshal(response, res)

	return res, err
}

// 获取在线客服
func (kf *Kf) GetOnlineKfList() (*ResOnlineInformation, error) {
	accessToken, err := kf.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?access_token=%s", kfOnlineListURL, accessToken)
	response, err := util.HTTPGet(uri)

	var res *ResOnlineInformation
	err = json.Unmarshal(response, res)

	return res, err
}

// 添加客服
func (kf *Kf) AddKf(req ReqAddUpdate) error {
	accessToken, err := kf.GetAccessToken()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s?access_token=%s", kfAddURL, accessToken)
	request, _ := json.Marshal(req)
	response, err := util.PostJSON(uri, request)

	return util.DecodeWithCommonError(response, "AddKf")
}

// 更新客服
func (kf *Kf) UpdateKf(req ReqAddUpdate) error {
	accessToken, err := kf.GetAccessToken()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s?access_token=%s", kfUpdateURL, accessToken)
	request, _ := json.Marshal(req)
	response, err := util.PostJSON(uri, request)

	return util.DecodeWithCommonError(response, "UpdateKf")
}

// 邀请绑定
func (kf *Kf) InviteKf(req ReqInviteWorker) error {
	accessToken, err := kf.GetAccessToken()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s?access_token=%s", kfInviteWorkerURL, accessToken)
	request, _ := json.Marshal(req)
	response, err := util.PostJSON(uri, request)

	return util.DecodeWithCommonError(response, "InviteKf")
}

// 上传客服头像
func (kf *Kf) UploadHeadimg(req ReqUploadimg) error {
	accessToken, err := kf.GetAccessToken()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s?access_token=%s&kf_account=%s", kfUploadHeadimgURL, accessToken, req.KfAccount)

	filename := fmt.Sprintf("%d", time.Now().Unix())
	var response []byte
	response, err = util.PostFileByBase64("media", filename, req.Media, uri)
	if err != nil {
		return err
	}
	return util.DecodeWithCommonError(response, "UploadHeadimg")
}

// 删除客服
func (kf *Kf) DeleteKf(req ReqDelete) error {
	accessToken, err := kf.GetAccessToken()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s?access_token=%s&kf_account=%s", kfDeleteURL, accessToken, req.KfAccount)
	response, err := util.HTTPGet(uri)

	if err != nil {
		return err
	}

	return util.DecodeWithCommonError(response, "DeleteKf")
}
