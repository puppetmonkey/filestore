package logic

import (
	"errors"
	"filestore/dao/ipfs"
	"filestore/dao/mysql"
	"filestore/dao/redis"
	"filestore/models"
	"filestore/pkg/eth"
	"filestore/util"
	"io"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	Errorpartfilevail = errors.New("invalid request")
)

// // UpdateFileMetaDB : 新增/更新文件元信息到mysql中
// func UpdateFileMetaDB(fmeta models.FileMeta) bool {
// 	return .OnFileUploadFinished(
// 		fmeta.FileSha1, fmeta.FileName, fmeta.FileSize, fmeta.Location)
// }

func createnewfile(fileMeta *models.FileMeta, o *models.Originfile) (*models.FileMeta, error) {
	newFile, err := os.Create(fileMeta.Location)
	if err != nil {
		zap.L().Error("create new_file failed", zap.Error(err))
		return nil, err
	}
	defer newFile.Close()

	fileMeta.FileSize, err = io.Copy(newFile, o.File)
	if err != nil {

		zap.L().Error("copy file failed", zap.Error(err))
		return nil, err
	}
	newFile.Seek(0, 0)
	fileMeta.FileSha1 = util.FileSha1(newFile)

	// // 游标重新回到文件头部
	// newFile.Seek(0, 0)

	return fileMeta, nil
}
func Uploadfile(o *models.Originfile, u *models.User) (*models.FileMeta, error) {
	//file元数据构建
	fileMeta := &models.FileMeta{
		FileName: o.Head.Filename,
		Location: "/home/xusir/file/" + o.Head.Filename,
	}
	fileMeta, err := createnewfile(fileMeta, o)
	if err != nil {
		zap.L().Error("create file failed", zap.Error(err))
		return nil, err
	}
	//上传ipfs
	cid, err := ipfs.Uploadfile(fileMeta, u)
	if err != nil {
		zap.L().Error("Ipfs.Uploadfile(fileMeta, u) FAILED", zap.Error(err))
		return nil, err
	}

	//将cid 存到链上
	_, err = eth.StoreFile2eth(cid, fileMeta)

	if err != nil {
		zap.L().Error("eth.NewFileStorage FAILED", zap.Error(err))
		return nil, err
	}

	err = mysql.UpdateFileMeta(fileMeta)
	if err != nil {
		zap.L().Error("mysql db upload file failed", zap.Error(err))
		return nil, err
	}
	userfile := &models.UserFile{
		UserName: u.Username,
		FileHash: fileMeta.FileSha1,
		FileName: fileMeta.FileName,
		FileSize: fileMeta.FileSize,
	}
	err = mysql.UploadUserfile(userfile)
	if err != nil {
		zap.L().Error("mysql db upload userfile failed", zap.Error(err))
		return nil, err
	}
	return fileMeta, nil
}

func UpdateFile(p *models.ParamUpdatefile) (data *models.FileMeta, err error) {
	//file元数据构建
	past_filemeta, err := mysql.GetFileMeta(p.FileSha1)
	if err != nil {
		zap.L().Error("mysql.GetFileMeta(p.FileSha1) failed",
			zap.String("filesha1", p.FileSha1),
			zap.Error(err))
		return
	}

	curMeta := &models.FileMeta{
		FileSha1: p.FileSha1,
		FileName: p.Filename,
		FileSize: past_filemeta.FileSize,
		Location: past_filemeta.Location,
	}
	err = mysql.UpdateFileMeta(curMeta)
	if err != nil {
		zap.L().Error("mysql db update file failed", zap.Error(err))
		return nil, err
	}
	return curMeta, nil
}
func DeleteFileMeta(p string) (data *models.FileMeta, err error) {
	//file元数据构建
	past_filemeta, err := mysql.GetFileMeta(p)
	if err != nil {
		zap.L().Error("mysql.GetFileMeta(p.FileSha1) failed",
			zap.String("filesha1", p),
			zap.Error(err))
		return
	}
	err = mysql.DeleteFileMeta(p)
	if err != nil {
		zap.L().Error("mysql db delete file failed", zap.Error(err))
		return nil, err
	}
	return past_filemeta, nil
}
func SetMultipartUploadInfo(p *models.MultipartUploadInfo) (err error) {
	return redis.SetMultipartUploadInfo(p)
}

func UploadPartHandler(c *gin.Context, p *models.UploadpartInfo) (err error) {
	fpath := "/home/xusir/data/" + p.UploadID + "/" + p.ChunkIndex
	zap.L().Debug("Creating file at path", zap.String("fpath", fpath))

	// 创建目录
	err = os.MkdirAll(path.Dir(fpath), 0744)
	if err != nil {
		zap.L().Error("Failed to create directory", zap.String("dir", path.Dir(fpath)), zap.Error(err))
		return err
	}

	// 创建文件
	fd, err := os.Create(fpath)
	if err != nil {
		zap.L().Error("Create file failed", zap.String("fpath", fpath), zap.Error(err))
		return err
	}
	defer fd.Close()

	// 4. 从请求体读取文件流，写入本地文件
	buf := make([]byte, 1024*1024) // 定义 1MB 的缓冲区
	for {
		n, err := c.Request.Body.Read(buf) // 从请求体中读取文件数据
		fd.Write(buf[:n])                  // 将读取到的数据写入文件
		if err != nil {
			break // 当遇到 EOF 或者其他错误时跳出循环
		}
	}
	// 4. 更新redis缓存状态
	err = redis.Updateuploadchkidx(p)
	return
}
func CompleteUpload(p *models.UserFile, upid string) (err error) {
	result, err := redis.GetDatafromID("MP_" + upid)

	if err != nil {
		zap.L().Error("edis.GetDatafromID", zap.Error(err))
		return err
	}
	// 初始化总分块数和已上传分块数
	totalCount := 0
	chunkCount := 0

	// 遍历结果，检查是否所有分块已上传
	for k, v := range result {
		if k == "chunkcount" {
			totalCount, _ = strconv.Atoi(v) // 获取总分块数
		} else if strings.HasPrefix(k, "chkidx_") && v == "1" {
			chunkCount++ // 统计已上传的分块数
		}
	}

	// 判断是否所有分块都已经上传完成
	if totalCount != chunkCount {
		zap.L().Error("totalCount != chunkCount", zap.Int("totalCount:", totalCount), zap.Int("chukcount:", chunkCount), zap.Error(err))
		return Errorpartfilevail
	}

	// 5. 更新唯一文件表及用户文件表
	filemta := &models.FileMeta{
		FileName: p.FileName,
		FileSha1: p.FileHash,
		FileSize: p.FileSize,
		Location: "",
	}
	err = mysql.UpdateFileMeta(filemta)

	if err != nil {
		zap.L().Error("mysql.UpdateFileMeta", zap.Error(err))
		return err
	}
	err = mysql.UploadUserfile(p)

	if err != nil {
		zap.L().Error("mysql.UploadUserfile", zap.Error(err))
		return err
	}

	return nil
}
