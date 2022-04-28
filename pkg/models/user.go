package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/kapitanmx/ATM/pkg/config"
)

var db *gorm.DB

type User struct {
	gorm.Model
	ID           string        `gorm:""json:"ID"`
	Name         string        `json:"name"`
	LastName     string        `json:"lastname"`
	Email        string        `json:"email"`
	Pin          string        `json:"pin"`
	Transactions []Transaction `json:"transactions"`
	Sum          int64
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) SetID() error {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	u.ID = id.String()
	return nil
}

func (u *User) CreateUser() *User {
	db.NewRecord(u)
	db.Create(&u)
	return u
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserById(ID string) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("ID=?", ID).Find(&getUser)
	return &getUser, db
}

func (u *User) GetUserTransactions() []Transaction {
	return u.Transactions
}

func DeleteUser(ID string) User {
	var user User
	db.Where("ID=?", ID).Delete(user)
	return user
}
