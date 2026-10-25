package store

// FeatureFlagKey represents a key for a Diamond Planner feature flag.
type FeatureFlagKey string

const (
    // FlagSentryEnabled controls if errors are forwarded to an external error tracking service
    // (in addition to Pocketbase logging).
	FlagSentryEnabled FeatureFlagKey = "SENTRY_ENABLED"
)