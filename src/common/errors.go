package common

import "errors"

var InternalError = errors.New("internal error")
var NotFoundError = errors.New("not found")
var LogicError = errors.New("logic error")
var ForbiddenError = errors.New("forbidden error")
