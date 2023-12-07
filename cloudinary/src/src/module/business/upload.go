package business

import (
	"context"
	"fmt"
	"go-cloudinary/src/src/module/entity"
	"path/filepath"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type UploadStorage interface {
	Upload(ctx context.Context, filePath string) (*uploader.UploadResult, error)
	SaveThumbInfo(ctx context.Context, thumbRes entity.ThumbCreatable, res *uploader.UploadResult) (*uint, error)
}

type uploadBusiness struct {
	storage UploadStorage
}

func NewUploadBusiness(storage UploadStorage) *uploadBusiness {
	return &uploadBusiness{storage: storage}
}

func (business *uploadBusiness) Upload(ctx context.Context, filePath string) (*uploader.UploadResult, error) {
	if res, err := business.storage.Upload(ctx, filepath.Base(filePath)); err != nil {
		fmt.Println("Error while upload in service: " + err.Error())
		return nil, err
	} else {
		return res, nil
	}
}
func (business *uploadBusiness) SaveThumbInfo(ctx context.Context, thumbRes entity.ThumbCreatable, res *uploader.UploadResult) (*uint, error) {
	if res, err := business.storage.SaveThumbInfo(ctx, thumbRes, res); err != nil {
		fmt.Println("Error while save thumb information in service: " + err.Error())
		return nil, err
	} else {
		return res, nil
	}
}
