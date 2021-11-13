package article

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/geshtng/go-rest/models"
	marticle "github.com/geshtng/go-rest/module/article"
)

type ErrResponse struct {
	Message string `json:"message"`
}

type ArticleHandler struct {
	DB *sql.DB
}

func InitArticleHandler(db *sql.DB) ArticleHandler {
	return ArticleHandler{
		DB: db,
	}
}

func (h ArticleHandler) FetchArticles(c echo.Context) error {
	articles, err := marticle.GetAllArticles(h.DB)
	if err != nil {
		resp := ErrResponse{
			Message: err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, resp)
	}

	return c.JSON(http.StatusOK, articles)
}

func (h ArticleHandler) CreateArticle(c echo.Context) error {
	var article models.Article

	err := c.Bind(&article)
	if err != nil {
		resp := ErrResponse{
			Message: err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, resp)
	}

	createdArticle, err := marticle.CreateArticle(h.DB, article)
	if err != nil {
		resp := ErrResponse{
			Message: err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, resp)
	}

	return c.JSON(http.StatusCreated, createdArticle)
}

func (h ArticleHandler) GetArticleByID(c echo.Context) error {
	articleID, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	article, err := marticle.GetArticleByID(h.DB, articleID)
	if err != nil {
		resp := ErrResponse{
			Message: err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, resp)
	}

	return c.JSON(http.StatusOK, article)
}

func (h ArticleHandler) UpdateArticle(c echo.Context) error {
	var article models.Article

	articleID, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	err := c.Bind(&article)
	if err != nil {
		resp := ErrResponse{
			Message: err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, resp)
	}

	updatedArticle, err := marticle.UpdateArticle(h.DB, articleID, article)
	if err != nil {
		resp := ErrResponse{
			Message: err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, resp)
	}

	return c.JSON(http.StatusOK, updatedArticle)
}

func (h ArticleHandler) DeleteArticle(c echo.Context) error {
	articleID, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	resp := ErrResponse{}

	err := marticle.DeleteArticle(h.DB, articleID)
	if err != nil {
		resp.Message = err.Error()

		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Message = "article deleted"

	return c.JSON(http.StatusOK, resp)
}
