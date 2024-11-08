package schemas

import "nokowebapi/cores"

type MessageBody struct {
	StatusOk   bool   `json:"statusOk,required"`
	StatusCode int    `json:"statusCode,required"`
	Status     string `json:"status,required"`
	Message    string `json:"message,required"`
	Data       any    `json:"data,required"`
}

func NewMessageBody(statusOk bool, statusCode int, status string, message string, data any) *MessageBody {
	return &MessageBody{
		StatusOk:   statusOk,
		StatusCode: statusCode,
		Status:     status,
		Message:    message,
		Data:       data,
	}
}

func NewMessageBodyOk(message string, data any) *MessageBody {
	status := cores.Default[cores.HttpStatusCodeValue]().FromCode(cores.HttpStatusCodeOk)
	return NewMessageBody(false, cores.HttpStatusCodeOk, string(status), message, data)
}

func NewMessageBodyCreated(message string, data any) *MessageBody {
	status := cores.Default[cores.HttpStatusCodeValue]().FromCode(cores.HttpStatusCodeCreated)
	return NewMessageBody(false, cores.HttpStatusCodeCreated, string(status), message, data)
}

func NewMessageBodyUnauthorized(message string, data any) *MessageBody {
	status := cores.Default[cores.HttpStatusCodeValue]().FromCode(cores.HttpStatusCodeUnauthorized)
	return NewMessageBody(false, cores.HttpStatusCodeUnauthorized, string(status), message, data)
}

func NewMessageBodyBadRequest(message string, data any) *MessageBody {
	status := cores.Default[cores.HttpStatusCodeValue]().FromCode(cores.HttpStatusCodeBadRequest)
	return NewMessageBody(false, cores.HttpStatusCodeBadRequest, string(status), message, data)
}

func NewMessageBodyNotFound(message string, data any) *MessageBody {
	status := cores.Default[cores.HttpStatusCodeValue]().FromCode(cores.HttpStatusCodeNotFound)
	return NewMessageBody(false, cores.HttpStatusCodeNotFound, string(status), message, data)
}

func NewMessageBodyInternalServerError(message string, data any) *MessageBody {
	status := cores.Default[cores.HttpStatusCodeValue]().FromCode(cores.HttpStatusCodeInternalServerError)
	return NewMessageBody(false, cores.HttpStatusCodeInternalServerError, string(status), message, data)
}
