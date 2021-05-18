package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type InfosLifeRestoreTimerFinish struct {
	Restored int
}

func (m InfosLifeRestoreTimerFinish) ProtocolId() retroproto.MsgSvrId {
	return retroproto.InfosLifeRestoreTimerFinish
}

func (m InfosLifeRestoreTimerFinish) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Restored), nil
}

func (m *InfosLifeRestoreTimerFinish) Deserialize(extra string) error {
	if len(extra) < 1 {
		return nil
	}

	restored, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.Restored = int(restored)

	return nil
}
