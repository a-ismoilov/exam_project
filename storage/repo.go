package storage

import "github.com/Abdur-Rohman/exam_project/model"

type Storage interface {
	ReadWords(page int, limit int) ([]map[string]int, error)
	WriteWords(body model.Items) error
}
