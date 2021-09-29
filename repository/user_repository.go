package repository

import (
	"github.com/ArneProductions/DISYS-exercise-1/models"
	"gorm.io/gorm"
	"log"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
	GetById(uint64) (models.User, error)
	Create(models.User) (models.User, error)
	Update(models.User) (models.User, error)
	Delete(uint64) error
	Migrate() error
}

type sqliteUserRepository struct {
	Db *gorm.DB
}

func NewSqliteUserRepository(db *gorm.DB) UserRepository {
	repo := sqliteUserRepository{
		Db: db,
	}

	err := repo.Migrate()
	if err != nil {
		log.Fatal("Repo migration failed", err)
	}

	return repo
}

func (s sqliteUserRepository) GetAll() (users []models.User, err error) {
	log.Println("{SQLITE USER REPOSITORY} GetAll")

	err = s.Db.Find(&users).Error

	return users, err
}

func (s sqliteUserRepository) GetById(userId uint64) (user models.User, err error) {
	log.Println("{SQLITE USER REPOSITORY} GetById")

	err = s.Db.First(&user, userId).Error

	return user, err
}

func (s sqliteUserRepository) Create(user models.User) (models.User, error) {
	log.Println("{SQLITE USER REPOSITORY} Create")

	err := s.Db.Create(&user).Error

	return user, err
}

func (s sqliteUserRepository) Update(user models.User) (models.User, error) {
	log.Println("{SQLITE USER REPOSITORY} Update")

	err := s.Db.Save(user).Error

	return user, err
}

func (s sqliteUserRepository) Delete(id uint64) error {
	log.Println("{SQLITE USER REPOSITORY} Delete")

	err := s.Db.Delete(&models.User{}, id).Error

	return err
}

func (s sqliteUserRepository) Migrate() error {
	log.Println("{SQLITE USER REPOSITORY} Create")

	err := s.Db.AutoMigrate(&models.User{})

	return err
}
