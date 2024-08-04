package enums

// SubscriptionAccessType represents the access type of a subscription.
// @name SubscriptionAccessType
type SubscriptionAccessType string

const (
	// ManualGranted represents a subscription access manually granted by Educology
	// @description Subscription access manually granted by Educology
	ManualGranted SubscriptionAccessType = "manual_granted"

	// SystemAuto represents a subscription access automatically granted through the payment system
	// @description Subscription access automatically granted through the payment system
	SystemAuto SubscriptionAccessType = "system_auto"
)
