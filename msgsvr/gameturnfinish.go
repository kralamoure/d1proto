// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GameTurnFinish struct{}

func (m GameTurnFinish) ProtocolId() retroproto.MsgSvrId {
	return retroproto.GameTurnFinish
}

func (m GameTurnFinish) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameTurnFinish) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
