package helper

type Token_Response struct {
	Code int               `json:"code"`
	Msg  string            `json:"msg"`
	Data map[string]string `json:"data"`
}

type Register_Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type Login_Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type Gene_Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
