package repository

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/idprm/go-telenor-pk/src/domain/entity"
)

const (
	queryInsertSubscription      = "INSERT INTO subscriptions(service_id, msisdn, channel, adnet, pub_id, aff_sub, latest_subject, latest_status, success, ip_address, is_active, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)"
	queryUpdateSubscription      = "UPDATE subscriptions SET channel = $1, adnet = $2, pub_id = $3, aff_sub = $4, latest_subject = $5, latest_status = $6, amount = amount + $7, trial_at = $8, renewal_at = $9, unsub_at = $10, charge_at = $11, retry_at = $12, success = success + $13, ip_address = $14, is_trial = $15, is_retry = $16, is_active = $17, total_firstpush = total_firstpush + $18, total_renewal = total_renewal + $19, total_amount_firstpush = total_amount_firstpush + $20, total_amount_renewal = total_amount_renewal + $21, updated_at = NOW() WHERE service_id = $22 AND msisdn = $23"
	queryCountSubscription       = "SELECT COUNT(*) as count FROM subscriptions WHERE service_id = $1 AND msisdn = $2"
	queryCountActiveSubscription = "SELECT COUNT(*) as count FROM subscriptions WHERE service_id = $1 AND msisdn = $2 AND is_active = true"
	querySelectSubscription      = "SELECT id, service_id, msisdn, channel, adnet, pub_id, aff_sub, latest_subject, latest_status, amount, trial_at, renewal_at, unsub_at, charge_at, retry_at, success, ip_address, total_firstpush, total_renewal, total_amount_firstpush, total_amount_renewal, is_trial, is_retry, is_active FROM subscriptions WHERE service_id = $1 AND msisdn = $2"
	querySelectPopulateRenewal   = "SELECT id, msisdn, service_id, channel, ip_address FROM subscriptions WHERE renewal_at IS NOT NULL AND DATE(renewal_at) <= DATE(NOW()) AND is_active = true AND deleted_at IS null ORDER BY success DESC"
	querySelectPopulateRetry     = "SELECT id, msisdn, service_id, channel, ip_address FROM subscriptions WHERE renewal_at IS NOT NULL AND DATE(renewal_at) = DATE(NOW() + interval '1 day') AND is_retry = true AND is_active = true AND deleted_at IS null ORDER BY success DESC"
	querySelectPopulateReminder  = "SELECT id, msisdn, service_id, channel, ip_address FROM subscriptions WHERE renewal_at IS NOT NULL AND DATE(renewal_at) = DATE(NOW() + interval '2 day') AND is_retry = false AND is_active = true AND deleted_at IS null ORDER BY success DESC"
)

type SubscriptionRepository struct {
	db *sql.DB
}

type ISubscriptionRepository interface {
	Save(*entity.Subscription) error
	Update(*entity.Subscription) error
	Count(int, string) (int, error)
	CountActive(int, string) (int, error)
	Get(int, string) (*entity.Subscription, error)
	Renewal() (*[]entity.Subscription, error)
	Retry() (*[]entity.Subscription, error)
	Reminder() (*[]entity.Subscription, error)
}

func NewSubscriptionRepository(db *sql.DB) *SubscriptionRepository {
	return &SubscriptionRepository{
		db: db,
	}
}

func (r *SubscriptionRepository) Save(s *entity.Subscription) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := r.db.PrepareContext(ctx, queryInsertSubscription)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, s.ServiceID, s.Msisdn, s.Channel, s.Adnet, s.PubID, s.AffSub, s.LatestSubject, s.LatestStatus, s.Success, s.IpAddress, s.IsActive, time.Now(), time.Now())
	if err != nil {
		log.Printf("Error %s when inserting row into subscriptions table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d subscriptions created ", rows)
	return nil
}

func (r *SubscriptionRepository) Update(s *entity.Subscription) error {

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := r.db.PrepareContext(ctx, queryUpdateSubscription)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, s.Channel, s.Adnet, s.PubID, s.AffSub, s.LatestSubject, s.LatestStatus, s.Amount, s.TrialAt, s.RenewalAt, s.UnsubAt, s.ChargeAt, s.RetryAt, s.Success, s.IpAddress, s.IsTrial, s.IsRetry, s.IsActive, s.TotalFirstpush, s.TotalRenewal, s.TotalAmountFirstpush, s.TotalAmountRenewal, s.ServiceID, s.Msisdn)
	if err != nil {
		log.Printf("Error %s when update row into subscriptions table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d subscriptions updated ", rows)

	return nil
}

func (r *SubscriptionRepository) Count(serviceId int, msisdn string) (int, error) {
	var count int
	err := r.db.QueryRow(queryCountSubscription, serviceId, msisdn).Scan(&count)
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionRepository) CountActive(serviceId int, msisdn string) (int, error) {
	var count int
	err := r.db.QueryRow(queryCountSubscription, serviceId, msisdn).Scan(&count)
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionRepository) Get(serviceId int, msisdn string) (*entity.Subscription, error) {
	var s entity.Subscription
	err := r.db.QueryRow(querySelectPackage, serviceId, msisdn).Scan(&s.ID, &s.ServiceID, &s.Msisdn, &s.Channel, &s.Adnet, &s.PubID, &s.AffSub, &s.LatestSubject, &s.LatestStatus, &s.Amount, &s.TrialAt, &s.RenewalAt, &s.UnsubAt, &s.ChargeAt, &s.RetryAt, &s.Success, &s.IpAddress, &s.TotalFirstpush, &s.TotalRenewal, &s.TotalAmountFirstpush, &s.TotalAmountRenewal, &s.IsTrial, &s.IsRetry, &s.IsActive)
	if err != nil {
		return &s, err
	}
	return &s, nil
}

func (r *SubscriptionRepository) Renewal() (*[]entity.Subscription, error) {
	rows, err := r.db.Query(querySelectPopulateRenewal)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subs []entity.Subscription

	for rows.Next() {

		var s entity.Subscription
		if err := rows.Scan(&s.ID, &s.Msisdn, &s.ServiceID, &s.Channel, &s.IpAddress); err != nil {
			return nil, err
		}
		subs = append(subs, s)
	}

	if err = rows.Err(); err != nil {
		return &subs, err
	}

	return &subs, nil
}

func (r *SubscriptionRepository) Retry() (*[]entity.Subscription, error) {
	rows, err := r.db.Query(querySelectPopulateRetry)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subs []entity.Subscription

	for rows.Next() {

		var s entity.Subscription
		if err := rows.Scan(&s.ID, &s.Msisdn, &s.ServiceID, &s.Channel, &s.IpAddress); err != nil {
			return nil, err
		}
		subs = append(subs, s)
	}

	if err = rows.Err(); err != nil {
		return &subs, err
	}

	return &subs, nil
}

func (r *SubscriptionRepository) Reminder() (*[]entity.Subscription, error) {
	rows, err := r.db.Query(querySelectPopulateReminder)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subs []entity.Subscription

	for rows.Next() {

		var s entity.Subscription
		if err := rows.Scan(&s.ID, &s.Msisdn, &s.ServiceID, &s.Channel, &s.IpAddress); err != nil {
			return nil, err
		}
		subs = append(subs, s)
	}

	if err = rows.Err(); err != nil {
		return &subs, err
	}

	return &subs, nil
}
