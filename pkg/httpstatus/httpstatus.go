package httpstatus

const (
	OK                  = 200
	Created             = 201
	Removed             = 204
	BadRequest          = 400
	Unauthorized        = 401
	Forbidden           = 403
	NotFound            = 404
	InternalServerError = 500
	NotImplemented      = 501
	ServiceUnavailable  = 503
)

var StatusNames = map[int]string{
	OK:                  "OK",
	Created:             "Created",
	Removed:             "Deleted",
	BadRequest:          "Bad Request",
	Unauthorized:        "Unauthorized",
	Forbidden:           "Forbidden",
	NotFound:            "Not Found",
	InternalServerError: "Internal Server Error",
	NotImplemented:      "Not Implemented",
	ServiceUnavailable:  "Service Unavailable",
}
