package auth

import "github.com/devexps/go-micro/v2/errors"

const (
	reason string = "UNAUTHORIZED"
)

var (
	ErrWrongContext         = errors.Unauthorized(reason, "wrong context for middleware")
	ErrMissingJwtToken      = errors.Unauthorized(reason, "no jwt token in context")
	ErrExtractSubjectFailed = errors.Unauthorized(reason, "extract subject failed")
)
