package models

type UserFile struct {
	UserName    string `db:"user_name"`
	FileHash    string `db:"file_sha1"`
	FileName    string `db:"file_name"`
	FileSize    int64  `db:"file_size"`
	UploadAt    string `db:"upload_at"`
	LastUpdated string `db:"last_update"`
}
