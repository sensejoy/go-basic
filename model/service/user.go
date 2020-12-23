package service

import (
	"go-basic/model/dao"
)

func GetUser(userId string) (*dao.User, error) {
	user := new(dao.User)
	/*
		if err := user.GetUserFromCache(userId); err != nil {
			return nil, err
		} else {
			return user, nil
		}
	*/
	//get user info from mysql\es\mongodb
	if err := user.GetUserFromDB(userId); err != nil {
		return nil, err
	}
	if err := user.GetUserInfoFromES(userId); err != nil {
		return nil, err
	}

	user.SetUserCache()

	return user, nil
}
