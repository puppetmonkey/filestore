package controller

import (
	"filestore/dao/mysql"
	"filestore/logic"
	"filestore/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UploadHandler(c *gin.Context) {
	//文件元数据上传
	file, head, err := c.Request.FormFile("file")
	if err != nil {
		fmt.Printf("Failed to get data, err: %s\n", err.Error())
		c.String(http.StatusBadRequest, "Failed to get file")
		return
	}
	defer file.Close()
	o := &models.Originfile{
		File: file,
		Head: head,
	}
	//用户，文件元数据上传
	username := c.Query("username")
	if username == "" {
		zap.L().Error("no userlogin", zap.Error(ErrorUserNotLogin))
		ResponseError(c, CodeUserNotExist)
		return
	}
	u := &models.User{
		Username: username,
	}
	zap.L().Debug("username", zap.String("username", username))
	_, err = logic.Uploadfile(o, u)
	if err != nil {
		zap.L().Error("logic.Uploadfile(o) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	c.Redirect(http.StatusFound, "http://"+c.Request.Host+"/static/view/home.html")
}
func GetFileMetaHandler(c *gin.Context) {
	p := c.Query("filesha1")
	zap.L().Debug("GetFileMetaHandler", zap.String("p", p))
	data, err := mysql.GetFileMeta(p)
	if err != nil {
		zap.L().Error("mysql.GetFileMeta(p)", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
func DownloadHandler(c *gin.Context) {
	p := c.Query("filesha1")
	data, err := mysql.GetFileMeta(p)
	if err != nil {
		zap.L().Error("mysql.GetFileMeta(p.FileSha1)", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func UpdateFileMeta(c *gin.Context) {
	p := &models.ParamUpdatefile{}
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("UpateFileMetaHandler with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.UpdateFile(p)
	if err != nil {
		zap.L().Error("logic.UpdateFile err", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func FileDeleteHandler(c *gin.Context) {
	p := c.Query("filesha1")
	data, err := logic.DeleteFileMeta(p)
	if err != nil {
		zap.L().Error("logic.UpdateFile err", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func FileQueryHandler(c *gin.Context) {
	limitStr := c.PostForm("limit")
	l, _ := strconv.Atoi(limitStr)
	zap.L().Debug("limit", zap.Int("limit", l))
	un := c.Query("username")
	// zap.L().Debug("username", zap.String("user", un))
	data, err := mysql.QueryUserFileMetas(un, l)
	if err != nil {
		zap.L().Error("mysql.QueryUserFileMetas", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func TryFastUploadHandler(c *gin.Context) {
	username := c.PostForm("username")
	filehash := c.PostForm("filehash")
	filename := c.PostForm("filename")
	filesize, _ := strconv.Atoi(c.PostForm("filesize"))
	// 2. 从文件表中查询相同hash的文件记录
	fileMeta, err := mysql.GetFileMeta(filehash)
	if err != nil {
		zap.L().Error(" mysql.GetFileMeta(filehash)", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 查不到记录则返回秒传失败
	if fileMeta == nil {
		ResponseError(c, Codefastuploadfail)
	}
	userfile := &models.UserFile{
		UserName: username,
		FileHash: filehash,
		FileName: filename,
		FileSize: int64(filesize),
	}
	err = mysql.UploadUserfile(userfile)
	if err != nil {
		ResponseSuccess(c, nil)
		return
	}
	ResponseError(c, CodeServerBusy)
}
