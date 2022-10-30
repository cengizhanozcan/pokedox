package service

import (
	"pokemon/middleware/database"
	"pokemon/model/domain"
)

type ITrainerService interface {
	GetTrainer(ID uint64) *domain.Trainer
	GetTrainerByNickname(nickname string) *domain.Trainer
	CreateTrainer(trainer *domain.Trainer)
}

type trainerService struct {
	db *database.Database
}

func (t *trainerService) GetTrainer(ID uint64) *domain.Trainer {
	trainer := domain.Trainer{}
	t.db.Instance.First(&trainer, "id=?", ID)
	if trainer.ID != 0 {
		return &trainer
	}
	return nil
}

func (t *trainerService) GetTrainerByNickname(nickname string) *domain.Trainer {
	trainer := domain.Trainer{}
	t.db.Instance.First(&trainer, "nickname=?", nickname)
	if trainer.ID != 0 {
		return &trainer
	}
	return nil
}

func (t *trainerService) CreateTrainer(trainer *domain.Trainer) {
	t.db.Instance.Create(trainer)
}

func NewTrainerService(db *database.Database) ITrainerService {
	return &trainerService{db: db}
}
