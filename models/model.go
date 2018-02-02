package models

type BaseResponse struct {
	Code int
	Msg  string
	Data interface{}
}
