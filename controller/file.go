package controller

import (
	"filestore/dao/mysql"
	"filestore/logic"
	"filestore/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UploadHandler(c *gin.Context) {
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
	data, err := logic.Uploadfile(o)
	if err != nil {
		zap.L().Error("logic.Uploadfile(o) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
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
