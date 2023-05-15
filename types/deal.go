package types

import "time"

type Deal struct {
	Id       int       `json:"id"`
	Symbol   string    `json:"symbol"` // короткое название инструмента
	Exchange string    `json:"exchange"`
	Quantity int       `json:"quantity"` // кол-во инструмента
	Type     string    `json:"type"`     // тип сделки - short/long
	DateOpen time.Time `json:"dateOpen"`
}
