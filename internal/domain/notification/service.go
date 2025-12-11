package notification

type Service struct {
	messenger Messenger
}

func NewNotificationService(messenger Messenger) *Service {
	serviceStruct := Service{messenger: messenger}
	return &serviceStruct
}

func (service *Service) Notify(destination string, message string) error {

	if destination == "" {
		return ErrEmptyDestination
	}

	if message == "" {
		return ErrEmptyMessage
	}

	return service.messenger.Send(destination, message)
}
