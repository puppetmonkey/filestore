package models

import "time"

type User struct {
	UserID     int64  `db:"user_id"`
	Username   string `db:"user_name"`
	Password   string `db:"user_pwd"`
	Token      string
	SignupAt   time.Time `db:"signup_at"`
	LastActive time.Time `db:"last_active"`
	Status     int       `db:"status"`
	Profile    string    `db:"profile"`
}
