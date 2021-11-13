package marticle

import (
	"database/sql"

	"github.com/geshtng/go-rest/models"
	rarticle "github.com/geshtng/go-rest/repository/article"
)

func GetAllArticles(db *sql.DB) ([]models.Article, error) {
	result, err := rarticle.GetAllArticles(db)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func CreateArticle(db *sql.DB, article models.Article) (*models.Article, error) {
	result, err := rarticle.InsertArticle(db, article)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetArticleByID(db *sql.DB, articleID int64) (*models.Article, error) {
	result, err := rarticle.GetArticleByID(db, articleID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func UpdateArticle(db *sql.DB, articleID int64, article models.Article) (*models.Article, error) {
	result, err := rarticle.UpdateArticle(db, articleID, article)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func DeleteArticle(db *sql.DB, articleID int64) error {
	err := rarticle.DeleteArticle(db, articleID)
	if err != nil {
		return err
	}

	return nil
}
