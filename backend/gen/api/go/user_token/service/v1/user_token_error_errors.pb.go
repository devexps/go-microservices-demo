// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/devexps/go-micro/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the micro package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsTokenNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserTokenErrorReason_TOKEN_NOT_FOUND.String() && e.Code == 401
}

func ErrorTokenNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(401, UserTokenErrorReason_TOKEN_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsCreateAccessTokenFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserTokenErrorReason_CREATE_ACCESS_TOKEN_FAILED.String() && e.Code == 402
}

func ErrorCreateAccessTokenFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(402, UserTokenErrorReason_CREATE_ACCESS_TOKEN_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsStoreRedisFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserTokenErrorReason_STORE_REDIS_FAILED.String() && e.Code == 403
}

func ErrorStoreRedisFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(403, UserTokenErrorReason_STORE_REDIS_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsCreateRefreshTokenFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserTokenErrorReason_CREATE_REFRESH_TOKEN_FAILED.String() && e.Code == 404
}

func ErrorCreateRefreshTokenFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(404, UserTokenErrorReason_CREATE_REFRESH_TOKEN_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsRemoveAccessTokenFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserTokenErrorReason_REMOVE_ACCESS_TOKEN_FAILED.String() && e.Code == 405
}

func ErrorRemoveAccessTokenFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(405, UserTokenErrorReason_REMOVE_ACCESS_TOKEN_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsRemoveRefreshTokenFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserTokenErrorReason_REMOVE_REFRESH_TOKEN_FAILED.String() && e.Code == 406
}

func ErrorRemoveRefreshTokenFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(406, UserTokenErrorReason_REMOVE_REFRESH_TOKEN_FAILED.String(), fmt.Sprintf(format, args...))
}
