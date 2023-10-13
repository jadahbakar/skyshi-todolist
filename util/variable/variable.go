package variable

import "errors"

// errors response
var (
	ErrNotFound       = errors.New("Not Found")             // 404
	ErrBadRequest     = errors.New("Bad Request")           // 400
	ErrInternalServer = errors.New("Internal Server Error") // 500

	Created = 201
	Success = 200
)
