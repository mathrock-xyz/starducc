package model

type File struct {
	ID     string
	Name   string `gorm:"uniqueIndex:idx_user_file_name"`
	Hash   string
	Size   int64
	UserID string `gorm:"uniqueIndex:idx_user_file_name"`
}
