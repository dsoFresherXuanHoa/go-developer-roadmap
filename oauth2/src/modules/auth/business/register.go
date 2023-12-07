package business

import (
	"context"
	"go-auth2/src/modules/auth/entity"
)

type SignUpStorage interface {
	SignUp(ctx context.Context, user *entity.UserCreatable) (*uint, error)
}

type signUpBusiness struct {
	storage SignUpStorage
}

func NewSignUpBusiness(storage SignUpStorage) *signUpBusiness {
	return &signUpBusiness{storage: storage}
}

func (business *signUpBusiness) SignUp(ctx context.Context, user *entity.UserCreatable) (*uint, error) {
	if userId, err := business.storage.SignUp(ctx, user); err != nil {
		return nil, err
	} else {
		return userId, nil
	}
}
