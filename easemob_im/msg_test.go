package easemob_im

import (
	"github.com/xiayuhes/go-easemob/types"
	"testing"
)

func GetMsg(t *testing.T) *Msg {
	auth, err := NewAuth(DefaultEaseMobConfig.AppKey, DefaultEaseMobConfig.ClientID, DefaultEaseMobConfig.ClientSecret, false)
	if err != nil {
		t.Error(err)
	}
	return NewMsg(auth)
}

func TestMsg_SendCustom(t *testing.T) {
	err := GetMsg(t).Send(types.Message{
		From: "12",
		To:   []string{"1"},
		Type: types.MsgTypeCustom,
		Body: types.MsgBodyCustom{
			CustomEvent: "applyFriend",
			CustomExts: map[string]string{
				"txt": "我是王丹蓝",
			},
		},
		SyncDevice: true,
		RouteType:  "",
		Attributes: map[string]interface{}{
			"tex":  1,
			"tex2": true,
		},
	})
	if err != nil {
		t.Error(err)
	}
}

func TestMsg_SendText(t *testing.T) {
	err := GetMsg(t).Send(types.Message{
		From:       "46",
		To:         []string{"8"},
		Type:       types.MsgTypeText,
		Body:       types.MsgBodyText{Msg: "test 你好"},
		SyncDevice: false,
		RouteType:  "",
		Attributes: map[string]interface{}{
			"tex":  1,
			"tex2": true,
		},
	})
	if err != nil {
		t.Error(err)
	}
}
