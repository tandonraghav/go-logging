package response

type GenericResp struct {
	Status string
	Result interface{}
	Err    ErrorResp
}

type ErrorResp struct {
	ErrorCode string
	ErrorDesc string
}
