package cores

type HttpStatusCode int
type HttpStatusText string

const (
	HttpStatusCodeContinue                      HttpStatusCode = 100
	HttpStatusCodeSwitchingProtocols                           = 101
	HttpStatusCodeProcessing                                   = 102
	HttpStatusCodeEarlyHints                                   = 103
	HttpStatusCodeOk                                           = 200
	HttpStatusCodeCreated                                      = 201
	HttpStatusCodeAccepted                                     = 202
	HttpStatusCodeNonAuthoritativeInformation                  = 203
	HttpStatusCodeNoContent                                    = 204
	HttpStatusCodeResetContent                                 = 205
	HttpStatusCodePartialContent                               = 206
	HttpStatusCodeMultiStatus                                  = 207
	HttpStatusCodeAlreadyReported                              = 208
	HttpStatusCodeImUsed                                       = 226
	HttpStatusCodeMultipleChoices                              = 300
	HttpStatusCodeMovedPermanently                             = 301
	HttpStatusCodeFound                                        = 302
	HttpStatusCodeSeeOther                                     = 303
	HttpStatusCodeNotModified                                  = 304
	HttpStatusCodeUseProxy                                     = 305
	HttpStatusCodeUnused                                       = 306
	HttpStatusCodeTemporaryRedirect                            = 307
	HttpStatusCodePermanentRedirect                            = 308
	HttpStatusCodeBadRequest                                   = 400
	HttpStatusCodeUnauthorized                                 = 401
	HttpStatusCodePaymentRequired                              = 402
	HttpStatusCodeForbidden                                    = 403
	HttpStatusCodeNotFound                                     = 404
	HttpStatusCodeMethodNotAllowed                             = 405
	HttpStatusCodeNotAcceptable                                = 406
	HttpStatusCodeProxyAuthenticationRequired                  = 407
	HttpStatusCodeRequestTimeout                               = 408
	HttpStatusCodeConflict                                     = 409
	HttpStatusCodeGone                                         = 410
	HttpStatusCodeLengthRequired                               = 411
	HttpStatusCodePreconditionFailed                           = 412
	HttpStatusCodePayloadTooLarge                              = 413
	HttpStatusCodeRequestUriTooLong                            = 414
	HttpStatusCodeUnsupportedMediaType                         = 415
	HttpStatusCodeRequestedRangeNotSatisfiable                 = 416
	HttpStatusCodeExpectationFailed                            = 417
	HttpStatusCodeImATeapot                                    = 418
	HttpStatusCodeInsufficientSpaceOnResource                  = 419
	HttpStatusCodeMethodFailure                                = 420
	HttpStatusCodeMisdirectedRequest                           = 421
	HttpStatusCodeUnprocessableEntity                          = 422
	HttpStatusCodeLocked                                       = 423
	HttpStatusCodeFailedDependency                             = 424
	HttpStatusCodeUpgradeRequired                              = 426
	HttpStatusCodePreconditionRequired                         = 428
	HttpStatusCodeTooManyRequests                              = 429
	HttpStatusCodeRequestHeaderFieldsTooLarge                  = 431
	HttpStatusCodeUnavailableForLegalReasons                   = 451
	HttpStatusCodeInternalServerError                          = 500
	HttpStatusCodeNotImplemented                               = 501
	HttpStatusCodeBadGateway                                   = 502
	HttpStatusCodeServiceUnavailable                           = 503
	HttpStatusCodeGatewayTimeout                               = 504
	HttpStatusCodeHttpVersionNotSupported                      = 505
	HttpStatusCodeVariantAlsoNegotiates                        = 506
	HttpStatusCodeInsufficientStorage                          = 507
	HttpStatusCodeLoopDetected                                 = 508
	HttpStatusCodeNotExtended                                  = 510
	HttpStatusCodeNetworkAuthenticationRequired                = 511
)

