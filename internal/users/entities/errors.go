package entities

import "errors"

var (
	ErrInternalServerError = errors.New("internal Server Error")
	
	ErrNotFound = errors.New("your requested Item is not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("item already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("given Param is not valid")

    ErrBadReqBody=errors.New("validation err")

	ErrEmailExist=errors.New("user email already exist")

	ErrInvalidRequestBody=errors.New("invalid request body")

	ErrDbFailure=errors.New("database failure")

	ErrUserNotfound=errors.New("user not found in database")

) 
