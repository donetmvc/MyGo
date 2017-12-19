package models

import (
	"fmt"
	"errors"
)

type LoginInfo struct {
	Id			int	`xorm:"BigInt"`
	Name		string
	Password	string
	Detail		*Detail `xorm:"extends"`
}

type Detail struct {
	Id			int	`xorm:"BigInt 'detail_id'"`
	Email		string
	Address		string
	Phone		string
}

func init() {
	engine.Sync2(new(LoginInfo), new(Detail))
}

func Login(name string, password string) (loginResult LoginInfo, err error) {
	var loginInfo LoginInfo
	has, err := engine.Alias("e").Where("e.name = ?", name).And("e.password = ?", password).Get(&loginInfo)

	fmt.Println(has)

	if !has {
		err = errors.New("用户不存在")
		return loginInfo, err
	} 

	return loginInfo, nil
}

func CreateUser(l *LoginInfo) (user int64, err error) {
	id, err := engine.Insert(l)

	if err != nil {
		return -1, err
	}

	return id, err
}