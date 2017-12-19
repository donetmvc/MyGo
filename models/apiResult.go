package models

type APIResult struct {
	Success			bool
	Result			string
	Error			*APIError	`xorm:"extends"`
}

type APIError struct {
	Code			int
	Message			string
	Detail			string
}