package repository

import (
	"github.com/ArneProductions/DISYS-exercise-1/models"
	"gorm.io/gorm"
	"log"
)

type StudentWorkloadRepository interface {
	Create(models.StudentWorkload) (models.StudentWorkload, error)
	GetByIds(uint64, uint64) (models.StudentWorkload, error)
	Migrate() error
}

type sqliteStudentWorkloadRepository struct {
	Db *gorm.DB
}

func NewSqliteStudentWorkloadRepository(db *gorm.DB) StudentWorkloadRepository {
	repo := sqliteStudentWorkloadRepository{
		Db: db,
	}

	err := repo.Migrate()
	if err != nil {
		log.Fatal("Repo migration failed", err)
	}

	return repo
}


func (s sqliteStudentWorkloadRepository) Create(studentworkload models.StudentWorkload) (models.StudentWorkload, error) {
	log.Println("{SQLITE STUDENTWORKLOAD REPOSITORY} Create")

	err := s.Db.Create(&studentworkload).Error

	return studentworkload, err
}

func (s sqliteStudentWorkloadRepository) GetByIds(studentId uint64, workladId uint64) (workload models.StudentWorkload, err error) {
	log.Println("{SQLITE WORKLOAD REPOSITORY} GetByIds")

	err = s.Db.First(&workload, studentId, workladId).Error

	return workload, err
}


func (s sqliteStudentWorkloadRepository) Migrate() error {
	log.Println("{SQLITE STUDENTWORKLOAD REPOSITORY} Create")

	err := s.Db.AutoMigrate(&models.StudentWorkload{})

	return err
}
