# Wallet App
This project is a simple web application (API) that simulates a wallet app that allows user to send money to one another

## Getting Started

### Prerequisites
---
* Install go binaries on your machine [Golang](https://golang.org/doc/install)

### Start up
To start this application, perform the following step in the order

* Clone this repo to your machine
* cd into the project folder and make sure you are on the root directory.
* In your terminal, enter `cp .env.example .env` , this will create .env file for you
* Fill in your details to the .env, most expecially the *Email Setup* part, so you'll be able to send email to reset password
* In your terminal, enter `go run main.go` to start server

## Built With
---
* [Golang](https://golang.org) - Language used
* [Gorilla Mux](https://github.com/gorilla/mux) - HTTP Router used
* [GORM](https://github.com/go-gorm/gorm) - ORM library used
* [MySQL](https://www.mysql.com) - Database used

## Documentation
---
Documentation for the Project can be found at [wallet-app](https://documenter.getpostman.com/view/947913/SzzehgDD)
