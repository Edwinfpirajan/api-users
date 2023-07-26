package user

import "github.com/Edwinfpirajan/curso-hex-arqu.git/model"

type UseCase interface {
	Create(m *model.User) error
}
