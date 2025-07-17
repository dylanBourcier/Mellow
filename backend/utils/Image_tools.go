package utils

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func GetFullImageURL(imagePath *string) *string {
	if imagePath == nil || *imagePath == "" {
		empty := ""
		return &empty
	}
	fullPath := "/api/uploads/" + *imagePath
	return &fullPath
}

func GetFullImageURLAvatar(imagePath *string) *string {
	if imagePath == nil {
		return nil
	}
	fullPath := "/api/uploads/avatars/" + *imagePath
	return &fullPath
}

func HandleImageUpload(header *multipart.FileHeader, file multipart.File, path []string) (*string, error) {
	var image_url *string
	defer file.Close()
	fileName := uuid.New().String() + filepath.Ext(header.Filename)
	savePath := filepath.Join(append(path, fileName)...)
	if err := os.MkdirAll(filepath.Dir(savePath), os.ModePerm); err != nil {
		return nil, err
	}
	dst, err := os.Create(savePath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()
	if _, err := io.Copy(dst, file); err != nil {
		return nil, err
	}
	image_url = &fileName
	return image_url, nil
}
