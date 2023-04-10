package services

import "github.com/idprm/go-telenor-pk/src/domain/repository"

type TransactionService struct {
	transactionRepo repository.ITransactionRepository
}

type ITransactionService interface {
	SaveTransaction()
	UpdateTransaction()
}

func NewTransactionService(transactionRepo repository.ITransactionRepository) *TransactionService {
	return &TransactionService{
		transactionRepo: transactionRepo,
	}
}

func (s *TransactionService) SaveTransaction() {

}

func (s *TransactionService) UpdateTransaction() {

}
