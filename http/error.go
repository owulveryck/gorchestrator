package http

type jsonErr struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}
