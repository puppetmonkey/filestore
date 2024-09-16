package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "userID"
const CtxUserNameKey = "username"

var ErrorUserNotLogin = errors.New("用户未登录")

func getCurrentUserName(c *gin.Context) (username string, err error) {
	uname, ok := c.Get(CtxUserNameKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	username, ok = uname.(string)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

// getCurrentUserID 获取当前登录的用户ID
func getCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

func getPageInfo(c *gin.Context) (int64, int64) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")

	var (
		page int64
		size int64
		err  error
	)

	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return page, size
}
