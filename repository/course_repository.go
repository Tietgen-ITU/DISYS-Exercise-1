package repository

import (
	"log"

	"github.com/ArneProductions/DISYS-exercise-1/models"
	"gorm.io/gorm"
)

type CourseRepository interface {
	CreateCourse(course models.Course) error
	DeleteCourse(courseId uint64) error
	AddStudent(courseId uint64, studentId uint64) error
	RemoveStudent(courseId uint64, studentId uint64) error
	GetCourses() ([]models.Course, error)
	GetCourse(courseId uint64) (models.Course, error)
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

	return c.Db.Create(&course).Error
}

func (c sqliteCourseRepository) DeleteCourse(courseId uint64) error {
	return c.Db.Delete(&models.Course{Id: courseId}).Error
}

func (c sqliteCourseRepository) AddStudent(courseId uint64, studentId uint64) error {

	return c.Db.Model(&models.Course{Id: courseId}).Association("Students").Append(models.User{ID: studentId})
}

func (c sqliteCourseRepository) RemoveStudent(courseId uint64, studentId uint64) error {
	return c.Db.Model(&models.Course{Id: courseId}).Association("Students").Delete(models.User{ID: studentId})
}

func (c sqliteCourseRepository) GetCourses() (courses []models.Course, err error) {

	err = c.Db.Find(&courses).Error
	return courses, err
}

func (c sqliteCourseRepository) GetCourse(courseId uint64) (course models.Course, err error) {
	err = c.Db.First(course, models.Course{Id: courseId}).Error
	return models.Course{}, err
}
