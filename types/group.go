package types

type GroupCreateReq struct {
	GroupName         string   `json:"groupname"`
	Description       string   `json:"description"`
	Public            bool     `json:"public"`
	Scale             string   `json:"scale"`
	MaxUsers          int64    `json:"maxusers"`
	AllowInvites      bool     `json:"allowinvites"`
	MembersOnly       bool     `json:"membersonly"`
	InviteNeedConfirm bool     `json:"invite_need_confirm"`
	Owner             string   `json:"owner"`
	Members           []string `json:"members"`
	Custom            string   `json:"custom"`
}

type GroupResp struct {
	BaseResp
	Data map[string]string `json:"data"`
}

type GroupDataBoolResp struct {
	BaseResp
	Data map[string]bool `json:"data"`
}

type GroupDataAnyResp struct {
	BaseResp
	Data map[string]interface{} `json:"data"`
}

type GroupAdminListResp struct {
	BaseResp
	Data []string `json:"data"`
}

type GroupEditReq struct {
	GroupName         string `json:"groupname"`
	Description       string `json:"description"`
	Public            bool   `json:"public"`
	MaxUsers          int64  `json:"maxusers"`
	AllowInvites      bool   `json:"allowinvites"`
	MembersOnly       bool   `json:"membersonly"`
	InviteNeedConfirm bool   `json:"invite_need_confirm"`
	Custom            string `json:"custom"`
}

type GroupEntity struct {
	Id                string `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	MembersOnly       bool   `json:"membersonly"`
	AllowInvites      bool   `json:"allowinvites"`
	MaxUsers          int64  `json:"maxusers"`
	Permission        string `json:"permission"`
	Owner             string `json:"owner"`
	Created           int64  `json:"created"`
	AffiliationsCount int64  `json:"affiliations_count"`
	Disabled          bool   `json:"disabled"`
	Mute              bool   `json:"mute"`
	Public            bool   `json:"public"`
	Custom            string `json:"custom"`
}
