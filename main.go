package main

import (
	"database/sql"
	"log"

	"github.com/geshtng/go-rest/config"
	"github.com/geshtng/go-rest/handler"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	conf := config.InitConfig()

	db, err := sql.Open("postgres", conf)
	if err != nil {
		log.Fatal(err)
	}

	echoServer := echo.New()

	handler.InitHandler(db, echoServer)

	echoServer.Start(":8080")
}
