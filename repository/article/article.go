package rarticle

import (
	"database/sql"
	"fmt"

	"github.com/geshtng/go-rest/models"
)

func GetAllArticles(db *sql.DB) ([]models.Article, error) {
	query := `
	SELECT id, title, body 
	FROM article`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	articles := []models.Article{}

	for rows.Next() {
		var article models.Article

		err := rows.Scan(&article.ID, &article.Title, &article.Body)
		if err != nil {
			return nil, err
		}

		articles = append(articles, article)
	}

	defer rows.Close()

	return articles, nil
}

func InsertArticle(db *sql.DB, article models.Article) (*models.Article, error) {
	query := `
	INSERT INTO article (title, body)
	VALUES ($1, $2)
	RETURNING id`

	var id int64

	err := db.QueryRow(query, article.Title, article.Body).Scan(&id)
	if err != nil {
		return nil, err
	}

	article.ID = id

	return &article, nil
}

func GetArticleByID(db *sql.DB, articleID int64) (*models.Article, error) {
	query := `
	SELECT id, title, body
	FROM article
	WHERE id=$1`

	article := models.Article{}

	err := db.QueryRow(query, articleID).Scan(&article.ID, &article.Title, &article.Body)
	if err != nil {
		return nil, err
	}

	return &article, nil
}

func UpdateArticle(db *sql.DB, articleID int64, article models.Article) (*models.Article, error) {
	query := `
	UPDATE article
	SET title=$2, body=$3
	WHERE id=$1`

	dbRes, err := db.Exec(query, articleID, article.Title, article.Body)
	if err != nil {
		return nil, err
	}

	totalUpdated, err := dbRes.RowsAffected()
	if err != nil {
		return nil, err
	}

	if totalUpdated <= 0 {
		errs := fmt.Errorf("nothing is changed")

		return nil, errs
	}

	article.ID = articleID

	return &article, nil
}

func DeleteArticle(db *sql.DB, articleID int64) error {
	query := `
	DELETE FROM article
	WHERE id=$1`

	_, err := db.Exec(query, articleID)
	if err != nil {
		return err
	}

	return nil
}
