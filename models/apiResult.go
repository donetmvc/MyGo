package models

type APIResult struct {
	Success			bool
	Result			string
	Error			*APIError	`orm:"rel(one)"`
}

type APIError struct {
	Code			int
	Message			string
	Detail			string
}