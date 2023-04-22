package errorcode

type ErrorCode struct {
	Code    string `json:"code" csv:"code"`
	Message string `json:"message" csv:"message"`
	Param   string `json:"param"`
}

func New(code, message string) ErrorCode {
	return ErrorCode{
		Code:    code,
		Message: message,
	}
}

func (ec ErrorCode) SetParam(s string) ErrorCode {
	ec.Param = s
	return ec
}
