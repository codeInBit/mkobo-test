package controllers

import (
	"github.com/codeInBit/mkobo-test/middlewares"
)

//LoadRoutes - This is where all the routes are defined
func (s *Server) LoadRoutes() {
	api := s.Router.PathPrefix("/api").Subrouter()
	//Home Route
	api.HandleFunc("", middlewares.CommonMiddleware(s.Home)).Methods("GET")

	//User routes
	api.HandleFunc("/users/register", middlewares.CommonMiddleware(s.Register)).Methods("POST")
	api.HandleFunc("/users/login", middlewares.CommonMiddleware(s.Login)).Methods("POST")
	api.HandleFunc("/users/forgot-password", middlewares.CommonMiddleware(s.ForgotPassword)).Methods("POST")
	api.HandleFunc("/reset-password/{token}", middlewares.CommonMiddleware(s.ResetPassword)).Methods("POST")

	api.HandleFunc("/users/wallets/transfer", middlewares.CommonMiddleware(middlewares.AuthenticationMiddleware(s.Transfer))).Methods("POST")
	api.HandleFunc("/users/wallets/transaction-history", middlewares.CommonMiddleware(middlewares.AuthenticationMiddleware(s.UserWalletTransactionHistory))).Methods("GET")

	//Wallet
	api.HandleFunc("/wallets/transaction-history", middlewares.CommonMiddleware(s.AllWalletTransactionHistory)).Methods("GET")

}
