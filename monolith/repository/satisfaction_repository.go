package repository

import (
	"log"

	"github.com/ArneProductions/DISYS-exercise-1/models"
	"gorm.io/gorm"
)

type CourseSatisfaction struct {
	AvgSatisfaction float32
}


type SatisfactionRepository interface {
	
	GetCourseSatisfactionById(uint64) (models.StudentSatisfaction, error)

	GetStudentSatisfactionById(uint64) (models.StudentSatisfaction, error)

	Create(models.StudentSatisfaction) (models.StudentSatisfaction, error)

	Migrate() error
}

type sqliteSatisfactionRepository struct {
	Db *gorm.DB
}

func NewSqliteSatisfactionRepository(db *gorm.DB) SatisfactionRepository {
	repo := sqliteSatisfactionRepository{
		Db: db,
	}

	err := repo.Migrate()
	if err != nil {
		log.Fatal("Repo migration failed", err)
	}

	return repo
}

func (s sqliteSatisfactionRepository) GetCourseSatisfactionById(course_id uint64) (satisfaction models.StudentSatisfaction , err error) {
	log.Println("{SQLITE Satisfaction REPOSITORY} GetCourseSatisfactionById")

	err = s.Db.Select("AVG(satisfaction) as Satisfaction").Group("course_id").Where("course_id = (?)", course_id).Find(&satisfaction).Error

	return satisfaction, err
}

func (s sqliteSatisfactionRepository) GetStudentSatisfactionById(student_id uint64) (satisfaction models.StudentSatisfaction, err error) {
	log.Println("{SQLITE Satisfaction REPOSITORY} GetStudentSatisfactionById")
	satisfaction.StudentId = int64(student_id)
	err = s.Db.First(&satisfaction, satisfaction).Error

	return satisfaction, err
}

func (s sqliteSatisfactionRepository) Create(satisfaction models.StudentSatisfaction) (models.StudentSatisfaction, error) {
	log.Println("{SQLITE SATISFACTION REPOSITORY} Create")

	err := s.Db.Create(&satisfaction).Error

	return satisfaction, err
}

func (s sqliteSatisfactionRepository) Migrate() error {
	log.Println("{SQLITE SATISFACTON REPOSITORY} Create")

	err := s.Db.AutoMigrate(&models.StudentSatisfaction{})

	return err
}