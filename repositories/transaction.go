package repositories

// import (
// 	"housy/models"

// 	"gorm.io/gorm"
// )

// type TransactionRepository interface {
// 	FindTransactions() ([]models.Transaction, error)
// 	GetTransaction(ID int) (models.Transaction, error)
// 	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
// }

// func RepositoryTransaction(db *gorm.DB) *repository {
// 	return &repository{db}
// }

// func (r *repository) FindTransactions() ([]models.Transaction, error) {
// 	var transactions []models.Transaction
// 	err := r.db.Preload("House.City").Preload("User").Find(&transactions).Error

// 	return transactions, err
// }

// func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
// 	var transaction models.Transaction
// 	err := r.db.Preload("House.City").Preload("User").First(&transaction, ID).Error

// 	return transaction, err
// }

// func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
// 	err := r.db.Preload("House.City").Preload("User").Create(&transaction).Error

// 	return transaction, err
// }
