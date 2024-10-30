package models

import "time"

type User struct {
	UserID     int64     `db:"user_id"`
	Username   string    `db:"user_name"`
	Password   string    `db:"user_pwd"`
	Token      string    `db:"token"` // 如果你希望在数据库中也存储 Token
	SignupAt   time.Time `db:"signup_at"`
	LastActive time.Time `db:"last_active"`
	Status     int       `db:"status"`
	Profile    string    `db:"profile"`
	PrivateKey string    `db:"private_key"` // 添加 PrivateKey 字段
	PublicKey  string    `db:"public_key"`  // 添加 PublicKey 字段
}
