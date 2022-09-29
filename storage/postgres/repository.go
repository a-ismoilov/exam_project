package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Abdur-Rohman/exam_project/model"
	"go.uber.org/zap"
	"log"
)

type Repository struct {
	logger *zap.Logger
	db     *sql.DB
}

func New(l *zap.Logger) *Repository {
	return &Repository{
		logger: l,
		db:     DB(),
	}
}

func (r *Repository) ReadWords(page int, limit int) ([]map[string]int, error) {
	var (
		key string
		val int
	)
	items := make([]map[string]int, 0)
	item := map[string]int{}
	offset := (page - 1) * limit
	query := `SELECT * FROM items LIMIT $1 OFFSET $2`
	log.Println(limit, page)
	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		r.logger.Error(fmt.Sprintf("error while reading from database, err: %v", err))
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(&key, &val); err != nil {
			return nil, err
		}

		item[key] = val
		items = append(items, item)
	}
	return items, nil
}
func (r *Repository) WriteWords(body model.Items) error {
	query := `INSERT INTO items VALUES ($1, $2)`

	for _, val := range body.Words {
		for k, v := range val {
			if _, err := r.db.Exec(query, k, v); err != nil {
				return err
			}
		}
	}
	return nil
}
