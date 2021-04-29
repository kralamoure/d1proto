package msgcli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1encoding"
)

type ExchangeBigStoreBuy struct {
	ItemId        int
	QuantityIndex int
	Price         int
}

func (m ExchangeBigStoreBuy) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ExchangeBigStoreBuy
}

func (m ExchangeBigStoreBuy) Serialized() (string, error) {
	return fmt.Sprintf("%d|%d|%d", m.ItemId, m.QuantityIndex, m.Price), nil
}

func (m *ExchangeBigStoreBuy) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) != 3 {
		return d1encoding.ErrInvalidMsg
	}

	itemId, err := strconv.ParseInt(sli[0], 10, 32)
	if err != nil {
		return err
	}
	m.ItemId = int(itemId)

	quantityIndex, err := strconv.ParseInt(sli[1], 10, 32)
	if err != nil {
		return err
	}
	m.QuantityIndex = int(quantityIndex)

	price, err := strconv.ParseInt(sli[2], 10, 32)
	if err != nil {
		return err
	}
	m.Price = int(price)

	return nil
}
