package utils

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
