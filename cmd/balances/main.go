package main

import (
	"database/sql"
	"fmt"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/database"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/get_balance"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/web"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/web/webserver"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// configMap := ckafka.ConfigMap{
	// 	"bootstrap.servers": "kafka:29092",
	// 	"group.id":          "wallet",
	// }
	// kafkaProducer := kafka.NewKafkaProducer(&configMap)
	balanceDb := database.NewBalancesDB(db)
	getBalanceUseCase := get_balance.NewGetBalanceUseCase(balanceDb)

	webserver := webserver.NewWebServer(":3003")

	balanceHandler := web.NewWebBalanceHandler(*getBalanceUseCase)

	webserver.AddHandler("/balances/{account_id}", balanceHandler.GetBalance)

	fmt.Println("Server is running")
	webserver.Start()
}
