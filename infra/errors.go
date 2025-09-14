package infra

import "errors"

var (
	ErrConnectionFailed    = errors.New("connection failed")
	ErrOrmConnectionFailed = errors.New("orm connection failed")
	ErrRecordNotFound      = errors.New("url not found")
	ErrQueryFailed         = errors.New("query execute failed")
	ErrTransactionFailed   = errors.New("transaction failed")
	ErrCommitFailed        = errors.New("commit failed")
	ErrMigrateFailed       = errors.New("migrate failed")
)
