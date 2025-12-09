package notification

// Messenger es la abstracción (interfaz) que define el comportamiento
// que deben cumplir todos los canales de mensajería.
//
// DIP: El dominio depende SOLO de esta interfaz, NO de implementaciones concretas.
//
// 🔑 CONCEPTO CLAVE - IMPLEMENTACIÓN IMPLÍCITA:
// ===============================================
// En Go, NO necesitas escribir "implements Messenger" como en Java.
// Si un tipo tiene TODOS los métodos que la interfaz declara, automáticamente
// la implementa.
//
// ¿Cómo sabe Go que *EmailMessenger implementa Messenger?
// 1. Messenger requiere: Send(string, string) error
// 2. *EmailMessenger tiene: func (m *EmailMessenger) Send(string, string) error
// 3. ✅ Coinciden → *EmailMessenger implementa Messenger (sin declararlo)
//
// Lo mismo aplica para *WhatsAppMessenger y cualquier otro tipo que tenga
// el método Send con la misma firma.
//
// VENTAJA: Puedes crear nuevos tipos que implementen esta interfaz sin
// modificar el código existente (Open/Closed Principle).
type Messenger interface {
	Send(destination string, message string) error
}

// 📝 EJEMPLO DE VERIFICACIÓN DEL COMPILADOR:
// ===========================================
// Cuando en main.go escribes:
//   emailMessenger := messaging.NewEmailMessenger("...")
//   service := notification.NewNotificationService(emailMessenger)
//
// El compilador hace esto:
// 1. NewNotificationService espera un parámetro de tipo Messenger
// 2. emailMessenger es de tipo *EmailMessenger
// 3. ¿Tiene *EmailMessenger el método Send(string, string) error? → SÍ (gracias al receptor)
// 4. ✅ Pasa la verificación → el código compila
