package handler

import (
	"database/sql"

	"github.com/labstack/echo/v4"

	"github.com/geshtng/go-rest/handler/article"
)

func InitHandler(db *sql.DB, echo *echo.Echo) {
	article.InitArticle(db, echo)
}
