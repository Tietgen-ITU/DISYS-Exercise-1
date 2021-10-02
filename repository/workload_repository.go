package repository

import (
	"github.com/ArneProductions/DISYS-exercise-1/models"
	"gorm.io/gorm"
	"log"
)

type WorkloadRepository interface {
	Create(models.Workload) (models.Workload, error)
	Migrate() error
}

type sqliteWorkloadRepository struct {
	Db *gorm.DB
}

func NewSqliteWorkloadRepository(db *gorm.DB) WorkloadRepository {
	repo := sqliteWorkloadRepository{
		Db: db,
	}

	err := repo.Migrate()
	if err != nil {
		log.Fatal("Repo migration failed", err)
	}

	return repo
}


func (s sqliteWorkloadRepository) Create(workload models.Workload) (models.Workload, error) {
	log.Println("{SQLITE WORKLOAD REPOSITORY} Create")

	err := s.Db.Create(&workload).Error

	return workload, err
}


func (s sqliteWorkloadRepository) Migrate() error {
	log.Println("{SQLITE WORKLOAD REPOSITORY} Create")

	err := s.Db.AutoMigrate(&models.Workload{})

	return err
}
