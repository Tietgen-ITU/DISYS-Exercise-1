package repository

import (
	"log"

	"github.com/ArneProductions/DISYS-exercise-1/models"
	"gorm.io/gorm"
)

type CourseRepository interface {
	CreateCourse(course models.Course) error
	DeleteCourse(courseId int) error
	AddStudent(courseId int, studentId int) error
	RemoveStudent(courseId int, studentId int) error
	GetCourses() ([]models.Course, error)
	GetCourse(courseId int) (models.Course, error)
}

func migrate(db *gorm.DB) error {
	log.Println("{SQLITE COURSE REPOSITORY} Create")

	err := db.AutoMigrate(&models.Course{})

	return err
}

func NewSqliteCourseRepository(db *gorm.DB) CourseRepository {
	repo := sqliteCourseRepository{
		Db: db,
	}

	err := migrate(repo.Db)
	if err != nil {
		log.Fatal("Repo migration failed", err)
	}

	return repo
}

type sqliteCourseRepository struct {
	Db *gorm.DB
}

func (c sqliteCourseRepository) CreateCourse(course models.Course) error {
	return nil
}

func (c sqliteCourseRepository) DeleteCourse(courseId int) error {
	return nil
}

func (c sqliteCourseRepository) AddStudent(courseId int, studentId int) error {
	return nil
}

func (c sqliteCourseRepository) RemoveStudent(courseId int, studentId int) error {
	return nil
}

func (c sqliteCourseRepository) GetCourses() ([]models.Course, error) {
	return nil, nil
}

func (c sqliteCourseRepository) GetCourse(courseId int) (models.Course, error) {
	return models.Course{}, nil
}
