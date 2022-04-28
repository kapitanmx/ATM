package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/kapitanmx/ATM/pkg/config"
)

var db *gorm.DB

type Transaction struct {
	gorm.Model
	ID           string
	Date         string
	Title        string
	Description  string
	SenderName   string
	SenderID     string
	ReceiverName string
	ReceiverID   string
	Sum          int64
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Transaction{})
}

func (t *Transaction) SetID() error {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	t.ID = id.String()
	return nil
}

func (t *Transaction) CreateTransaction() *Transaction {
	db.NewRecord(t)
	db.Create(&t)
	return t
}

func (t *Transaction) GetAllTransactions() []Transaction {
	var Transactions []Transaction
	db.Find(&Transactions)
	return Transactions
}

func (t *Transaction) GetTransactionById(id string) (*Transaction, *gorm.DB) {
	var getTransaction Transaction
	t.ID = id
	ID := t.ID
	db := db.Where("ID=?", ID).Find(&getTransaction)
	return &getTransaction, db
}

func (t *Transaction) DeleteTransaction() Transaction {
	var transaction Transaction
	ID := t.ID
	db.Where("ID=?", ID).Delete(transaction)
	return transaction
}
