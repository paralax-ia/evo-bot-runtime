package errors

import "errors"

var (
	ErrPipelineCancelled   = errors.New("bot_runtime: pipeline cancelled")
	ErrAITimeout           = errors.New("bot_runtime: ai call timeout")
	ErrDispatchInterrupted = errors.New("bot_runtime: dispatch interrupted")
	ErrLockFailed          = errors.New("bot_runtime: distributed lock failed")
	ErrPostbackRejected    = errors.New("bot_runtime: postback rejected by crm (e.g. status 422)")
)
