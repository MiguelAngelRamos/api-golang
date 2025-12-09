package messaging

import "fmt"

type EmailMessenger struct {
	fromAddress string // Campo privado (minúscula inicial)
}

func NewEmailMessenger(fromAddress string) *EmailMessenger {

	return &EmailMessenger{
		fromAddress: fromAddress,
	}
}

// *EmailMessenger IMPLEMENTA la interfaz Messenger (de forma IMPLÍCITA, sin declararlo)
func (messenger *EmailMessenger) Send(destination string, message string) error {

	fmt.Printf("[EMAIL] De: %s → Para: %s | Mensaje: %s\n",
		messenger.fromAddress, destination, message)
	return nil
}
