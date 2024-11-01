package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"filestore/models"
)

// 把每一步数据库操作封装成函数
// 待logic层根据业务需求调用

const secret = "liwenzhou.com"

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from tbl_user where user_name = ?`
	var count int64
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// InsertUser 想数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	user.Password = encryptPassword(user.Password)
	user.PrivateKey = encryptPassword(user.PrivateKey)
	// 执行SQL语句入库
	sqlStr := `insert into tbl_user(user_id, user_name, user_pwd, private_key, public_key) values(?,?,?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password, user.PrivateKey, user.PublicKey)
	return
}

// encryptPassword 密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *models.User) (err error) {
	oPassword := user.Password // 用户登录的密码
	sqlStr := `select user_id, user_name, user_pwd from tbl_user where user_name=?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		// 查询数据库失败
		return err
	}
	// 判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return
}

// GetUserById 根据id获取用户信息
func GetUserById(uid int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id, user_name, signup_at, last_active from tbl_user where user_id = ?`
	err = db.Get(user, sqlStr, uid)
	return
}

func GetUserByName(uid string) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_name, user_id, signup_at, last_active from tbl_user where user_name = ?`
	err = db.Get(user, sqlStr, uid)
	return
}
