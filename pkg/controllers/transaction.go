package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kapitanmx/ATM/pkg/models"
	"github.com/kapitanmx/ATM/pkg/utils"
)

var NewTransaction models.Transaction

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	newTransactions := models.GetAllTransactions()
	res, _ := json.Marshal(newTransactions)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTransactionById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["transactionId"]
	transactionDetails, _ := models.GetTransactionById(ID)
	res, _ := json.Marshal(transactionDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	CreateTransaction := &models.Transaction{}
	utils.ParseBody(r, CreateTransaction)
	b := CreateTransaction.CreateTransaction()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["transactionId"]
	transaction := models.DeleteTransaction(ID)
	res, _ := json.Marshal(transaction)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	var updateTransaction = &models.Transaction{}
	utils.ParseBody(r, updateTransaction)
	vars := mux.Vars(r)
	ID := vars["transactionId"]
	transactionDetails, db := models.GetTransactionById(ID)
	if updateTransaction.Date != "" {
		transactionDetails.Date = updateTransaction.Date
	}
	if updateTransaction.Title != "" {
		transactionDetails.Title = updateTransaction.Title
	}
	if updateTransaction.Description != "" {
		transactionDetails.Description = updateTransaction.Description
	}
	if updateTransaction.SenderName != "" {
		transactionDetails.SenderName = updateTransaction.SenderName
	}
	if updateTransaction.SenderID != "" {
		transactionDetails.SenderID = updateTransaction.SenderID
	}
	if updateTransaction.ReceiverName != "" {
		transactionDetails.ReceiverName = updateTransaction.ReceiverName
	}
	if updateTransaction.ReceiverID != "" {
		transactionDetails.ReceiverID = updateTransaction.ReceiverID
	}
	if updateTransaction.Sum > 0 {
		transactionDetails.Sum = updateTransaction.Sum
	}
	db.Save(&transactionDetails)
	res, _ := json.Marshal(transactionDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
