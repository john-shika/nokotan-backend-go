package schemas

type Message struct {
	StatusOk   bool   `json:"statusOk"`
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

func NewMessage(statusOk bool, statusCode int, status string, message string, data any) *Message {
	return &Message{
		StatusOk:   statusOk,
		StatusCode: statusCode,
		Status:     status,
		Message:    message,
		Data:       data,
	}
}

func NewMessageSuccess(message string, data any) *Message {
	return NewMessage(true, 200, "OK", message, data)
}

func NewMessageCreated(message string, data any) *Message {
	return NewMessage(true, 201, "created", message, data)
}

func NewMessageUnauthorized(message string) *Message {
	return NewMessage(false, 401, "unauthorized", message, nil)
}

func NewMessageBadRequest(message string) *Message {
	return NewMessage(false, 400, "bad_request", message, nil)
}

func NewMessageNotFound(message string) *Message {
	return NewMessage(false, 404, "not_found", message, nil)
}

func NewMessageError(message string) *Message {
	return NewMessage(false, 500, "error", message, nil)
}
