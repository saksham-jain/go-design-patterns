package push

import (
	"fmt"
	"go-design-patterns/decoractor_pattern/notifier"
)

type Push struct {
	name     string
	notifier *notifier.Notifier
}

func NewPush(name string, notifier *notifier.Notifier) notifier.Notifier {
	return &Push{
		name:     name,
		notifier: notifier,
	}
}

func (p *Push) Notify() {
	fmt.Printf("notifying via %s\n", p.name)
	if p.notifier != nil {
		(*p.notifier).Notify() // notify futher
	}
}
