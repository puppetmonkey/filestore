package ipfs

import (
	"bytes"
	"filestore/models"
	"os"

	"go.uber.org/zap"
)

// UploadData 上传数据到 IPFS
func UploadData(data []byte) (string, error) {
	// 创建一个 Reader 以供 IPFS 上传
	reader := bytes.NewReader(data)

	// 上传数据到 IPFS
	cid, err := ipfs.Add(reader)
	if err != nil {
		zap.L().Error("failed to add data to IPFS", zap.Error(err))
		return "", err
	}

	zap.L().Info("data uploaded to IPFS", zap.String("CID", cid))
	return cid, nil
}

// UploadData 上传数据到IPFS
func Uploadfile(o *models.FileMeta, u *models.User) (cid string, err error) {
	// 读取文件数据以准备上传到 IPFS
	fileData, err := os.ReadFile(o.Location)
	if err != nil {
		zap.L().Error("ipfs read file failed", zap.Error(err))
		return "", err
	}

	// 上传到 IPFS
	cid, err = UploadData(fileData)
	if err != nil {
		zap.L().Error("upload to IPFS failed", zap.Error(err))
		return "", err
	}

	return cid, nil
}
