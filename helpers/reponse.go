package helpers

const (
	STATUS_NOT_FOUND = "Not Found"
	STATUS_SUCCESS   = "Success"
	MESSAGE_SUCCESS  = "Success"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *Response) Create(Status string, Message string, Data interface{}) *Response {
	r.Status = Status
	r.Message = Message
	r.Data = Data
	return r
}
