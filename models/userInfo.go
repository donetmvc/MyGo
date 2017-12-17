package models

import (
	"time"
)

type UserInfo struct {
	Id int64
	UserName string
	Password string
	Email string

	InDateTime   time.Time `xorm:"created"`
	ModiDateTime time.Time `xorm:"updated"` 
}