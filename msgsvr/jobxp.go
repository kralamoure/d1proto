// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type JobXP struct{}

func (m JobXP) ProtocolId() retroproto.MsgSvrId {
	return retroproto.JobXP
}

func (m JobXP) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *JobXP) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
