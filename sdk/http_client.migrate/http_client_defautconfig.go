package httpclient

import (
	"time"
)

const (
	DefaultLogLevel                  = LogLevelInfo
	DefaultMaxRetryAttempts          = 3
	DefaultEnableDynamicRateLimiting = true
	DefaultMaxConcurrentRequests     = 5
	DefaultTokenBufferPeriod         = 5 * time.Minute
	DefaultTotalRetryDuration        = 5 * time.Minute
	DefaultTimeout                   = 10 * time.Second
)
