package main

import (
	"fmt"
	"walletserver/api"
	"walletserver/app"
)

func main() {
	engine := app.NewGinEngine()
	router := api.NewRouter()
	server := app.NewServer(engine, router)
	fmt.Println("swagger doc: http://127.0.0.1:8089/swagger/index.html")
	server.StartWallet()
}
