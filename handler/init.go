package handler

import (
	"database/sql"

	"github.com/geshtng/go-rest/handler/article"
	"github.com/labstack/echo/v4"
)

func InitHandler(db *sql.DB, echo *echo.Echo) {
	article.InitArticle(db, echo)
}
