package logic

import (
	"filestore/dao/mysql"
	"filestore/models"
	"filestore/util"
	"io"
	"os"

	"go.uber.org/zap"
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
func Uploadfile(o *models.Originfile) (*models.FileMeta, error) {
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
	err = mysql.UpdateFileMeta(fileMeta)
	if err != nil {
		zap.L().Error("mysql db upload file failed", zap.Error(err))
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
