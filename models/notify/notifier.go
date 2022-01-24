package notify

type Notifier interface {
	Notify(eventName string, id string, event []byte) error
}
