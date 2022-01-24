package notify

type Notifier interface {
	Notify(eventName string, id string, event []byte) error
}

type Notification struct {
	not Notifier
}

func New(not Notifier) Notification {
	return Notification{not: not}
}

// Notify - notify another system of a change in state
func (s *Notification) Notify(eventName string, id string, event []byte) error {
	s.Notify(eventName, id, event)
	return nil
}
