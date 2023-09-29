package authentication

import "errors"

var ErrUnexpected = errors.New("Unexpected internal error")
var ErrUnAuthorized = errors.New("Unauthorized")
