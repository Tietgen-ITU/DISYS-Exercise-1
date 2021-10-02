package models

type User struct {
	ID uint64 `json:"id,omitempty" gorm:"primaryKey;autoIncrement"`
}

func IntToUser(id uint64) User {
	return User{id}
}

func IntArrayToUser(ids []uint64) []User {
	models := make([]User, len(ids))

	for i, value := range ids {
		models[i] = IntToUser(value)
	}

	return models
}

func UserToInt(user User) uint64 {
	return user.ID
}

func UserArrayToInt(users []User) []uint64 {
	models := make([]uint64, len(users))

	for i, value := range users {
		models[i] = UserToInt(value)
	}

	return models
}
