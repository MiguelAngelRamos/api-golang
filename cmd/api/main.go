// cmd/api/main.go
package main

import (
	"fmt"

	"github.com/MiguelAngelRamos/go-clean-api/internal/domain/notification"
	"github.com/MiguelAngelRamos/go-clean-api/internal/infra/messaging"
)

func main() {
	fmt.Println("=== DIP + Clean Code en Go ===")

	emailMessenger := messaging.NewEmailMessenger("soporte@miempresa.cl")
	whatsAppMessenger := messaging.NewWhatsAppMessenger("+56912345678")
	emailService := notification.NewNotificationService(emailMessenger)
	whatsAppService := notification.NewNotificationService(whatsAppMessenger)
	emailService.Notify("cliente@correo.com", "Mensaje por EMAIL")
	whatsAppService.Notify("+56998765432", "Mensaje por WHATSAPP")

	fmt.Println("\nValidaciones:")
	fmt.Println(emailService.Notify("", "destino vacío")) // error esperado: ErrEmptyDestination
	fmt.Println(emailService.Notify("x@x.com", ""))       // error esperado: ErrEmptyMessage
}
