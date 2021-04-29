package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type BasicsGetDate struct{}

func (m BasicsGetDate) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.BasicsGetDate
}

func (m BasicsGetDate) Serialized() (string, error) {
	return "", nil
}

func (m *BasicsGetDate) Deserialize(extra string) error {
	return nil
}
