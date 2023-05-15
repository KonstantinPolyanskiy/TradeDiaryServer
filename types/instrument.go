package types

import (
	"github.com/shopspring/decimal"
	"time"
)

type Instrument struct {
	Id               int             `json:"id"`
	Name             string          `json:"name"`
	ExpirationDate   time.Time       `json:"expirationDate"`
	UnderlyingSymbol string          `json:"underlyingSymbol"`
	ContractSize     int             `json:"contractSize"`
	TickSize         decimal.Decimal `json:"tickSize"`
}
