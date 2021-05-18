// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GameZoneData struct{}

func (m GameZoneData) ProtocolId() retroproto.MsgSvrId {
	return retroproto.GameZoneData
}

func (m GameZoneData) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameZoneData) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
