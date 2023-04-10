package entity

type Service struct {
	ID              int     `json:"id"`
	Category        string  `json:"category"`
	Code            string  `json:"code"`
	Name            string  `json:"name"`
	Package         string  `json:"package"`
	AuthUser        string  `json:"auth_user"`
	AuthPass        string  `json:"auth_pass"`
	PartnerID       string  `json:"partner_id"`
	ProductID       string  `json:"product_id"`
	TransID         string  `json:"trans_id"`
	CorrelationID   string  `json:"correlation_id"`
	Price           float64 `json:"price"`
	RenewalDay      int     `json:"renewal_day"`
	TrialDay        int     `json:"trial_day"`
	UrlHost         string  `json:"url_host"`
	UrlTelco        string  `json:"url_telco"`
	UrlCallback     string  `json:"url_callback"`
	UrlNotifSub     string  `json:"url_notif_sub"`
	UrlNotifUnsub   string  `json:"url_notif_unsub"`
	UrlNotifRenewal string  `json:"url_notif_renewal"`
	UrlPostback     string  `json:"url_postback"`
}

func (s *Service) GetAuthUser() string {
	return s.AuthUser
}

func (s *Service) GetAuthPass() string {
	return s.AuthPass
}

func (s *Service) GetPartnerID() string {
	return s.PartnerID
}

func (s *Service) GetProductID() string {
	return s.ProductID
}

func (s *Service) GetTransID() string {
	return s.TransID
}

func (s *Service) GetCorrelationID() string {
	return s.CorrelationID
}

func (s *Service) GetPrice() float64 {
	return s.Price
}

func (s *Service) GetRenewalDay() int {
	return s.RenewalDay
}

func (s *Service) GetTrialDay() int {
	return s.TrialDay
}

func (s *Service) GetUrlHost() string {
	return s.UrlHost
}

func (s *Service) GetUrlTelco() string {
	return s.UrlTelco
}

func (s *Service) GetUrlCallback() string {
	return s.UrlCallback
}

func (s *Service) GetUrlNotifSub() string {
	return s.UrlNotifSub
}

func (s *Service) GetUrlNotifUnsub() string {
	return s.UrlNotifUnsub
}

func (s *Service) GetUrlNotifRenewal() string {
	return s.UrlNotifRenewal
}

func (s *Service) GetUrlPostback() string {
	return s.UrlPostback
}
