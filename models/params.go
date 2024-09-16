package models

type ParamFileHash struct {
	FileSha1 string `json:"filesha1"`
}
type ParamUpdatefile struct {
	FileSha1 string `json:"filesha1" binding:"required"`
	Filename string `json:"filename" binding:"required"`
}

type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登录请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type ParamUsername struct {
	Username string `json:"username" binding:"required"`
}
