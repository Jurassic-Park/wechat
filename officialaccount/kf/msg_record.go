package kf

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/v2/officialaccount/context"
	"github.com/silenceper/wechat/v2/util"
)

// 聊天记录
const (
	msgRecordURL = "https://api.weixin.qq.com/customservice/msgrecord/getmsglist"
)

//MsgRecord struct
type MsgRecord struct {
	*context.Context
}

//NewMsgRecord 实例
func NewMsgRecord(context *context.Context) *MsgRecord {
	msgRecord := new(MsgRecord)
	msgRecord.Context = context
	return msgRecord
}

type ReqGetMsgList struct {
	Starttime int `json:"starttime"`
	Endtime   int `json:"endtime"`
	Msgid     int `json:"msgid"`
	Number    int `json:"number"`
}

type ResGetMsgList struct {
	Recordlist []ResGetMsgListRecord `json:"recordlist"`
	Number     int                   `json:"number"`
	Msgid      int                   `json:"msgid"`
}
type ResGetMsgListRecord struct {
	Openid   string `json:"openid"`
	Opercode int    `json:"opercode"`
	Text     string `json:"text"`
	Time     int    `json:"time"`
	Worker   string `json:"worker"`
}

func (msgRecord *MsgRecord) GetMsgList(req ReqGetMsgList) (*ResGetMsgList, error) {
	accessToken, err := msgRecord.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?access_token=%s", msgRecordURL, accessToken)
	request, _ := json.Marshal(req)
	response, err := util.PostJSON(uri, request)

	var msgList *ResGetMsgList
	err = json.Unmarshal(response, msgList)

	return msgList, err
}
