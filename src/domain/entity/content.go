package entity

type Content struct {
	ID        int `json:"id"`
	ServiceID int `json:"service_id"`
	Service   *Service
	Name      string `json:"name"`
	Value     string `json:"value"`
}

func (c *Content) GetName() string {
	return c.Name
}

func (c *Content) GetValue() string {
	return c.Value
}
