// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type HousesSellError struct{}

func (m HousesSellError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.HousesSellError
}

func (m HousesSellError) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *HousesSellError) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
