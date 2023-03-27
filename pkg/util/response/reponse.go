package response

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Errors     interface{} `json:"errors,omitempty"`
}

func (r *Response) SuccessResponse(data interface{}) {

}

type UserData struct {
	ID    uint
	Email string
	Phone string
}
