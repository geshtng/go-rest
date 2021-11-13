package article

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func InitArticle(db *sql.DB, echo *echo.Echo) {
	handler := InitArticleHandler(db)

	echo.GET("/articles", handler.FetchArticles)
	echo.GET("/article/:id", handler.GetArticleByID)

	echo.POST("/article", handler.CreateArticle)

	echo.PUT("/article/:id", handler.UpdateArticle)

	echo.DELETE("/article/:id", handler.DeleteArticle)
}
