package models

import (
	"mime/multipart"
	"time"
)

type FileMeta struct {
	FileSha1 string    `json:"filesha1" db:"file_sha1"`    // 与数据库中的 char(40) 对应
	FileName string    `json:"filename" db:"file_name"`    // 与数据库中的 varchar(256) 对应
	FileSize int64     `json:"filesize" db:"file_size"`    // 与数据库中的 bigint 对应
	Location string    `json:"fileaddr" db:"file_addr"`    // 对应 file_addr (varchar 1024)
	Status   int       `json:"status" db:"status"`         // 对应 status 列 (int)
	CreateAt time.Time `json:"create_at" db:"create_at"`   // 对应 create_at 列 (datetime)
	UpdateAt time.Time `json:"update_at " db:"update_at "` // 对应 update_at 列 (datetime)
}

type Originfile struct {
	File multipart.File
	Head *multipart.FileHeader
}
