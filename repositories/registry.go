package repositories

import (
	"gorm.io/gorm"
	repositories "payment-service/repositories/payment"
	repositories2 "payment-service/repositories/paymenthistory"
)

type Registry struct {
	db *gorm.DB
}

type IRepositoryRegistry interface {
	GetPayment() repositories.IPaymentRepository
	GetPaymentHistory() repositories2.IPaymentHistoryRepository
	GetTx() *gorm.DB
}

func NewRepositoryRegistry(db *gorm.DB) IRepositoryRegistry {
	return &Registry{db: db}
}

func (r *Registry) GetPayment() repositories.IPaymentRepository {
	return repositories.NewPaymentRepository(r.db)
}

func (r *Registry) GetPaymentHistory() repositories2.IPaymentHistoryRepository {
	return repositories2.NewPaymentHistoryRepository(r.db)
}

func (r *Registry) GetTx() *gorm.DB {
	return r.db
}
