package easemob_im

import (
	"errors"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/xiayuhes/go-easemob/types"
)

type Msg struct {
	auth *Auth
}

func NewMsg(auth *Auth) *Msg {
	return &Msg{
		auth: auth,
	}
}

func (s *Msg) Send(msg types.Message) error {
	body := msg.ToMap()
	uri := s.auth.BuildURI("/messages/users")
	var res types.BaseResp
	err := HttpPost(uri, body, &res, s.auth.Headers())
	if !gvar.New(err).IsEmpty() {
		return err
	}
	return nil
}

// Delete
// channel	String	是	要删除的会话 ID。该参数的值取决于会话类型 type 的值:
// - type 为 chat，即单聊时，会话 ID 为对端用户 ID；
// - type 为 groupchat，即群聊时，会话 ID 为群组 ID。
// type	String	是	会话类型。
// - chat：单聊会话；
// -groupchat：群聊会话。
// delete_roam	Bool	是	是否删除该会话在服务端的漫游消息。
// - true：是。若删除了该会话的服务端消息，则用户无法从服务器拉取该会话的漫游消息。
// - false：否。用户仍可以从服务器拉取该会话的漫游消息。
func (s *Msg) Delete(username, channel string, chatType types.ChatType, deleteRoam bool) error {
	body := map[string]interface{}{
		"channel":     channel,
		"type":        chatType,
		"delete_roam": deleteRoam,
	}
	uri := s.auth.BuildURI("/users/" + username + "/user_channel")
	var res types.DataStringResp
	err := HttpDelete(uri, body, &res, s.auth.Headers())
	if !gvar.New(err).IsEmpty() {
		return err
	}
	if res.Data["result"] != "ok" {
		return errors.New(res.Data["result"])
	}
	return nil
}
