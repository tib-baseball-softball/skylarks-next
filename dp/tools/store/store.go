package store

// FeatureFlagKey represents a key for a Diamond Planner feature flag.
// 
// All feature flags should be saved as boolean values in the app-wide store.
// There is no way to enforce it at the compiler level as Pocketbase defines values as `any`.
type FeatureFlagKey string

const (
    // FlagSentryEnabled controls if errors are forwarded to an external error tracking service
    // (in addition to Pocketbase logging).
	FlagSentryEnabled FeatureFlagKey = "SENTRY_ENABLED"
)