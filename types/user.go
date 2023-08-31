package types

import "fmt"

type UserCreateReq struct {
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	Nickname string `yaml:"nickname" json:"nickname"`
}

type BaseResp struct {
	Path      string `json:"path,omitempty"`
	Uri       string `json:"uri,omitempty"`
	Timestamp int64  `json:"timestamp"`
	Action    string `json:"action,omitempty"`
	Duration  int    `json:"duration,omitempty"`

	ApplicationName string `json:"applicationName,omitempty"`
	Organization    string `json:"organization,omitempty"`
	Application     string `json:"application,omitempty"`
}

type UserEntity struct {
	Uuid      string `json:"uuid"`
	Type      string `json:"type"`
	Created   int64  `json:"created"`
	Modified  int64  `json:"modified"`
	Username  string `json:"username"`
	Activated bool   `json:"activated"`
}

type UserResp struct {
	BaseResp
	Entities []*UserEntity `json:"entities"`
}

type UserListResp struct {
	BaseResp
	Entities []*UserEntity `json:"entities"`
	Cursor   string        `json:"cursor"`
	Count    int           `json:"count"`
}

type ListUserReq struct {
	Limit     int64
	Cursor    string
	Activated bool
}

var ListUserReqDefault = ListUserReq{
	Limit:     10,
	Cursor:    "",
	Activated: true,
}

func (s *ListUserReq) BuildQuery() string {
	if s.Limit <= 0 {
		s.Limit = 10
	}
	if s.Limit > 100 {
		s.Limit = 100
	}
	activated := 0
	if s.Activated {
		activated = 1
	}
	return fmt.Sprintf("limit=%d&cursor=%s&activated=%d", s.Limit, s.Cursor, activated)
}

type UserResultResp struct {
	BaseResp
	Data struct {
		Result bool `json:"result"`
	} `json:"data"`
}

type UserStatusResp struct {
	BaseResp
	Data map[string]string `json:"data"`
}

type UsersStatusResp struct {
	BaseResp
	Data []map[string]string `json:"data"`
}

// UserMutesReq 禁言
type UserMutesReq struct {
	Username  string `json:"username"`
	Chat      int    `json:"chat"` // 0单聊消息禁言时长，单位为秒，最大值为 2147483647。 - > 0：该用户 ID 具体的单聊消息禁言时长。 - 0：取消该用户的单聊消息禁言。 - -1：该用户被设置永久单聊消息禁言。
	GroupChat int    `json:"groupchat"`
	Chatroom  int    `json:"chatroom"`
}

type UserMutesData struct {
	Userid    string `json:"userid"`
	Chat      int    `json:"chat"`
	Groupchat int    `json:"groupchat"`
	Chatroom  int    `json:"chatroom"`
	Unixtime  int    `json:"unixtime"`
}
type UserMutesResp struct {
	BaseResp
	Data *UserMutesData `json:"data"`
}

// UserOfflineMsgCountResp 用户离线消息个数
type UserOfflineMsgCountResp struct {
	BaseResp
	Data map[string]int64 `json:"data"`
}

// UserMetadata 用户属性
type UserMetadata struct {
	Nickname  string `json:"nickname"`  // 用户昵称。长度在 64 个字符内。
	Avatarurl string `json:"avatarurl"` // 用户头像 URL 地址。长度在 256 个字符内。
	Phone     string `json:"phone"`     // 用户联系方式。长度在 32 个字符内。
	Mail      string `json:"mail"`      // 用户邮箱。长度在 64 个字符内。
	Gender    int    `json:"gender"`    // 用户性别 0未知 1男 2女
	Sign      string `json:"sign"`      // 用户签名。长度在 256 个字符内。
	Birth     string `json:"birth"`     // 用户生日。长度在 64 个字符内。
	Ext       string `json:"ext"`       // 扩展字段。
}

type UserMetadataResp struct {
	BaseResp
	Data *UserMetadata `json:"data"`
}
