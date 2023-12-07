package business

import (
	"context"
	"fmt"
	"go-swagger/src/api/entity"
)

type AccountStorage interface {
	Create(ctx context.Context, account entity.AccountCreatable) (*string, error)
}

type createBusiness struct {
	storage AccountStorage
}

func NewCreateBusiness(storage AccountStorage) *createBusiness {
	return &createBusiness{storage: storage}
}

func (business *createBusiness) SignUp(ctx context.Context, account *entity.AccountCreatable) (*string, error) {
	if err := account.Validate(); err != nil {
		fmt.Println("Error while validate account: " + err.Error())
		return nil, err
	} else {
		accountCreatable := account.Mask()
		if registerEmail, err := business.storage.Create(ctx, *accountCreatable); err != nil {
			fmt.Println("Error while sign up in account business: " + err.Error())
			return nil, err
		} else {
			return registerEmail, nil
		}
	}
}
