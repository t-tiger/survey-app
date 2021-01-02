package repository

import "github.com/t-tiger/survey/server/entity"

type User interface {
	Create(email, password string) (entity.User, error)
}
