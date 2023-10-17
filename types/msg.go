package types

import "github.com/gogf/gf/v2/util/gconv"

type MsgType string

const (
	MsgTypeText   MsgType = "txt"
	MsgTypeImg    MsgType = "img"
	MsgTypeCmd    MsgType = "cmd"
	MsgTypeCustom MsgType = "custom"
)

type ChatType string

const (
	Chat      ChatType = "chat"
	GroupChat ChatType = "groupchat"
)

const (
	RouteTypeOnline = "ROUTE_ONLINE"
)

type Message struct {
	From       string      `json:"from"`
	To         string      `json:"to"`
	Type       MsgType     `json:"type"`
	Body       interface{} `json:"body"`
	SyncDevice bool        `json:"sync_device"` // sync 消息发送成功后，是否将消息同步到发送方
	RouteType  string      `json:"routetype"`   // 若传入该参数，其值为 ROUTE_ONLINE，表示接收方只有在线时才能收到消息，若接收方离线则无法收到消息。若不传入该参数，无论接收方在线还是离线都能收到消息。
	MsgExt     `json:"ext"`
}

type MsgExt struct {
	// 是否发送静默消息：
	//- true：是；
	//- （默认）false：否。
	//发送静默消息指用户离线时，环信即时通讯 IM 服务不会通过第三方厂商的消息推送服务向该用户的设备推送消息通知。因此，用户不会收到消息推送通知。当用户再次上线时，会收到离线期间的所有消息。发送静默消息和免打扰模式下均为不推送消息，区别在于发送静默消息为发送方设置不推送消息，而免打扰模式为接收方设置在指定时间段内不接收推送通知。
	EmIgnoreNotification bool `json:"em_ignore_notification"`
}

func (m Message) ToMap() map[string]interface{} {
	data := gconv.MapDeep(m)
	if m.RouteType == "" {
		delete(data, "routetype")
	}
	return data
}

type MsgBodyText struct {
	Msg string `json:"msg"`
}

type MsgBodyCMD struct {
	Action string `json:"action"`
}

type MsgBodyImg struct {
	Filename string `json:"filename"`
	Secret   string `json:"secret"`
	Size     struct {
		Height int `json:"height"`
		Width  int `json:"width"`
	} `json:"size"`
	Url string `json:"url"`
}

type MsgBodyCustom struct {
	CustomEvent string            `json:"customEvent"` // 用户自定义的事件类型。该参数的值必须满足正则表达式 [a-zA-Z0-9-_/\.]{1,32}，长度为 1-32 个字符。
	CustomExts  map[string]string `json:"customExts"`  // 用户自定义的事件属性，类型必须是 Map<String,String>，最多可以包含 16 个元素。customExts 是可选的，不需要可以不传。
}

type DataStringResp struct {
	BaseResp
	Data map[string]string `json:"data"`
}
