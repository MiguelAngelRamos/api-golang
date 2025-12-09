package messaging

import "fmt"

// WhatsAppMessenger es otra implementación concreta del contrato Messenger.
// IMPORTANTE: Al igual que EmailMessenger, este struct solo contiene DATOS.
type WhatsAppMessenger struct {
	fromNumber string // Solo un campo de datos
}

// 🏭 CONSTRUCTOR
// Devuelve *WhatsAppMessenger (un PUNTERO), no el valor completo.
// Esto es crucial para que funcione con la interfaz Messenger.
func NewWhatsAppMessenger(fromNumber string) *WhatsAppMessenger {
	// & crea un puntero: "dame la dirección de memoria de este struct"
	return &WhatsAppMessenger{fromNumber: fromNumber}
}

// 🎯 MÉTODO CON RECEPTOR
// (messenger *WhatsAppMessenger) es el RECEPTOR que vincula Send a *WhatsAppMessenger.
//
// NOTA IMPORTANTE: Aunque WhatsAppMessenger solo tiene el campo "fromNumber",
// este método Send está "pegado" al tipo mediante el RECEPTOR.
//
// Gracias a esta línea, el compilador sabe que:
// *WhatsAppMessenger tiene el método Send(string, string) error
// Por lo tanto, *WhatsAppMessenger IMPLEMENTA Messenger (implícitamente)
func (messenger *WhatsAppMessenger) Send(destination string, message string) error {
	// "messenger" es el equivalente a "this" en Java
	// Accedemos a messenger.fromNumber gracias al receptor
	fmt.Printf("[WHATSAPP] De: %s → Para: %s | Mensaje: %s\n",
		messenger.fromNumber, destination, message)
	return nil
}

// 📝 PATRÓN REPETIDO:
// ===================
// 1. Struct con solo DATOS (fromNumber)
// 2. Constructor que devuelve PUNTERO (*WhatsAppMessenger) usando &
// 3. Método vinculado mediante RECEPTOR (messenger *WhatsAppMessenger)
// 4. Implementa Messenger IMPLÍCITAMENTE (sin declararlo explícitamente)
