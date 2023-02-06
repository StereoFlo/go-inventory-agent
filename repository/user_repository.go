package repository

import (
	"go-agent/entity"
	"os/user"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur UserRepository) GetUser() (*entity.User, error) {
	current, err := user.Current()
	if err != nil {
		return nil, err
	}
	userEntity := &entity.User{
		Uid:      current.Uid,
		Gid:      current.Gid,
		Username: current.Username,
		Name:     current.Name,
		HomeDir:  current.HomeDir,
	}

	return userEntity, nil
}
