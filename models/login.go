package models

import (
	"github.com/astaxie/beego/orm"
	"errors"
)

type LoginInfo struct {
	Id			int
	Name		string
	Password	string
	Detail		*Detail `orm:"rel(one)"`
}

type Detail struct {
	Id			int
	Email		string
	Address		string
	Phone		string
}

func Login(name string, password string) (loginResult LoginInfo, err error) {
	db := orm.NewOrm()
	db.Using("eoas")
	loginInfo := LoginInfo{ Name: name, Password: password}

	err = db.Read(&loginInfo)

	if err == orm.ErrNoRows {
		err = errors.New("用户不存在")
	} else if err == orm.ErrMissPK {
		err = errors.New("找不到主键")
	}

	return loginInfo, err
}

// func CreateUser(user *U)