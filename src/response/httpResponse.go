package response

type (
	HttpResponse struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	ErrorResponse struct {
		Message string `json:"message"`
	}
)