const (
	HttpStatusTextContinue                      HttpStatusText = "CONTINUE"
	HttpStatusTextSwitchingProtocols                           = "SWITCHING_PROTOCOLS"
	HttpStatusTextProcessing                                   = "PROCESSING"
	HttpStatusTextEarlyHints                                   = "EARLY_HINTS"
	HttpStatusTextOk                                           = "OK"
	HttpStatusTextCreated                                      = "CREATED"
	HttpStatusTextAccepted                                     = "ACCEPTED"
	HttpStatusTextNonAuthoritativeInformation                  = "NON_AUTHORITATIVE_INFORMATION"
	HttpStatusTextNoContent                                    = "NO_CONTENT"
	HttpStatusTextResetContent                                 = "RESET_CONTENT"
	HttpStatusTextPartialContent                               = "PARTIAL_CONTENT"
	HttpStatusTextMultiStatus                                  = "MULTI_STATUS"
	HttpStatusTextAlreadyReported                              = "ALREADY_REPORTED"
	HttpStatusTextImUsed                                       = "IM_USED"
	HttpStatusTextMultipleChoices                              = "MULTI_CHOICES"
	HttpStatusTextMovedPermanently                             = "MOVED_PERMANENTLY"
	HttpStatusTextFound                                        = "FOUND"
	HttpStatusTextSeeOther                                     = "SEE_OTHER"
	HttpStatusTextNotModified                                  = "NOT_MODIFIED"
	HttpStatusTextUseProxy                                     = "USE_PROXY"
	HttpStatusTextUnused                                       = "UNUSED"
	HttpStatusTextTemporaryRedirect                            = "TEMPORARY_REDIRECT"
	HttpStatusTextPermanentRedirect                            = "PERMANENT_REDIRECT"
	HttpStatusTextBadRequest                                   = "BAD_REQUEST"
	HttpStatusTextUnauthorized                                 = "UNAUTHORIZED"
	HttpStatusTextPaymentRequired                              = "PAYMENT_REQUIRED"
	HttpStatusTextForbidden                                    = "FORBIDDEN"
	HttpStatusTextNotFound                                     = "NOT_FOUND"
	HttpStatusTextMethodNotAllowed                             = "METHOD_NOT_ALLOWED"
	HttpStatusTextNotAcceptable                                = "NOT_ACCEPTABLE"
	HttpStatusTextProxyAuthenticationRequired                  = "PROXY_AUTHENTICATION_REQUIRED"
	HttpStatusTextRequestTimeout                               = "REQUEST_TIMEOUT"
	HttpStatusTextConflict                                     = "CONFLICT"
	HttpStatusTextGone                                         = "GONE"
	HttpStatusTextLengthRequired                               = "LENGTH_REQUIRED"
	HttpStatusTextPreconditionFailed                           = "PRECONDITION_FAILED"
	HttpStatusTextPayloadTooLarge                              = "PAYLOAD_TOO_LARGE"
	HttpStatusTextRequestUriTooLong                            = "REQUEST_URI_TOO_LONG"
	HttpStatusTextUnsupportedMediaType                         = "UNSUPPORTED_MEDIA_TYPE"
	HttpStatusTextRequestedRangeNotSatisfiable                 = "REQUESTED_RANGE_NOT_SATISFIABLE"
	HttpStatusTextExpectationFailed                            = "EXPECTATION_FAILED"
	HttpStatusTextImATeapot                                    = "IM_A_TEAPOT"
	HttpStatusTextInsufficientSpaceOnResource                  = "INSUFFICIENT_SPACE_ON_RESOURCE"
	HttpStatusTextMethodFailure                                = "METHOD_FAILURE"
	HttpStatusTextMisdirectedRequest                           = "MISDIRECTED_REQUEST"
	HttpStatusTextUnprocessableEntity                          = "UNPROCESSABLE_ENTITY"
	HttpStatusTextLocked                                       = "LOCKED"
	HttpStatusTextFailedDependency                             = "FAILED_DEPENDENCY"
	HttpStatusTextUpgradeRequired                              = "UPGRADE_REQUIRED"
	HttpStatusTextPreconditionRequired                         = "PRECONDITION_REQUIRED"
	HttpStatusTextTooManyRequests                              = "TOO_MANY_REQUESTS"
	HttpStatusTextRequestHeaderFieldsTooLarge                  = "REQUEST_HEADER_FIELDS_TOO_LARGE"
	HttpStatusTextUnavailableForLegalReasons                   = "UNAVAILABLE_FOR_LEGAL_REASONS"
	HttpStatusTextInternalServerError                          = "INTERNAL_SERVER_ERROR"
	HttpStatusTextNotImplemented                               = "NOT_IMPLEMENTED"
	HttpStatusTextBadGateway                                   = "BAD_GATEWAY"
	HttpStatusTextServiceUnavailable                           = "SERVICE_UNAVAILABLE"
	HttpStatusTextGatewayTimeout                               = "GATEWAY_TIMEOUT"
	HttpStatusTextHttpVersionNotSupported                      = "HTTP_VERSION_NOT_SUPPORTED"
	HttpStatusTextVariantAlsoNegotiates                        = "VARIANT_ALSO_NEGOTIATES"
	HttpStatusTextInsufficientStorage                          = "INSUFFICIENT_STORAGE"
	HttpStatusTextLoopDetected                                 = "LOOP_DETECTED"
	HttpStatusTextNotExtended                                  = "NOT_EXTENDED"
	HttpStatusTextNetworkAuthenticationRequired                = "NETWORK_AUTHENTICATION_REQUIRED"
)

