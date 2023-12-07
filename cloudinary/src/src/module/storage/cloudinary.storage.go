package storage

import (
	"github.com/cloudinary/cloudinary-go/v2"
)

type cloudinaryStorage struct {
	cld *cloudinary.Cloudinary
}

func NewCloudinaryStore(cld *cloudinary.Cloudinary) *cloudinaryStorage {
	return &cloudinaryStorage{cld: cld}
}
