package persistence

import (
	"context"

	"github.com/t-tiger/survey/server/cerrors"
	"github.com/t-tiger/survey/server/entity"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{db: db}
}

func (p *User) FindBy(ctx context.Context, id string) (entity.User, error) {
	var u entity.User
	if err := p.db.WithContext(ctx).Where(&entity.User{ID: id}).First(&u).Error; err != nil {
		return entity.User{}, err
	}
	return u, nil
}

func (p *User) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var u entity.User
	if err := p.db.WithContext(ctx).Where(&entity.User{Email: email}).First(&u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, cerrors.Errorf(cerrors.DatabaseErr, err.Error())
	}
	return &u, nil
}

func (p *User) Create(ctx context.Context, name, email, password string) (entity.User, error) {
	u := entity.User{Name: name, Email: email, PasswordDigest: password}
	if err := p.db.WithContext(ctx).Table("users").Create(&u).Error; err != nil {
		return entity.User{}, cerrors.Errorf(cerrors.DatabaseErr, err.Error())
	}
	return u, nil
}
