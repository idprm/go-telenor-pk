package entity

type Campaign struct {
	ID        uint `json:"id"`
	ServiceID int  `json:"service_id"`
	Service   *Service
	Adnet     string `json:"adnet"`
	LimitMo   int    `json:"limit_mo"`
	TotalMo   int    `json:"total_mo"`
}
