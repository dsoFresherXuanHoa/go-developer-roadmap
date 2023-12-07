package storage

import (
	"context"
	"go-gomail/src/module/entity"
)

func (s *sqlStorage) CreateUser(ctx context.Context, user *entity.UserCreatable) (*uint, error) {
	if err := s.db.Table(entity.UserCreatable{}.GetTableName()).Create(&user).Error; err != nil {
		return nil, err
	}
	return &user.ID, nil
}
