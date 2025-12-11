package messaging

import "fmt"

type InstagramMessenger struct {
	userAdress string // Solo un campo de datos
}

func NewInstagramMessenger(userAdress string) *InstagramMessenger {
	messengerStruct := InstagramMessenger{userAdress: userAdress}
	return &messengerStruct
}

func (instagram *InstagramMessenger) Send(destination string, message string) error {
	fmt.Printf("[INSTAGAM] De: %s → Para: %s | Mensaje: %s\n",
		instagram.userAdress, destination, message)
	return nil
}
