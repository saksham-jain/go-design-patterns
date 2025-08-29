package email

import (
	"fmt"
	"go-design-patterns/decoractor_pattern/notifier"
)

type Email struct {
	name     string
	notifier *notifier.Notifier
}

func NewEmail(name string, notifier *notifier.Notifier) notifier.Notifier {
	return &Email{
		name:     name,
		notifier: notifier,
	}
}

func (p *Email) Notify() {
	fmt.Printf("notifying via %s\n", p.name)
	if p.notifier != nil {
		(*p.notifier).Notify() // notify futher
	}
}
