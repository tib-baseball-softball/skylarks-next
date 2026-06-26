package dp

import (
	"github.com/getsentry/sentry-go"
	"github.com/pocketbase/pocketbase/core"
)

type ErrorContext struct {
	Key    string
	Values map[string]any
}

// LogErrorInternalExternal logs the error to the internal PocketBase logger 
// and forwards it to the external error tracking system.
func LogErrorInternalExternal(app core.App, err error, context *ErrorContext, tags *map[string]string) {
	app.Logger().Error(err.Error(), "context", context, "tags", tags)
	ForwardErrorToExternalSystem(err, context, tags)
}

// Forwards the error to the external error tracking system 
// (Sentry-compatible, using self-hosted Bugsink instance).
func ForwardErrorToExternalSystem(err error, context *ErrorContext, tags *map[string]string) {
	sentry.ConfigureScope(func(scope *sentry.Scope) {
		if context != nil {
			scope.SetContext(context.Key, context.Values)
		}

		if tags != nil {
			scope.SetTags(*tags)
		}
		sentry.CaptureException(err)
	})
}
