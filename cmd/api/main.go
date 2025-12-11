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

	emailService := notification.NewNotificationService(emailMessenger)

	emailService.Notify("cliente@correo.com", "Mensaje por EMAIL")

}
