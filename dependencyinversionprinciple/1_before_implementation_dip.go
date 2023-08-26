package dependencyinversionprinciple

// PROBLEM
// The OrderProcessor is tightly coupled with the EmailNotifier.
// If we decide to notify the user through another medium, like SMS or push notifications, this design would require significant changes to OrderProcessor.

type EmailNotifierNotDIP struct {
	// some fields for email configuration
}

func (en *EmailNotifierNotDIP) SendEmail(to string, subject string, body string) {
	// logic to send an email
}

type OrderProcessorNotDIP struct {
	notifier EmailNotifierNotDIP
}

func (op *OrderProcessorNotDIP) CompletePurchase(userEmail string, product string) {
	// logic to complete purchase
	op.notifier.SendEmail(userEmail, "Purchase Completed", "Thank you for purchasing "+product)
}
