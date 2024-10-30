package eth

import "filestore/models"

func StoreFile2eth(cid string, u *models.FileMeta) (string, error) {
	client, err := ConnectToClient("http://localhost:8545")
	if err != nil {
		return "", err
	}

	instance, err := GetContractInstance(client, "e7f1725e7734ce288f8367e1bb143e90bb3f0512")
	if err != nil {
		return "", err
	}

	auth, err := GetAuth(client)
	if err != nil {
		return "", err
	}

	txHash, err := StoreFile(instance, auth, u.FileSha1, cid, u.FileName)
	if err != nil {
		return "", err
	}

	return txHash, nil
}
func Getfilecid(filesha1 string) (cid string, err error) {
	client, err := ConnectToClient("http://localhost:8545")
	if err != nil {
		return "", err
	}

	instance, err := GetContractInstance(client, "e7f1725e7734ce288f8367e1bb143e90bb3f0512")
	if err != nil {
		return "", err
	}
	_, cid, _, _, err = GetFileByHash(instance, filesha1)
	if err != nil {
		return "", err
	}

	return cid, nil
}
