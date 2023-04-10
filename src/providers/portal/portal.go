package portal

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/idprm/go-telenor-pk/src/config"
	"github.com/idprm/go-telenor-pk/src/domain/entity"
	"github.com/idprm/go-telenor-pk/src/logger"
	"github.com/sirupsen/logrus"
)

type Portal struct {
	cfg          *config.Secret
	logger       *logger.Logger
	service      *entity.Service
	subscription *entity.Subscription
}

func NewPortal(cfg *config.Secret, logger *logger.Logger, service *entity.Service, subscription *entity.Subscription) *Portal {
	return &Portal{
		cfg:          cfg,
		logger:       logger,
		service:      service,
		subscription: subscription,
	}
}

func (p *Portal) Subscription() ([]byte, error) {
	l := p.logger.Init("notif", true)

	urlAPI := p.service.GetUrlNotifSub()

	payload := url.Values{}
	payload.Add("msisdn", p.subscription.GetMsisdn())
	payload.Add("trxid", "")
	payload.Add("package", "weekly")
	payload.Add("event", "reg")

	req, err := http.NewRequest("GET", urlAPI+"?"+payload.Encode(), nil)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	l.WithFields(logrus.Fields{
		"request_url": urlAPI + "?" + payload.Encode(),
		"msisdn":      p.subscription.GetMsisdn(),
	}).Info("")

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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return body, nil
}

func (p *Portal) UnSubscription() ([]byte, error) {
	l := p.logger.Init("notif", true)

	urlAPI := p.service.GetUrlNotifUnsub()

	payload := url.Values{}
	payload.Add("msisdn", p.subscription.GetMsisdn())
	payload.Add("event", "unreg")

	req, err := http.NewRequest("GET", urlAPI+"?"+payload.Encode(), nil)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	l.WithFields(logrus.Fields{
		"request_url": urlAPI + "?" + payload.Encode(),
		"msisdn":      p.subscription.GetMsisdn(),
	}).Info("")

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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return body, nil
}

func (p *Portal) Renewal() ([]byte, error) {
	l := p.logger.Init("notif", true)

	urlAPI := p.service.GetUrlNotifRenewal()

	payload := url.Values{}
	payload.Add("msisdn", p.subscription.GetMsisdn())
	payload.Add("trxid", "")
	payload.Add("package", "weekly")
	payload.Add("event", "renewal")

	req, err := http.NewRequest("GET", urlAPI+"?"+payload.Encode(), nil)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	l.WithFields(logrus.Fields{
		"request_url": urlAPI + "?" + payload.Encode(),
		"msisdn":      p.subscription.GetMsisdn(),
	}).Info("")

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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return body, nil
}
