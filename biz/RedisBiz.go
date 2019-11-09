package biz

import (
	db "github.com/catcherwong/rest-api-sample/common"
	"time"
)

type RedisBiz struct {
}

func NewRedisBiz() *RedisBiz {
	return &RedisBiz{}
}

func (b RedisBiz) SetString(k, v string, exp int64) (bool, error) {

	t := time.Second * time.Duration(exp)

	err := db.RedisClient.Set(k, v, t).Err()

	if err != nil {
		return false, err
	}

	return true, nil
}

func (b RedisBiz) DeleteValue(k string) (bool, error) {

	err := db.RedisClient.Del(k).Err()

	if err != nil {

		return false, err
	}

	return true, nil
}

func (b RedisBiz) GetString(k string) (string, error) {

	v, err := db.RedisClient.Get(k).Result()

	if err != nil {

		return "", err
	}

	return v, err
}
