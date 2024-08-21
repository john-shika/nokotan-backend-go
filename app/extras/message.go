package extras

type Message struct {
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

func NewMessage(status string, statusCode int, message string, data any) *Message {
	return &Message{
		Status:     status,
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
}

func NewMessageSuccess(message string, data any) *Message {
	return NewMessage("ok", 200, message, data)
}

func NewMessageCreated(message string, data any) *Message {
	return NewMessage("created", 201, message, data)
}

func NewMessageUnauthorized(message string) *Message {
	return NewMessage("unauthorized", 401, message, nil)
}

func NewMessageBadRequest(message string) *Message {
	return NewMessage("bad_request", 400, message, nil)
}

func NewMessageNotFound(message string) *Message {
	return NewMessage("not_found", 404, message, nil)
}

func NewMessageError(message string) *Message {
	return NewMessage("error", 500, message, nil)
}
