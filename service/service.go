package service

import (
	"github.com/Abdur-Rohman/exam_project/model"
	"github.com/Abdur-Rohman/exam_project/storage"
)

type Service struct {
	repo storage.Storage
}

func New(storage storage.Storage) *Service {
	return &Service{
		repo: storage,
	}
}

func (s *Service) WriteWords(items model.Items) error {
	if err := s.repo.WriteWords(items); err != nil {
		return err
	}
	return nil
}

func (s *Service) ReadWords(page int, limit int) ([]map[string]int, error) {
	body, err := s.repo.ReadWords(page, limit)
	if err != nil {
		return nil, err
	}
	return body, nil
}
