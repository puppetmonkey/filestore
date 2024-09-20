package redis

import "filestore/models"

func SetMultipartUploadInfo(p *models.MultipartUploadInfo) (err error) {
	pipeline := client.TxPipeline()
	pipeline.HSet("MP_"+p.UploadID, "chunkcount", p.ChunkCount)
	pipeline.HSet("MP_"+p.UploadID, "filehash", p.FileHash)
	pipeline.HSet("MP_"+p.UploadID, "filesize", p.FileSize)
	_, err = pipeline.Exec()
	return err
}

func Updateuploadchkidx(p *models.UploadpartInfo) (err error) {
	return client.HSet("MP_"+p.UploadID, "chkidx_"+p.ChunkIndex, 1).Err()
}

func GetDatafromID(id string) (resule map[string]string, err error) {
	return client.HGetAll(id).Result()
}
