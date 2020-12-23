package dao

import (
	"encoding/json"
	"errors"
	"go-basic/lib"
	"strconv"

	"go.uber.org/zap"

	"github.com/gomodule/redigo/redis"
)

const (
	REDIS_PREFIX = "app:user:"
)

type User struct {
	UserId int64  `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Age    int    `json:"age"`
	Email  string `json:"email"`
}

func (user *User) GetUserFromCache(userId string) error {
	conn := lib.RedisPool.Get()
	defer conn.Close()

	key := REDIS_PREFIX + userId
	userJsonData, err := redis.Bytes(conn.Do("get", key))
	if err != nil {
		lib.Logger.Error("getUserFromCache get fail", zap.String("error", err.Error()))
		return err
	}

	err = json.Unmarshal(userJsonData, user)
	if err != nil {
		lib.Logger.Error("getUserFromCache unmarshal fail", zap.String("error", err.Error()))
		return err
	}
	return nil
}

func (user *User) SetUserCache() error {
	if user == nil {
		return errors.New("invalid user")
	}
	conn := lib.RedisPool.Get()
	defer conn.Close()

	userJsonData, err := json.Marshal(user)
	if err != nil {
		lib.Logger.Error("setUserCache fail", zap.String("error", err.Error()))
		return err
	}

	key := REDIS_PREFIX + strconv.FormatInt(user.UserId, 10)
	_, err = conn.Do("set", key, userJsonData)
	if err != nil {
		lib.Logger.Error("setUserCache fail", zap.String("error", err.Error()))
		return err
	}
	return nil
}

func (user *User) GetUserFromDB(userId string) error {
	err := lib.DB.Get(user, "SELECT * FROM user WHERE id=?", userId)
	if err != nil {
		lib.Logger.Error("GetUserFromDB fail", zap.String("error", err.Error()))
		return err
	}
	return nil
}

func (user *User) GetUserInfoFromES(userId string) error {
	response, err := lib.ESClient.GetSource("user", userId)
	if err != nil {
		lib.Logger.Error("GetUserFromES fail:", zap.String("error", err.Error()))
		return err
	}
	if response.StatusCode != 200 {
		lib.Logger.Error("GetUserFromES fail:", zap.String("response", response.String()))
		return errors.New(response.String())
	}

	err = json.NewDecoder(response.Body).Decode(user)
	if err != nil {
		lib.Logger.Error("GetUserFromES json unmarshal fail:", zap.String("error", err.Error()))
		return err
	}
	return nil
}
