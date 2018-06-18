package models

type Conf struct {
	Name   string `form:"name"`
	Header string `form:"header"`
	Banner string `form:"banner"`
	Logo   string `form:"logo"`
	Msg    string `form:"msg"`
}
