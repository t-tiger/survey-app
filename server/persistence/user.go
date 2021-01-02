package persistence

import (
	"github.com/t-tiger/survey/server/entity"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{db: db}
}

func (p *User) Create(email, password string) (entity.User, error) {
	panic("implement me")
}
