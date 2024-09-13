package mysql

import "filestore/models"

func UpdateFileMeta(filemeta *models.FileMeta) (err error) {
	sqlStr := `
        INSERT INTO tbl_file (
            file_sha1, file_name, file_size, file_addr
        ) VALUES (?, ?, ?, ?)
        ON DUPLICATE KEY UPDATE 
            file_name = VALUES(file_name),
            file_size = VALUES(file_size),
            file_addr = VALUES(file_addr)
    `
	_, err = db.Exec(sqlStr, filemeta.FileSha1, filemeta.FileName, filemeta.FileSize, filemeta.Location)
	return
}

func GetFileMeta(fileSha1 string) (filemeta *models.FileMeta, err error) {
	filemeta = new(models.FileMeta)
	sqlStr := `
        SELECT
            file_sha1, file_name, file_size, file_addr
        FROM tbl_file
        WHERE file_sha1 = ?
    `
	err = db.Get(filemeta, sqlStr, fileSha1)
	return
}

func DeleteFileMeta(fileSha1 string) (err error) {
	// SQL 删除语句，根据 file_sha1 删除记录
	sqlStr := `DELETE FROM tbl_file WHERE file_sha1 = ?`

	// 执行 SQL 删除语句
	_, err = db.Exec(sqlStr, fileSha1)
	return
}
