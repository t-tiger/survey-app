package entity

type User struct {
	ID             string `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	PasswordDigest string `json:"-"`
}
