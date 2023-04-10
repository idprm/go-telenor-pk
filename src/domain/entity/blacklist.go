package entity

import "time"

type Blacklist struct {
	ID        uint64    `json:"id"`
	Msisdn    string    `json:"msisdn"`
	CreatedAt time.Time `json:"created_at"`
}

func (b *Blacklist) GetMsisdn() string {
	return b.Msisdn
}
