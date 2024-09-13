package models

type ParamFileHash struct {
	FileSha1 string `json:"filesha1"`
}
type ParamUpdatefile struct {
	FileSha1 string `json:"filesha1" binding:"required"`
	Filename string `json:"filename" binding:"required"`
}
