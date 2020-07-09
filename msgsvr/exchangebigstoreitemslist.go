package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1proto"
)

type ExchangeBigStoreItemsList struct {
	ItemType        int
	ItemTemplateIds []int
}

func (m ExchangeBigStoreItemsList) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeBigStoreItemsList
}

func (m ExchangeBigStoreItemsList) Serialized() (string, error) {
	itemTemplateIds := make([]string, len(m.ItemTemplateIds))
	for i, v := range m.ItemTemplateIds {
		itemTemplateIds[i] = fmt.Sprintf("%d", v)
	}

	return fmt.Sprintf("%d|%s", m.ItemType, strings.Join(itemTemplateIds, ";")), nil
}

func (m *ExchangeBigStoreItemsList) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) != 2 {
		return d1proto.ErrInvalidMsg
	}

	itemType, err := strconv.ParseInt(sli[0], 10, 32)
	if err != nil {
		return err
	}
	m.ItemType = int(itemType)

	if sli[1] != "" {
		itemTemplateIds := strings.Split(sli[1], ";")
		m.ItemTemplateIds = make([]int, len(itemTemplateIds))
		for i, v := range itemTemplateIds {
			itemTemplateId, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				return err
			}
			m.ItemTemplateIds[i] = int(itemTemplateId)
		}
	}

	return nil
}
