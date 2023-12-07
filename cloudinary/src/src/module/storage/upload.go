package storage

import (
	"context"
	"fmt"
	"go-cloudinary/src/src/module/entity"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"gorm.io/gorm"
)

type uploadStorage struct {
	db  *gorm.DB
	cld *cloudinary.Cloudinary
}

func NewUploadStore(db *gorm.DB, cld *cloudinary.Cloudinary) *uploadStorage {
	return &uploadStorage{db: db, cld: cld}
}

func (s *uploadStorage) SaveThumbInfo(ctx context.Context, thumbRes entity.ThumbCreatable, res *uploader.UploadResult) (*uint, error) {
	thumb := entity.ThumbCreatable{Url: &res.URL, Description: thumbRes.Description}
	if err := s.db.Table(entity.ThumbCreatable{}.GetTableName()).Create(&thumb).Error; err != nil {
		fmt.Println("Error while save thumb information to database!")
		return nil, err
	} else {
		return &thumb.ID, nil
	}
}

func (s *uploadStorage) Upload(ctx context.Context, filePath string) (*uploader.UploadResult, error) {
	if resp, err := s.cld.Upload.Upload(ctx, filePath, uploader.UploadParams{
		Folder: "go-cloudinary",
	}); err != nil {
		fmt.Println("Error while upload image to cloudinary in repository: " + err.Error())
		return nil, err
	} else {
		return resp, nil
	}
}
