package telco

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/idprm/go-telenor-pk/src/config"
	"github.com/idprm/go-telenor-pk/src/domain/entity"
)

type Telco struct {
	cfg          *config.Secret
	service      *entity.Service
	content      *entity.Content
	subscription *entity.Subscription
}

func NewTelco(cfg *config.Secret, service *entity.Service, content *entity.Content, subscription *entity.Subscription) *Telco {
	return &Telco{
		cfg:          cfg,
		service:      service,
		content:      content,
		subscription: subscription,
	}
}

func (t *Telco) Token() (string, error) {
	payload := url.Values{}
	payload.Add("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", t.service.GetUrlTelco()+payload.Encode(), nil)

	req.SetBasicAuth(t.service.GetAuthUser(), t.service.GetAuthPass())
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return "", errors.New(err.Error())
	}

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: tr,
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New(err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", errors.New(err.Error())
	}

	bodyString := string(body)

	var res entity.TokenResponse
	json.Unmarshal([]byte(bodyString), &res)

	return res.GetAccessToken(), nil
}

func (t *Telco) Charge() ([]byte, error) {

	accessToken, err := t.Token()
	if err != nil {
		return nil, err
	}

	var bearer = "Bearer " + accessToken

	data := &entity.ChargeRequestBody{
		Msisdn:          t.subscription.GetMsisdn(),
		ChargableAmount: t.service.GetPrice(),
		CorrelationID:   t.service.GetCorrelationID(),
		PartnerID:       t.service.GetPartnerID(),
	}

	jsonStr, err := json.Marshal(data)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	req, err := http.NewRequest("POST", t.service.GetUrlTelco()+"/payment/v1/charge", bytes.NewBuffer(jsonStr))

	if err != nil {
		return nil, errors.New(err.Error())
	}

	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return nil, errors.New(err.Error())
	}

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: tr,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return res, nil
}

func (t *Telco) Message() ([]byte, error) {

	accessToken, err := t.Token()
	if err != nil {
		return nil, err
	}

	var bearer = "Bearer " + accessToken

	data := &entity.MessageRequestBody{
		Msisdn:  t.subscription.GetMsisdn(),
		Message: t.content.GetValue(),
	}

	jsonStr, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", t.service.GetUrlTelco()+"/sms/v1/send", bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return nil, err
	}

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: tr,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (t *Telco) Activate() ([]byte, error) {
	accessToken, err := t.Token()
	if err != nil {
		return nil, err
	}

	var bearer = "Bearer " + accessToken

	data := &entity.ActivateRequestBody{
		Msisdn:    t.subscription.GetMsisdn(),
		ServiceId: "",
		Channel:   "API",
	}

	jsonStr, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", t.service.GetUrlTelco()+"/dpdp/v1/subscriber", bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return nil, err
	}

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: tr,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (t *Telco) Deactive() ([]byte, error) {
	accessToken, err := t.Token()
	if err != nil {
		return nil, err
	}

	var bearer = "Bearer " + accessToken

	data := &entity.ActivateRequestBody{
		Msisdn:    t.subscription.GetMsisdn(),
		ServiceId: "",
		Channel:   "API",
	}

	jsonStr, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", t.service.GetUrlTelco()+"/dpdp/v1/subscriber", bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return nil, err
	}

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: tr,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
