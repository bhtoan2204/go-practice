package response

const (
	Success             = 2000
	Created             = 2010
	NoContent           = 2040
	ErrorUnauthorized   = 4010
	ErrorInvalidParams  = 4001
	ErrorNotFound       = 4040
	ErrorBadRequest     = 4000
	ErrorInternalServer = 5000
)

var msg = map[int]string{
	Success:             "Success",
	Created:             "Created",
	NoContent:           "No Content",
	ErrorUnauthorized:   "Unauthorized",
	ErrorInvalidParams:  "Invalid Parameters",
	ErrorNotFound:       "Not Found",
	ErrorBadRequest:     "Bad Request",
	ErrorInternalServer: "Internal Server Error",
}
