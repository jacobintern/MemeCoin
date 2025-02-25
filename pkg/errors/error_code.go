package errors

var (
	// 500 server error
	InternalServerPanic  = NewCustomError("500001", StatusInternalServerError, "internal server panic.")
	InternalServerError  = NewCustomError("500002", StatusInternalServerError, "internal server error.")
	MethodNotImplement   = NewCustomError("500003", StatusInternalServerError, "method not implement.")
	ReadConfigFatal      = NewCustomError("500004", StatusInternalServerError, "read config failed.")
	UnmarshalConfigFatal = NewCustomError("500005", StatusInternalServerError, "unmarshal config struct failed.")
	ValidConfigFatal     = NewCustomError("500006", StatusInternalServerError, "valid config struct failed.")

	// 504
	GatewayTimeout = NewCustomError("504001", StatusGatewayTimeout, "request time out.")

	// 400 bad request
	InvalidArgument = NewCustomError("400001", StatusBadRequest, "invalid arguments.")
	InvalidParams   = NewCustomError("400002", StatusBadRequest, "parse parameter failed.")
	IDParseFailed   = NewCustomError("400003", StatusBadRequest, "id parse failed.")

	// 404 not found
	RouteNotFound    = NewCustomError("404001", StatusNotFound, "route not found.")
	MemeCoinNotFound = NewCustomError("404002", StatusNotFound, "meme coin not found.")
)
