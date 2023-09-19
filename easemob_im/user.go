package easemob_im

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/xiayuhes/go-easemob/types"
	"strings"
)

type User struct {
	auth *Auth
}

func NewUser(auth *Auth) *User {
	return &User{
		auth: auth,
	}
}

func (u *User) Create(users ...types.UserCreateReq) ([]*types.UserEntity, error) {
	for _, us := range users {
		if err := u.authUser(us); err != nil {
			return nil, err
		}
	}
	uri := u.auth.BuildURI("/users")
	var res types.UserResp
	err := HttpPost(uri, users, &res, u.auth.Headers())
	if !gvar.New(err).IsEmpty() {
		return nil, err
	}
	return res.Entities, nil
}

func (u *User) Get(username string) (*types.UserEntity, error) {
	username = strings.TrimSpace(username)
	if username == "" {
		return nil, errors.New("please enter your username")
	}
	uri := u.auth.BuildURI("/users/" + username)
	var res types.UserListResp
	err := HttpGet(uri, &res, u.auth.Headers())
	if !gvar.New(err).IsEmpty() {
		return nil, err
	}
	return res.Entities[0], err
}
func (u *User) authUser(up types.UserCreateReq) error {
	if up.Username == "" || up.Password == "" {
		return errors.New("please enter your username and password")
	}
	return nil
}

func (u *User) ListUsers(req *types.ListUserReq) ([]*types.UserEntity, string, error) {
	uri := u.auth.GetBaseURI() + "/users?" + req.BuildQuery()
	var res types.UserListResp
	err := HttpGet(uri, &res, u.auth.Headers())
	if !gvar.New(err).IsEmpty() {
		return nil, "", err
	}
	return res.Entities, res.Cursor, nil
}

func (u *User) Delete(username string) (bool, error) {
	username = strings.TrimSpace(username)
	if username == "" {
		return false, errors.New("please enter your username")
	}
	uri := u.auth.BuildURI("/users/" + username)
	var res types.UserResp
	err := HttpDelete(uri, nil, &res, u.auth.Headers())
	if !gvar.New(err).IsEmpty() {
		return false, err
	}
	if len(res.Entities) != 1 {
		return false, errors.New("delete error")
	}
	return res.Entities[0].Activated == false, nil
}

func (u *User) BatchDelete(limit int64) ([]*types.UserEntity, error) {
	if limit <= 0 {
		limit = 1
	}
	uri := u.auth.BuildURI(fmt.Sprintf("/users?limit=%d", limit))
	var res types.UserResp
	err := HttpDelete(uri, nil, &res, u.auth.Headers())
	if !gvar.New(err).IsEmpty() {
		return nil, err
	}
	return res.Entities, nil
}

func (u *User) UpdatePassword(username, password string) error {
	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)
	if username == "" || password == "" {
		return errors.New("please enter your username and password")
	}
	uri := u.auth.BuildURI(fmt.Sprintf("/users/%s/password", username))
	body := map[string]string{
		"newpassword": password,
	}
	var res types.BaseResp
	err := HttpPut(uri, body, &res, u.auth.Headers())
	if !gvar.New(err).IsEmpty() {
		return err
	}
	return nil
}

func (u *User) IsUserOnline(username string) (bool, error) {
	username = strings.TrimSpace(username)
	if username == "" {
		return false, errors.New("please enter your username and password")
	}
	return u.UserStatus(username)
}

func (u *User) IsUsersOnline(usernames []string) (bool, error) {
	if usernames == nil || len(usernames) == 0 {
		return false, errors.New("please enter user name array")
	}
	users, err := u.UsersStatus(usernames)
	if err != nil {
		return false, err
	}

	for _, v := range usernames {
		if o, ok := users[v]; !ok || !o {
			return false, nil
		}
	}
	return true, nil
}

func (u *User) ForceLogoutAllDevices(username string) (bool, error) {
	username = strings.TrimSpace(username)
	if username == "" {
		return false, errors.New("please enter your username and password")
	}
	uri := u.auth.BuildURI(fmt.Sprintf("/users/%s/disconnect", username))
	var res types.UserResultResp
	err := HttpGet(uri, &res, u.auth.Headers())
	if !gvar.New(err).IsEmpty() {
		return false, err
	}
	return res.Data.Result, nil
}

func (u *User) UserStatus(username string) (bool, error) {
	username = strings.TrimSpace(username)
	if username == "" {
		return false, errors.New("please enter your username and password")
	}
	uri := u.auth.BuildURI(fmt.Sprintf("/users/%s/status", username))
	var res types.UserStatusResp
	err := HttpGet(uri, &res, u.auth.Headers())
	if !gvar.New(err).IsEmpty() {
		return false, err
	}
	return res.Data[username] == "online", nil
}

