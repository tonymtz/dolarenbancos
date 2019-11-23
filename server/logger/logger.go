package logger

import (
	"github.com/getsentry/raven-go"
)

var loggingCategory = map[string]string{"category": "logging"}

func init() {
	raven.SetDSN("https://08161331edeb4ebe805a7475a1b3e6bd@sentry.io/1831978")
	raven.SetEnvironment("localhost")
}

func Log(message string) {
	raven.CaptureMessage(message, loggingCategory)
}

func Error(err error) {
	raven.CaptureError(err, nil)
}

func Panic(err error) {
	raven.CapturePanic(func() {
		panic(err)
	}, nil)
}
