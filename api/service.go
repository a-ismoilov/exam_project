package api

import "github.com/Abdur-Rohman/exam_project/model"

type IService interface {
	WriteWords(items model.Items) error
	ReadWords(page int, limit int) ([]map[string]int, error)
}
