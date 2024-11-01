package logic

import (
	"filestore/dao/mysql"
	"filestore/models"
	"filestore/pkg/jwt"
	"filestore/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 1.判断用户存不存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 2.生成UID
	userID := snowflake.GenID()

	// //构建用户私钥，公钥
	// privKey, address := eth.GenerateKeyPair()
	// 构造一个User实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
		// PrivateKey: hex.EncodeToString(privKey.D.Bytes()),
		// PublicKey:  address.Hex(),
	}
	// 3.保存进数据库
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 传递的是指针，就能拿到user.UserID
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	// 生成JWT
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}

func GetUserinfo(p string) (user *models.User, err error) {
	user = &models.User{}
	if user, err = mysql.GetUserByName(p); err != nil {
		return nil, err
	}
	return user, nil
}
