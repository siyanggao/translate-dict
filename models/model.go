package models

type BaseResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Word struct {
	Word       string `orm:"pk;size(20)"`
	Definition string `orm:"size(60)"`
}
