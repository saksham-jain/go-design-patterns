package main

import (
	"go-design-patterns/decoractor_pattern/notifier/email"
	"go-design-patterns/decoractor_pattern/notifier/push"
)

func main() {
	push1 := push.NewPush("PushToApp", nil)               // independent object
	email1 := email.NewEmail("PushToEmailMailChimp", nil) // independent object

	push1.Notify()
	email1.Notify()
	// obove object will only trigger single notification

	decoratedObj := email.NewEmail("PushToEmailMailChimp", &push1) // this object notifier to email as well as push
	decoratedObj.Notify()
}
