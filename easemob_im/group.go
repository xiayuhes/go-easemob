package easemob_im

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/xiayuhes/go-easemob/types"
)

type Group struct {
	auth *Auth
}

func NewGroup(auth *Auth) *Group {
	return &Group{
		auth: auth,
	}
}

func (s *Group) Create(body types.GroupCreateReq) (groupId string, err error) {
	uri := s.auth.BuildURI("/chatgroups")
	var res types.GroupResp
	err = HttpPost(uri, body, &res, s.auth.Headers())
	if !gvar.New(err).IsEmpty() {
		return
	}
	return res.Data["groupid"], nil
}

func (s *Group) Disable(groupId string) (err error) {
	uri := s.auth.BuildURI(fmt.Sprintf("/chatgroups/%s/disable", groupId))
	var res types.DataBoolResp
	err = HttpPost(uri, nil, &res, s.auth.Headers())
	if !gvar.New(err).IsEmpty() {
		return
	}
	if res.Data["disabled"] {
		return nil
	}
	return errors.New("disable error")
}

func (s *Group) Enable(groupId string) (err error) {
	uri := s.auth.BuildURI(fmt.Sprintf("/chatgroups/%s/enable", groupId))
	var res types.DataBoolResp
	err = HttpPost(uri, nil, &res, s.auth.Headers())
	if !gvar.New(err).IsEmpty() {
		return
	}
	if !res.Data["disabled"] {
		return nil
	}
	return errors.New("enable error")
}

func (s *Group) Edit(groupId string, body types.GroupEditReq) error {
	uri := s.auth.BuildURI(fmt.Sprintf("/chatgroups/%s", groupId))
	var res types.DataBoolResp
	err := HttpPut(uri, body, &res, s.auth.Headers())
	if !gvar.New(err).IsEmpty() {
		return err
	}
	for k, v := range res.Data {
		if v == false {
			return fmt.Errorf("edit %s failed", k)
		}
	}
	return nil
}

func (s *Group) Delete(groupId string) error {
	uri := s.auth.BuildURI(fmt.Sprintf("/chatgroups/%s", groupId))
	var res types.DataAnyResp
	err := HttpDelete(uri, nil, &res, s.auth.Headers())
	if !gvar.New(err).IsEmpty() {
		return err
	}

	if res.Data["success"].(bool) {
		return nil
	}
	return errors.New("delete failed")
}

func (s *Group) ChangeOwner(groupId, newOwner string) error {
	uri := s.auth.BuildURI(fmt.Sprintf("/chatgroups/%s", groupId))
	var res types.DataBoolResp
	req := make(map[string]string)
	req["newowner"] = newOwner
	err := HttpPut(uri, req, &res, s.auth.Headers())
	if !gvar.New(err).IsEmpty() {
		return err
	}
	if res.Data["newowner"] {
		return nil
	}
	return errors.New("change owner failed")
}

func (s *Group) AdminList(groupId string) (out []string) {
	uri := s.auth.BuildURI(fmt.Sprintf("/chatgroups/%s/admin", groupId))
	var res types.GroupAdminListResp
	err := HttpGet(uri, &res, s.auth.Headers())
	if !gvar.New(err).IsEmpty() {
		return
	}
	out = res.Data
	return
}

func (s *Group) AddAdmin(groupId string, admin string) error {
	uri := s.auth.BuildURI(fmt.Sprintf("/chatgroups/%s/admin", groupId))
	var res types.GroupResp
	req := make(map[string]string)
	req["newadmin"] = admin
	err := HttpPost(uri, req, &res, s.auth.Headers())
	if !gvar.New(err).IsEmpty() {
		return err
	}
	if res.Data["result"] == "success" {
		return nil
	}
	return err
}

func (s *Group) RemoveAdmin(groupId string, admin string) error {
	uri := s.auth.BuildURI(fmt.Sprintf("/chatgroups/%s/admin/%s", groupId, admin))
	var res types.GroupResp
	err := HttpDelete(uri, nil, &res, s.auth.Headers())
	if !gvar.New(err).IsEmpty() {
		return err
	}
	if res.Data["result"] == "success" {
		return nil
	}
	return err
}

func (s *Group) AddUsername(groupId string, username string) error {
	uri := s.auth.BuildURI(fmt.Sprintf("/chatgroups/%s/users/%s", groupId, username))
	var res types.DataAnyResp
	err := HttpPost(uri, nil, &res, s.auth.Headers())
	if !gvar.New(err).IsEmpty() {
		return err
	}
	return nil
}

func (s *Group) AddUsernames(groupId string, usernames []string) error {
	uri := s.auth.BuildURI(fmt.Sprintf("/chatgroups/%s/users", groupId))
	var res types.DataAnyResp
	req := make(map[string]interface{})
	req["usernames"] = usernames
	err := HttpPost(uri, req, &res, s.auth.Headers())
	if !gvar.New(err).IsEmpty() {
		return err
	}
	return nil
}

func (s *Group) RemoveUsername(groupId string, username string) error {
	uri := s.auth.BuildURI(fmt.Sprintf("/chatgroups/%s/users/%s", groupId, username))
	var res types.DataAnyResp
	err := HttpDelete(uri, nil, &res, s.auth.Headers())
	if !gvar.New(err).IsEmpty() {
		return err
	}
	return nil
}

func (s *Group) RemoveUsernames(groupId string, usernames string) error {
	uri := s.auth.BuildURI(fmt.Sprintf("/chatgroups/%s/users/%s", groupId, usernames))
	var res types.DataAnyResp
	err := HttpDelete(uri, nil, &res, s.auth.Headers())
	if !gvar.New(err).IsEmpty() {
		return err
	}
	return nil
}
