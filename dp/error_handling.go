package dp

import "github.com/getsentry/sentry-go"

type ErrorContext struct {
	Key    string
	Values map[string]any
}

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