func (u *User) UsersStatus(usernames []string) (map[string]bool, error) {
	uri := u.auth.BuildURI("/users/batch/status")
	body := map[string]interface{}{
		"usernames": usernames,
	}
	var res types.UsersStatusResp
	err := HttpPost(uri, body, &res, u.auth.Headers())
	if !gvar.New(err).IsEmpty() {
		return nil, err
	}
	ret := map[string]bool{}
	for _, uu := range res.Data {
		for dk, du := range uu {
			ret[dk] = false
			if du == "online" {
				ret[du] = true
			}
		}
	}
	return ret, nil
}

func (u *User) UserMutes(req *types.UserMutesReq) (bool, error) {
	uri := u.auth.BuildURI("/mutes")

	var res types.UserResultResp
	err := HttpPost(uri, req, &res, u.auth.Headers())
	if !gvar.New(err).IsEmpty() {
		return false, err
	}
	return res.Data.Result, nil
}

// GetUserMutes 查询单个用户 ID 全局禁言
func (u *User) GetUserMutes(username string) (*types.UserMutesData, error) {
	uri := u.auth.BuildURI("/mutes/" + username)

	var res types.UserMutesResp
	err := HttpGet(uri, &res, u.auth.Headers())
	if !gvar.New(err).IsEmpty() {
		return nil, err
	}
	return res.Data, nil
}

// GetOfflineMsgCount 获取用户离线消息数量
func (u *User) GetOfflineMsgCount(username string) (int64, error) {
	uri := u.auth.BuildURI(fmt.Sprintf("users/%s/offline_msg_count", username))

	var res types.UserOfflineMsgCountResp
	err := HttpGet(uri, &res, u.auth.Headers())
	if !gvar.New(err).IsEmpty() {
		return 0, err
	}
	return res.Data[username], nil
}

func (u *User) SetMetadata(username string, req *types.UserMetadata) (*types.UserMetadata, error) {
	uri := u.auth.BuildURI("/metadata/user/" + username)

	var res types.UserMetadataResp
	err := HttpPut(uri, req, &res, u.auth.Headers())
	if !gvar.New(err).IsEmpty() {
		return nil, err
	}
	return res.Data, nil
}

// SetPushNickname 设置离线推送时显示的昵称
func (u *User) SetPushNickname(username, nickname string) error {
	uri := u.auth.BuildURI("/users/" + username)

	req := make(map[string]string)
	req["nickname"] = nickname

	var res types.BaseResp
	err := HttpPut(uri, &req, &res, u.auth.Headers())
	if !gvar.New(err).IsEmpty() {
		return err
	}
	return nil
}

type UserTokenByUsernameReq struct {
	Username        string
	AutoCreateUser  bool  // 当用户不存在时，是否自动创建用户, 自动创建用户时，需保证授权方式（grant_type）必须为 inherit
	ExpireInSeconds int64 // 设置为 -1 则 token 有效期为永久。若不传该参数，有效期默认为 60 天
}

// GetTokenByUsername 通过用户 ID 获取用户 token
func (u *User) GetTokenByUsername(req *UserTokenByUsernameReq) (*types.AccessTokenResp, error) {
	body := map[string]interface{}{
		"grant_type":     "inherit",
		"username":       req.Username,
		"autoCreateUser": req.AutoCreateUser,
	}
	if req.ExpireInSeconds > 0 {
		body["ttl"] = req.ExpireInSeconds
	} else if req.ExpireInSeconds == -1 {
		body["ttl"] = 0
	}
	uri := u.auth.BuildURI("/token")
	res := types.AccessTokenResp{}
	err := HttpPost(uri, body, &res, u.auth.Headers())
	if !gvar.New(err).IsEmpty() {
		return nil, err
	}
	return &res, nil
}

func (u *User) Deactivate(username string) error {
	uri := u.auth.BuildURI("/users/" + username + "/deactivate")
	var res types.UserResp
	err := HttpPost(uri, nil, &res, u.auth.Headers())
	if !gvar.New(err).IsEmpty() {
		return err
	}
	if len(res.Entities) != 1 {
		return errors.New("deactivate error")
	}
	ok := res.Entities[0].Activated == false
	if !ok {
		return errors.New("deactivate failed")
	}
	return nil
}

func (u *User) Activate(username string) error {
	uri := u.auth.BuildURI("/users/" + username + "/activate")
	var res types.BaseResp
	err := HttpPost(uri, nil, &res, u.auth.Headers())
	if !gvar.New(err).IsEmpty() {
		return err
	}
	return nil
}

func (u *User) AddContact(owner, username string) error {
	uri := u.auth.BuildURI("/users/" + owner + "/contacts/users/" + username)
	var res types.BaseResp
	err := HttpPost(uri, nil, &res, u.auth.Headers())
	if !gvar.New(err).IsEmpty() {
		return err
	}
	return nil
}
