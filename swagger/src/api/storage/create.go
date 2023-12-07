package storage

import (
	"context"
	"fmt"
	"go-swagger/src/api/entity"
)

func (s *sqlStorage) Create(ctx context.Context, account entity.AccountCreatable) (*string, error) {
	if err := s.db.Table(entity.AccountCreatable{}.GetTableName()).Create(&account).Error; err != nil {
		fmt.Println("Error while create an account in account storage: " + err.Error())
		return nil, err
	}
	return account.Email, nil
}
