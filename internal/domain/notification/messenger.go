package notification

type Messenger interface {
	Send(destination string, message string) error
}
