package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1encoding"
)

type SpellsUpgradeSpellSuccess struct {
	Id    int
	Level int
}

func (m SpellsUpgradeSpellSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.SpellsUpgradeSpellSuccess
}

func (m SpellsUpgradeSpellSuccess) Serialized() (string, error) {
	return fmt.Sprintf("%d~%d", m.Id, m.Level), nil
}

func (m *SpellsUpgradeSpellSuccess) Deserialize(extra string) error {
	sli := strings.Split(extra, "~")
	if len(sli) != 2 {
		return d1encoding.ErrInvalidMsg
	}

	id, err := strconv.ParseInt(sli[0], 10, 32)
	if err != nil {
		return err
	}
	m.Id = int(id)

	level, err := strconv.ParseInt(sli[1], 10, 32)
	if err != nil {
		return err
	}
	m.Level = int(level)

	return nil
}
