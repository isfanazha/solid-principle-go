package dependencyinversionprinciple

// Solution
// The OrderProcessor depends on the abstraction (Notifier) and not on the concrete implementation (EmailNotifier or SMSNotifier).
// If we decide to change our notification method, or even introduce a new one, we can do so easily without altering OrderProcessor — we simply introduce a new implementation of the Notifier interface.

type Notifier interface {
	Notify(to string, subject string, message string)
}

type EmailNotifier struct {
	// some fields for email configuration
}

func (en *EmailNotifier) Notify(to string, subject string, message string) {
	// logic to send an email
}

type SMSNotifier struct {
	// some fields for SMS configuration
}

func (sn *SMSNotifier) Notify(to string, subject string, message string) {
	// logic to send an SMS (Note: In reality, SMS might not have a "subject", this is just for illustrative purposes)
}

type OrderProcessor struct {
	notifier Notifier // depends on the interface, not concrete implementation
}

func (op *OrderProcessor) CompletePurchase(userContact string, product string) {
	// logic to complete purchase
	op.notifier.Notify(userContact, "Purchase Completed", "Thank you for purchasing "+product)
}
