# Polimorfismo y Interfaces en Go

## Pregunta 1: ¿Qué se guarda en `messenger` del Service?

### Pregunta :
> Entiendo que el servicio "service.go" puede recibir en su constructor NewNotificationService(messenger Messenger) un tipo de la interfaz y emailMessenger que esta en main.go es del tipo de la interfaz por que contiene una implementacion de Messenger que es la interfaz, pero lo que no entiendo es que se guarda en:
> ```go
> return &Service{
>     messenger: messenger,
> }
> ```
> 
> Es decir que pasa con `fromAddress: fromAddress,` que se guarda en messenger que es la clave del diccionaro por llamarlo asi de:
> ```go
> return &Service{
>     messenger: messenger,
> }
> ```

### Respuesta:

Lo que se guarda en `messenger` del struct `Service` es:

El **PUNTERO COMPLETO** al objeto `EmailMessenger` (o `WhatsAppMessenger`), **incluyendo todos sus datos** (como `fromAddress`).

#### Paso a paso:

```go
// En main.go
emailMessenger := messaging.NewEmailMessenger("soporte@miempresa.cl")
// emailMessenger es un *EmailMessenger que apunta a un objeto con:
// { fromAddress: "soporte@miempresa.cl" }

// Luego lo pasas al constructor
emailService := notification.NewNotificationService(emailMessenger)
```

#### Dentro del constructor:

```go
func NewNotificationService(messenger Messenger) *Service {
    return &Service{
        messenger: messenger,  // Guarda el PUNTERO completo
    }
}
```

#### Lo que contiene `Service.messenger`:

```
Service {
    messenger: *EmailMessenger {  // ← El PUNTERO al objeto completo
        fromAddress: "soporte@miempresa.cl"  // ← Los datos están aquí
    }
}
```

#### Cuando llamas a `Send`:

```go
service.messenger.Send(...)
// ↑ usa el puntero almacenado
// ↑ que tiene acceso a fromAddress
```

**Resumen:** El campo `messenger` del struct `Service` guarda **una referencia (puntero)** al objeto **completo** de `EmailMessenger`, que incluye **todos sus campos internos** como `fromAddress`. No solo guarda la interfaz, guarda el objeto real con sus datos. 🎯

---

## Pregunta 2: ¿Por qué el constructor acepta objetos con campos adicionales?

### Pregunta:
> Por que el constructor del servicio acepta que se pase por parametro otro campo que no es de el se que se implementa la interfaz pero no es lo unico que viene sino tambien viene fromAddress

### Respuesta:

Esto es **polimorfismo** y la clave del **Dependency Inversion Principle (DIP)**.

#### El constructor del servicio NO le importa qué campos extras tenga el objeto:

```go
func NewNotificationService(messenger Messenger) *Service {
    // ↑ Solo pide algo que cumpla con Messenger
    // NO le importa si tiene fromAddress, fromNumber, o 100 campos más
}
```

#### ¿Por qué acepta `EmailMessenger` si tiene `fromAddress`?

Porque Go usa **Duck Typing**:

> "Si camina como pato y hace cuac como pato, entonces ES un pato"

**Traducido a tu código:**

> "Si tiene el método `Send(string, string) error`, entonces ES un Messenger"

#### Lo que pasa:

```go
// EmailMessenger tiene:
type EmailMessenger struct {
    fromAddress string  // ← El servicio NO sabe que esto existe
}

func (e *EmailMessenger) Send(...) error {
    // ← El servicio SOLO sabe que este método existe
}
```

#### El contrato (interfaz) solo exige:

```go
type Messenger interface {
    Send(destination string, message string) error
    // ← Solo necesita este método
}
```

#### El servicio solo ve:

```
messenger Messenger
    ↓
"Tengo algo que puede hacer Send()"

NO sabe:
- Si es Email o WhatsApp
- Si tiene fromAddress o fromNumber
- Cualquier otro detalle interno
```

#### La magia:

Cuando llamas `service.messenger.Send()`, Go **automáticamente** usa el método correcto del objeto real (`EmailMessenger` o `WhatsAppMessenger`) y ese método **sí tiene acceso** a sus propios campos internos.

**Resumen:** El constructor acepta cualquier cosa que **implemente la interfaz**, sin importar qué campos extras tenga. Solo le importa que pueda ejecutar `Send()`. 🎯

---

## Conceptos Clave

### 1. Polimorfismo
El `Service` puede trabajar con **cualquier** implementación de `Messenger` sin saber los detalles internos.

### 2. Duck Typing
Go no requiere declarar explícitamente `implements Messenger`. Si el tipo tiene los métodos correctos, automáticamente implementa la interfaz.

### 3. Dependency Inversion Principle (DIP)
El `Service` depende de la **abstracción** (`Messenger`), no de las implementaciones concretas (`EmailMessenger`, `WhatsAppMessenger`).

### 4. Encapsulación
Los campos privados como `fromAddress` están ocultos para el `Service`, pero accesibles dentro del método `Send` del `EmailMessenger`.
