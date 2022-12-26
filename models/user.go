package models

type User struct {
	ID    uint   `json:"user_id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Posts []Post `json:"posts" gorm:"foreignKey:UserId"`
}
