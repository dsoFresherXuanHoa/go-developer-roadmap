package storage

import (
	"context"
	"go-auth2/src/modules/auth/entity"
)

func (s *sqlStorage) SignUp(ctx context.Context, user *entity.UserCreatable) (*uint, error) {
	if err := s.db.Table(entity.UserCreatable{}.GetTableName()).Create(&user).Error; err != nil {
		return nil, err
	}
	return &user.ID, nil
}
