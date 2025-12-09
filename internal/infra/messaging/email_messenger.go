package messaging

import "fmt"

// EmailMessenger es un detalle concreto que implementa Messenger.
// El dominio NO depende de esto; vive abajo en "infra".
//
// 📦 DATOS (STRUCT): Solo define los campos/atributos que tendrá el tipo.
// Aquí NO vemos métodos dentro de las llaves (diferente a Java/C#).
type EmailMessenger struct {
	fromAddress string // Campo privado (minúscula inicial)
}

// 🏭 CONSTRUCTOR (Factory Function)
// Esta función crea una nueva instancia de EmailMessenger y devuelve un PUNTERO (*).
//
// ¿Por qué devuelve *EmailMessenger y no EmailMessenger?
// - El * significa "puntero a EmailMessenger"
// - En lugar de copiar todos los datos del struct, devolvemos su dirección de memoria
// - Similar a "new EmailMessenger(...)" en Java, pero explícito con &
func NewEmailMessenger(fromAddress string) *EmailMessenger {
	// El operador & obtiene la DIRECCIÓN DE MEMORIA del struct creado
	// & = "dame la dirección donde vive este EmailMessenger en la memoria"
	return &EmailMessenger{
		fromAddress: fromAddress,
	}
}

// 🎯 MÉTODO (con RECEPTOR)
// Send implementa el contrato Messenger.
// Aquí iría la lógica real de envío por SMTP.
//
// ⚡ LA CLAVE ESTÁ AQUÍ: (messenger *EmailMessenger) es el RECEPTOR
//
// Esta sintaxis dice: "Esta función Send PERTENECE al tipo *EmailMessenger"
// No es una función suelta del paquete. Es un MÉTODO del tipo.
//
// Desglose:
// func                           → declara una función
// (messenger *EmailMessenger)    → RECEPTOR: vincula la función al tipo *EmailMessenger
//
//	"messenger" es como "this" en Java (pero explícito)
//	El * indica que el método pertenece al PUNTERO
//
// Send(...)                      → nombre del método y parámetros
// error                          → tipo de retorno
//
// Gracias a este receptor, el compilador reconoce que *EmailMessenger
// tiene el método Send(string, string) error, por lo tanto:
// *EmailMessenger IMPLEMENTA la interfaz Messenger (de forma IMPLÍCITA, sin declararlo)
func (messenger *EmailMessenger) Send(destination string, message string) error {
	// Aquí "messenger" funciona como "this" en Java
	// Podemos acceder a messenger.fromAddress porque el receptor nos da acceso
	fmt.Printf("[EMAIL] De: %s → Para: %s | Mensaje: %s\n",
		messenger.fromAddress, destination, message)
	return nil
}

// 📝 RESUMEN PARA PRINCIPIANTES:
// ================================
// 1. El struct EmailMessenger solo tiene DATOS (fromAddress)
// 2. El método Send se define FUERA, pero se VINCULA mediante el RECEPTOR (messenger *EmailMessenger)
// 3. Gracias al receptor, Send es propiedad exclusiva de *EmailMessenger
// 4. Como *EmailMessenger tiene Send(string, string) error, automáticamente implementa Messenger
// 5. No necesitamos declarar "implements Messenger" como en Java. Go lo detecta solo.
//
// DIFERENCIA VISUAL:
// - Función del paquete:  func Send(...) error { }          ← NO pertenece a nadie
// - Método del tipo:      func (m *EmailMessenger) Send(...) error { }  ← Pertenece a *EmailMessenger
