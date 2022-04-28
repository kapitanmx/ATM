package routes

import (
	"github.com/gorilla/mux"
	"github.com/kapitanmx/ATM/pkg/controllers"
)

var Routes = func(router *mux.Router) {
	// User routes
	router.HandleFunc("/user/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{userId}", controllers.DeleteUser).Methods("DELETE")
	// Transaction routes
	router.HandleFunc("/transaction/", controllers.CreateTransaction).Methods("POST")
	router.HandleFunc("/transaction/", controllers.GetTransaction).Methods("GET")
	router.HandleFunc("/transaction/{transactionId}", controllers.GetTransactionById).Methods("GET")
	router.HandleFunc("/transaction/{transactionId}", controllers.UpdateTransaction).Methods("PUT")
	router.HandleFunc("/transaction/{transactionId}", controllers.DeleteTransaction).Methods("DELETE")

}
