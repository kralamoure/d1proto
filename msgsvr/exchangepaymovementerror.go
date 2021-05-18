// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangePayMovementError struct{}

func (m ExchangePayMovementError) ProtocolId() retroproto.MsgSvrId {
	return retroproto.ExchangePayMovementError
}

func (m ExchangePayMovementError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangePayMovementError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
