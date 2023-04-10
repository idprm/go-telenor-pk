package entity

import "time"

type Vip struct {
	ID        int       `json:"id"`
	Msisdn    string    `json:"msisdn"`
	CreatedAt time.Time `json:"created_at"`
}

func (v *Vip) GetMsisdn() string {
	return v.Msisdn
}
