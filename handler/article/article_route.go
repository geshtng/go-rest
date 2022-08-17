package article

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func InitArticle(db *sql.DB, echo *echo.Echo) {
	handler := InitArticleHandler(db)

	echo.GET("/articles", handler.FetchArticles)
	echo.GET("/articles/:id", handler.GetArticleByID)

	echo.POST("/articles", handler.CreateArticle)

	echo.PUT("/articles/:id", handler.UpdateArticle)

	echo.DELETE("/articles/:id", handler.DeleteArticle)
}
