package main

import (
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"github.com/geshtng/go-rest/config"
	"github.com/geshtng/go-rest/handler"
)

func main() {
	conf := config.InitConfigDsn()
	port := config.InitConfigServer()

	db, err := sql.Open("postgres", conf)
	if err != nil {
		log.Fatal(err)
	}

	echoServer := echo.New()

	handler.InitHandler(db, echoServer)

	echoServer.Start(port)
}
