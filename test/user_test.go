package test

import (
	"fmt"
	"go-basic/model/dao"
	"testing"
)

func TestGetUserInfoFromCache(t *testing.T) {
	userId := "1"
	user := dao.GetUserFromCache(userId)
	fmt.Println(user)
	t.Run()
}
