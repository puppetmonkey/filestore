package mysql

import (
	"filestore/models"
)

func UploadUserfile(uf *models.UserFile) (err error) {
	sqlStr := `
		INSERT ignore INTO tbl_user_file (
			user_name, file_sha1, file_size, file_name
		) VALUES (?, ?, ?, ?)
	`
	_, err = db.Exec(sqlStr, uf.UserName, uf.FileHash, uf.FileSize, uf.FileName)
	return
}
func QueryUserFileMetas(username string, limit int) ([]models.UserFile, error) {
	sqlStr := `
		SELECT file_sha1, file_name, file_size, upload_at, last_update
		FROM tbl_user_file
		WHERE user_name = ?
		ORDER BY upload_at DESC
		LIMIT ?
	`
	userFiles := make([]models.UserFile, 0, limit)
	err := db.Select(&userFiles, sqlStr, username, limit)
	return userFiles, err
}
