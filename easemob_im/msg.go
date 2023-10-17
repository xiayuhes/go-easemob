package easemob_im

import (
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
