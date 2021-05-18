package msgcli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retroproto"
)

type GameActionCancel struct {
	Id     int
	Params string
}

func (m GameActionCancel) ProtocolId() retroproto.MsgCliId {
	return retroproto.GameActionCancel
}

func (m GameActionCancel) Serialized() (string, error) {
	return fmt.Sprintf("%d|%s", m.Id, m.Params), nil
}

func (m *GameActionCancel) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")

	if sli[0] == "" {
		return retroproto.ErrInvalidMsg
	}

	id, err := strconv.ParseInt(sli[0], 10, 32)
	if err != nil {
		return err
	}
	m.Id = int(id)

	if len(sli) >= 2 {
		m.Params = sli[1]
	}

	return nil
}
