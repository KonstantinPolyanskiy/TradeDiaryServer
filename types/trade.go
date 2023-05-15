package types

import (
	"github.com/shopspring/decimal"
	"time"
)

type Trade struct {
	Id            int             `json:"id"`
	InstrumentId  int             `json:"instrumentId"`
	TimestampOpen time.Time       `json:"timestampOpen"` // время открытия
	Price         decimal.Decimal `json:"price"`
	Quantity      int             `json:"quantity"`  // Кол-во контрактов
	Direction     string          `json:"direction"` // Направление: long/short
	Commission    decimal.Decimal `json:"commission"`
}
