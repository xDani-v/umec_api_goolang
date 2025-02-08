package utils

type ResponseMsg struct {
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
	Status int         `json:"code"`
}
