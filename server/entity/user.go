package entity

type User struct {
	ID    string `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name  string
	Email string
}
