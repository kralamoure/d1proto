// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GuildChangeMemberProfile struct{}

func (m GuildChangeMemberProfile) ProtocolId() retroproto.MsgCliId {
	return retroproto.GuildChangeMemberProfile
}

func (m GuildChangeMemberProfile) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildChangeMemberProfile) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