func GetHttpStatusTextFromCode(code HttpStatusCode) HttpStatusText {
	switch code {
	case HttpStatusCodeContinue:
		return HttpStatusTextContinue
	case HttpStatusCodeSwitchingProtocols:
		return HttpStatusTextSwitchingProtocols
	case HttpStatusCodeProcessing:
		return HttpStatusTextProcessing
	case HttpStatusCodeEarlyHints:
		return HttpStatusTextEarlyHints
	case HttpStatusCodeOk:
		return HttpStatusTextOk
	case HttpStatusCodeCreated:
		return HttpStatusTextCreated
	case HttpStatusCodeAccepted:
		return HttpStatusTextAccepted
	case HttpStatusCodeNonAuthoritativeInformation:
		return HttpStatusTextNonAuthoritativeInformation
	case HttpStatusCodeNoContent:
		return HttpStatusTextNoContent
	case HttpStatusCodeResetContent:
		return HttpStatusTextResetContent
	case HttpStatusCodePartialContent:
		return HttpStatusTextPartialContent
	case HttpStatusCodeMultiStatus:
		return HttpStatusTextMultiStatus
	case HttpStatusCodeAlreadyReported:
		return HttpStatusTextAlreadyReported
	case HttpStatusCodeImUsed:
		return HttpStatusTextImUsed
	case HttpStatusCodeMultipleChoices:
		return HttpStatusTextMultipleChoices
	case HttpStatusCodeMovedPermanently:
		return HttpStatusTextMovedPermanently
	case HttpStatusCodeFound:
		return HttpStatusTextFound
	case HttpStatusCodeSeeOther:
		return HttpStatusTextSeeOther
	case HttpStatusCodeNotModified:
		return HttpStatusTextNotModified
	case HttpStatusCodeUseProxy:
		return HttpStatusTextUseProxy
	case HttpStatusCodeUnused:
		return HttpStatusTextUnused
	case HttpStatusCodeTemporaryRedirect:
		return HttpStatusTextTemporaryRedirect
	case HttpStatusCodePermanentRedirect:
		return HttpStatusTextPermanentRedirect
	case HttpStatusCodeBadRequest:
		return HttpStatusTextBadRequest
	case HttpStatusCodeUnauthorized:
		return HttpStatusTextUnauthorized
	case HttpStatusCodePaymentRequired:
		return HttpStatusTextPaymentRequired
	case HttpStatusCodeForbidden:
		return HttpStatusTextForbidden
	case HttpStatusCodeNotFound:
		return HttpStatusTextNotFound
	case HttpStatusCodeMethodNotAllowed:
		return HttpStatusTextMethodNotAllowed
	case HttpStatusCodeNotAcceptable:
		return HttpStatusTextNotAcceptable
	case HttpStatusCodeProxyAuthenticationRequired:
		return HttpStatusTextProxyAuthenticationRequired
	case HttpStatusCodeRequestTimeout:
		return HttpStatusTextRequestTimeout
	case HttpStatusCodeConflict:
		return HttpStatusTextConflict
	case HttpStatusCodeGone:
		return HttpStatusTextGone
	case HttpStatusCodeLengthRequired:
		return HttpStatusTextLengthRequired
	case HttpStatusCodePreconditionFailed:
		return HttpStatusTextPreconditionFailed
	case HttpStatusCodePayloadTooLarge:
		return HttpStatusTextPayloadTooLarge
	case HttpStatusCodeRequestUriTooLong:
		return HttpStatusTextRequestUriTooLong
	case HttpStatusCodeUnsupportedMediaType:
		return HttpStatusTextUnsupportedMediaType
	case HttpStatusCodeRequestedRangeNotSatisfiable:
		return HttpStatusTextRequestedRangeNotSatisfiable
	case HttpStatusCodeExpectationFailed:
		return HttpStatusTextExpectationFailed
	case HttpStatusCodeImATeapot:
		return HttpStatusTextImATeapot
	case HttpStatusCodeInsufficientSpaceOnResource:
		return HttpStatusTextInsufficientSpaceOnResource
	case HttpStatusCodeMethodFailure:
		return HttpStatusTextMethodFailure
	case HttpStatusCodeMisdirectedRequest:
		return HttpStatusTextMisdirectedRequest
	case HttpStatusCodeUnprocessableEntity:
		return HttpStatusTextUnprocessableEntity
	case HttpStatusCodeLocked:
		return HttpStatusTextLocked
	case HttpStatusCodeFailedDependency:
		return HttpStatusTextFailedDependency
	case HttpStatusCodeUpgradeRequired:
		return HttpStatusTextUpgradeRequired
	case HttpStatusCodePreconditionRequired:
		return HttpStatusTextPreconditionRequired
	case HttpStatusCodeTooManyRequests:
		return HttpStatusTextTooManyRequests
	case HttpStatusCodeRequestHeaderFieldsTooLarge:
		return HttpStatusTextRequestHeaderFieldsTooLarge
	case HttpStatusCodeUnavailableForLegalReasons:
		return HttpStatusTextUnavailableForLegalReasons
	case HttpStatusCodeInternalServerError:
		return HttpStatusTextInternalServerError
	case HttpStatusCodeNotImplemented:
		return HttpStatusTextNotImplemented
	case HttpStatusCodeBadGateway:
		return HttpStatusTextBadGateway
	case HttpStatusCodeServiceUnavailable:
		return HttpStatusTextServiceUnavailable
	case HttpStatusCodeGatewayTimeout:
		return HttpStatusTextGatewayTimeout
	case HttpStatusCodeHttpVersionNotSupported:
		return HttpStatusTextHttpVersionNotSupported
	case HttpStatusCodeVariantAlsoNegotiates:
		return HttpStatusTextVariantAlsoNegotiates
	case HttpStatusCodeInsufficientStorage:
		return HttpStatusTextInsufficientStorage
	case HttpStatusCodeLoopDetected:
		return HttpStatusTextLoopDetected
	case HttpStatusCodeNotExtended:
		return HttpStatusTextNotExtended
	case HttpStatusCodeNetworkAuthenticationRequired:
		return HttpStatusTextNetworkAuthenticationRequired
	default:
		panic("invalid http status code")
	}
}

