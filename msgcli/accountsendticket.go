package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type AccountSendTicket struct {
	Ticket string
}

func (m AccountSendTicket) ProtocolId() retroproto.MsgCliId {
	return retroproto.AccountSendTicket
}

func (m AccountSendTicket) Serialized() (string, error) {
	return m.Ticket, nil
}

func (m *AccountSendTicket) Deserialize(extra string) error {
	m.Ticket = extra

	return nil
}
