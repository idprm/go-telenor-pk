package repository

import (
	"database/sql"

	"github.com/idprm/go-telenor-pk/src/domain/entity"
)

const (
	queryCountService  = "SELECT COUNT(*) as count FROM services WHERE category = $1 AND package = $2"
	querySelectService = "SELECT id, category, code, name, package, auth_user, auth_pass, partner_id, product_id, trans_id, correlation_id, price, renewal_day, trial_day, url_host, url_callback, url_notif_sub, url_notif_unsub, url_notif_renewal, url_postback FROM services WHERE id = $1 LIMIT 1"
	querySelectPackage = "SELECT id, category, code, name, package, auth_user, auth_pass, partner_id, product_id, trans_id, correlation_id, price, renewal_day, trial_day, url_host, url_callback, url_notif_sub, url_notif_unsub, url_notif_renewal, url_postback FROM services WHERE code = $1 LIMIT 1"
)

type ServiceRepository struct {
	db *sql.DB
}

type IServiceRepository interface {
	Count(string, string) (int, error)
	GetById(int) (*entity.Service, error)
	GetByCode(string) (*entity.Service, error)
}

func NewServiceRepository(db *sql.DB) *ServiceRepository {
	return &ServiceRepository{
		db: db,
	}
}

func (r *ServiceRepository) Count(category string, pkg string) (int, error) {
	var count int
	err := r.db.QueryRow(queryCountService, category, pkg).Scan(&count)
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *ServiceRepository) GetById(id int) (*entity.Service, error) {
	var s entity.Service
	err := r.db.QueryRow(querySelectService, id).Scan(&s.ID, &s.Category, &s.Code, &s.Name, &s.Package, &s.AuthUser, &s.AuthPass, &s.PartnerID, &s.ProductID, &s.TransID, &s.CorrelationID, &s.Price, &s.RenewalDay, &s.TrialDay, &s.UrlHost, &s.UrlCallback, &s.UrlNotifSub, &s.UrlNotifUnsub, &s.UrlNotifRenewal, &s.UrlPostback)
	if err != nil {
		return &s, err
	}
	return &s, nil
}

func (r *ServiceRepository) GetByCode(code string) (*entity.Service, error) {
	var s entity.Service
	err := r.db.QueryRow(querySelectPackage, code).Scan(&s.ID, &s.Category, &s.Code, &s.Name, &s.Package, &s.AuthUser, &s.AuthPass, &s.PartnerID, &s.ProductID, &s.TransID, &s.CorrelationID, &s.Price, &s.RenewalDay, &s.TrialDay, &s.UrlHost, &s.UrlCallback, &s.UrlNotifSub, &s.UrlNotifUnsub, &s.UrlNotifRenewal, &s.UrlPostback)
	if err != nil {
		return &s, err
	}
	return &s, nil
}
