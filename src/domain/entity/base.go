package entity

import (
	"strings"

	"github.com/idprm/go-telenor-pk/src/config"
)

type MORequest struct {
	cfg     *config.Secret
	Msisdn  string `json:"msisdn"`
	Service string `json:"service"`
}

type ChargeRequestBody struct {
	Msisdn          string  `json:"msisdn"`
	ChargableAmount float64 `json:"chargableAmount"`
	CorrelationID   string  `json:"correlationId"`
	PartnerID       string  `json:"PartnerID"`
}

type MessageRequestBody struct {
	Msisdn  string `json:"recipientMsisdn"`
	Message string `json:"messageBody"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}

type ActivateRequestBody struct {
	Msisdn    string `json:"msisdn"`
	ServiceId string `json:"serviceId"`
	Channel   string `json:"channel"`
}

type ActivateResponse struct {
	Status            string `json:"status"`
	ActivationTime    int    `json:"activationTime"`
	ExpireTime        int    `json:"expireTime"`
	ActivationChannel string `json:"activationChannel"`
}

func (s *MORequest) IsValidPrefix() bool {
	return strings.HasPrefix(s.Msisdn, s.cfg.Telco.Prefix)
}

func (s *MORequest) IsValidVIP() bool {
	return strings.HasPrefix(s.Msisdn, s.cfg.Telco.PrefixVIP)
}

func (t *TokenResponse) GetAccessToken() string {
	return t.AccessToken
}

func (t *TokenResponse) SetAccessToken(accessToken string) {
	t.AccessToken = accessToken
}

func (t *TokenResponse) GetExpiresIn() string {
	return t.ExpiresIn
}