func GetHttpStatusCodeParseText(text HttpStatusText) HttpStatusCode {
	switch HttpStatusText(ToSnakeCaseUpper(string(text))) {
	case HttpStatusTextContinue:
		return HttpStatusCodeContinue
	case HttpStatusTextSwitchingProtocols:
		return HttpStatusCodeSwitchingProtocols
	case HttpStatusTextProcessing:
		return HttpStatusCodeProcessing
	case HttpStatusTextEarlyHints:
		return HttpStatusCodeEarlyHints
	case HttpStatusTextOk:
		return HttpStatusCodeOk
	case HttpStatusTextCreated:
		return HttpStatusCodeCreated
	case HttpStatusTextAccepted:
		return HttpStatusCodeAccepted
	case HttpStatusTextNonAuthoritativeInformation:
		return HttpStatusCodeNonAuthoritativeInformation
	case HttpStatusTextNoContent:
		return HttpStatusCodeNoContent
	case HttpStatusTextResetContent:
		return HttpStatusCodeResetContent
	case HttpStatusTextPartialContent:
		return HttpStatusCodePartialContent
	case HttpStatusTextMultiStatus:
		return HttpStatusCodeMultiStatus
	case HttpStatusTextAlreadyReported:
		return HttpStatusCodeAlreadyReported
	case HttpStatusTextImUsed:
		return HttpStatusCodeImUsed
	case HttpStatusTextMultipleChoices:
		return HttpStatusCodeMultipleChoices
	case HttpStatusTextMovedPermanently:
		return HttpStatusCodeMovedPermanently
	case HttpStatusTextFound:
		return HttpStatusCodeFound
	case HttpStatusTextSeeOther:
		return HttpStatusCodeSeeOther
	case HttpStatusTextNotModified:
		return HttpStatusCodeNotModified
	case HttpStatusTextUseProxy:
		return HttpStatusCodeUseProxy
	case HttpStatusTextUnused:
		return HttpStatusCodeUnused
	case HttpStatusTextTemporaryRedirect:
		return HttpStatusCodeTemporaryRedirect
	case HttpStatusTextPermanentRedirect:
		return HttpStatusCodePermanentRedirect
	case HttpStatusTextBadRequest:
		return HttpStatusCodeBadRequest
	case HttpStatusTextUnauthorized:
		return HttpStatusCodeUnauthorized
	case HttpStatusTextPaymentRequired:
		return HttpStatusCodePaymentRequired
	case HttpStatusTextForbidden:
		return HttpStatusCodeForbidden
	case HttpStatusTextNotFound:
		return HttpStatusCodeNotFound
	case HttpStatusTextMethodNotAllowed:
		return HttpStatusCodeMethodNotAllowed
	case HttpStatusTextNotAcceptable:
		return HttpStatusCodeNotAcceptable
	case HttpStatusTextProxyAuthenticationRequired:
		return HttpStatusCodeProxyAuthenticationRequired
	case HttpStatusTextRequestTimeout:
		return HttpStatusCodeRequestTimeout
	case HttpStatusTextConflict:
		return HttpStatusCodeConflict
	case HttpStatusTextGone:
		return HttpStatusCodeGone
	case HttpStatusTextLengthRequired:
		return HttpStatusCodeLengthRequired
	case HttpStatusTextPreconditionFailed:
		return HttpStatusCodePreconditionFailed
	case HttpStatusTextPayloadTooLarge:
		return HttpStatusCodePayloadTooLarge
	case HttpStatusTextRequestUriTooLong:
		return HttpStatusCodeRequestUriTooLong
	case HttpStatusTextUnsupportedMediaType:
		return HttpStatusCodeUnsupportedMediaType
	case HttpStatusTextRequestedRangeNotSatisfiable:
		return HttpStatusCodeRequestedRangeNotSatisfiable
	case HttpStatusTextExpectationFailed:
		return HttpStatusCodeExpectationFailed
	case HttpStatusTextImATeapot:
		return HttpStatusCodeImATeapot
	case HttpStatusTextInsufficientSpaceOnResource:
		return HttpStatusCodeInsufficientSpaceOnResource
	case HttpStatusTextMethodFailure:
		return HttpStatusCodeMethodFailure
	case HttpStatusTextMisdirectedRequest:
		return HttpStatusCodeMisdirectedRequest
	case HttpStatusTextUnprocessableEntity:
		return HttpStatusCodeUnprocessableEntity
	case HttpStatusTextLocked:
		return HttpStatusCodeLocked
	case HttpStatusTextFailedDependency:
		return HttpStatusCodeFailedDependency
	case HttpStatusTextUpgradeRequired:
		return HttpStatusCodeUpgradeRequired
	case HttpStatusTextPreconditionRequired:
		return HttpStatusCodePreconditionRequired
	case HttpStatusTextTooManyRequests:
		return HttpStatusCodeTooManyRequests
	case HttpStatusTextRequestHeaderFieldsTooLarge:
		return HttpStatusCodeRequestHeaderFieldsTooLarge
	case HttpStatusTextUnavailableForLegalReasons:
		return HttpStatusCodeUnavailableForLegalReasons
	case HttpStatusTextInternalServerError:
		return HttpStatusCodeInternalServerError
	case HttpStatusTextNotImplemented:
		return HttpStatusCodeNotImplemented
	case HttpStatusTextBadGateway:
		return HttpStatusCodeBadGateway
	case HttpStatusTextServiceUnavailable:
		return HttpStatusCodeServiceUnavailable
	case HttpStatusTextGatewayTimeout:
		return HttpStatusCodeGatewayTimeout
	case HttpStatusTextHttpVersionNotSupported:
		return HttpStatusCodeHttpVersionNotSupported
	case HttpStatusTextVariantAlsoNegotiates:
		return HttpStatusCodeVariantAlsoNegotiates
	case HttpStatusTextInsufficientStorage:
		return HttpStatusCodeInsufficientStorage
	case HttpStatusTextLoopDetected:
		return HttpStatusCodeLoopDetected
	case HttpStatusTextNotExtended:
		return HttpStatusCodeNotExtended
	case HttpStatusTextNetworkAuthenticationRequired:
		return HttpStatusCodeNetworkAuthenticationRequired
	default:
		panic("invalid http status text")
	}
}
