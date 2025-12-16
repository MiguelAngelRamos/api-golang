# 🔍 Métodos de Verificación en Testify Mock

## 📚 Índice

1. [¿Qué hace AssertExpectations?](#assert-expectations)
2. [Todos los Métodos de Verificación](#metodos-verificacion)
3. [Ejemplos Prácticos](#ejemplos-practicos)
4. [Comparación con Java/Mockito](#comparacion-java)
5. [Casos de Uso Comunes](#casos-uso)
6. [Mejores Prácticas](#mejores-practicas)
7. [Tabla de Referencia Rápida](#tabla-referencia)

---

## 🎯 ¿Qué hace AssertExpectations? {#assert-expectations}

### Definición

```go
mockMessenger.AssertExpectations(t)
```

**`AssertExpectations`** verifica que **TODAS** las expectativas configuradas con `.On()` fueron llamadas exactamente como se esperaba.

### ¿Qué verifica?

1. ✅ **Que el método fue llamado** (si configuraste `.On("Send", ...)`)
2. ✅ **Con los argumentos correctos** (los mismos que especificaste)
3. ✅ **El número correcto de veces** (por defecto: al menos 1 vez)

### Ejemplo Básico

```go
// ARRANGE: Configurar expectativa
mockMessenger := new(MockMessenger)
mockMessenger.On("Send", "user@test.com", "Hello").Return(nil)
//            ^^^
//            Expectativa registrada

// ACT: Usar el mock
service.Notify("user@test.com", "Hello")
// Internamente llama: mockMessenger.Send("user@test.com", "Hello")

// ASSERT: Verificar que la expectativa se cumplió
mockMessenger.AssertExpectations(t)
// ✅ Verifica: ¿Se llamó Send con ("user@test.com", "Hello")? → SÍ
```

### ¿Qué pasa si falla?

Si el método **NO** fue llamado o fue llamado con **argumentos diferentes**:

```go
mockMessenger.On("Send", "user@test.com", "Hello").Return(nil)

// ACT: Llamar con argumentos DIFERENTES
service.Notify("otro@test.com", "Bye")  // ❌ Argumentos no coinciden

// ASSERT: Esto FALLA
mockMessenger.AssertExpectations(t)
// ❌ Error: FAIL: Send("user@test.com", "Hello") was not called
```

---

## 📋 Todos los Métodos de Verificación {#metodos-verificacion}

### 1. **AssertExpectations(t)** - Verificar TODAS las expectativas

```go
func (m *Mock) AssertExpectations(t TestingT) bool
```

**Propósito:** Verifica que todas las llamadas configuradas con `.On()` se cumplieron.

**Cuándo usar:** Al final de cada test para asegurar que el mock fue usado como esperabas.

```go
mockMessenger.On("Send", "dest", "msg").Return(nil)
// ... código que usa el mock ...
mockMessenger.AssertExpectations(t)  // ✅ Verifica TODAS las expectativas
```

**Equivalente en Mockito (Java):**
```java
verify(mock).method(args);
```

---

### 2. **AssertCalled(t, method, args...)** - Verificar llamada específica

```go
func (m *Mock) AssertCalled(t TestingT, methodName string, arguments ...interface{}) bool
```

**Propósito:** Verifica que un método **fue llamado** con argumentos específicos (al menos 1 vez).

**Diferencia con AssertExpectations:** No necesitas configurar `.On()` previamente.

```go
// NO necesitas .On() para usar AssertCalled
mockMessenger := new(MockMessenger)
mockMessenger.On("Send", mock.Anything, mock.Anything).Return(nil)

service.Notify("user@test.com", "Hello")

// Verificar que fue llamado con estos argumentos exactos
mockMessenger.AssertCalled(t, "Send", "user@test.com", "Hello")  // ✅
```

**Ejemplo práctico del proyecto:**

```go
func TestNotify_ValidInput_Success(t *testing.T) {
    mockMessenger := new(MockMessenger)
    mockMessenger.On("Send", "user@example.com", "Hello World").Return(nil)
    service := NewNotificationService(mockMessenger)
    
    service.Notify("user@example.com", "Hello World")
    
    // Opción 1: Verificar todas las expectativas
    mockMessenger.AssertExpectations(t)
    
    // Opción 2: Verificar llamada específica (redundante aquí, pero posible)
    mockMessenger.AssertCalled(t, "Send", "user@example.com", "Hello World")
}
```

---

### 3. **AssertNotCalled(t, method)** - Verificar que NO fue llamado

```go
func (m *Mock) AssertNotCalled(t TestingT, methodName string, arguments ...interface{}) bool
```

**Propósito:** Verifica que un método **NO fue llamado** (con argumentos específicos o en absoluto).

**Cuándo usar:** Cuando la lógica debe prevenir una llamada (por ejemplo, validaciones que fallan).

```go
func TestNotify_EmptyMessage_ReturnsError(t *testing.T) {
    mockMessenger := new(MockMessenger)
    service := NewNotificationService(mockMessenger)
    
    // ACT: Mensaje vacío → validación falla
    err := service.Notify("user@example.com", "")
    
    // ASSERT: El messenger NO debe ser llamado
    mockMessenger.AssertNotCalled(t, "Send")  // ✅ Verifica que Send nunca fue llamado
}
```

**Con argumentos específicos:**

```go
// Verificar que NO fue llamado con estos argumentos específicos
mockMessenger.AssertNotCalled(t, "Send", "blocked@example.com", mock.Anything)
```

---

### 4. **AssertNumberOfCalls(t, method, n)** - Verificar número exacto de llamadas

```go
func (m *Mock) AssertNumberOfCalls(t TestingT, methodName string, expectedCalls int) bool
```

**Propósito:** Verifica que un método fue llamado **exactamente N veces**.

**Cuándo usar:** Cuando importa el número de veces que se llama un método.

```go
func TestNotify_MultipleMessages(t *testing.T) {
    mockMessenger := new(MockMessenger)
    mockMessenger.On("Send", mock.Anything, mock.Anything).Return(nil)
    service := NewNotificationService(mockMessenger)
    
    // ACT: Enviar 3 notificaciones
    service.Notify("user1@test.com", "Msg 1")
    service.Notify("user2@test.com", "Msg 2")
    service.Notify("user3@test.com", "Msg 3")
    
    // ASSERT: Verificar que fue llamado exactamente 3 veces
    mockMessenger.AssertNumberOfCalls(t, "Send", 3)  // ✅
}
```

**Ejemplo del proyecto:**

```go
func TestNotify_WithAnything_AcceptsAnyArgument(t *testing.T) {
    mockMessenger := new(MockMessenger)
    mockMessenger.On("Send", mock.Anything, mock.Anything).Return(nil)
    service := NewNotificationService(mockMessenger)
    
    service.Notify("test1@example.com", "Message 1")
    service.Notify("test2@example.com", "Message 2")
    
    // Verificar que fue llamado exactamente 2 veces
    mockMessenger.AssertNumberOfCalls(t, "Send", 2)  // ✅
}
```

---

### 5. **MethodCalled(method, args...)** - Registrar llamada manual (avanzado)

```go
func (m *Mock) MethodCalled(methodName string, arguments ...interface{}) Arguments
```

**Propósito:** Registra una llamada manualmente (usado internamente por `.Called()`).

**Nota:** Raramente lo usarás directamente; es parte de la implementación interna.

```go
// Dentro del método mock
func (m *MockMessenger) Send(dest, msg string) error {
    args := m.Called(dest, msg)  // ← Internamente usa MethodCalled
    return args.Error(0)
}
```

---

### 6. **On(method, args...)** - Configurar expectativa

```go
func (m *Mock) On(methodName string, arguments ...interface{}) *Call
```

**Propósito:** Configurar el comportamiento esperado del mock.

**No es un método de verificación**, pero es esencial para configurar expectativas.

```go
// Configurar expectativa
mockMessenger.On("Send", "dest", "msg").Return(nil)
//            ^^
//            Método de CONFIGURACIÓN
```

**Métodos encadenables de `.On()`:**

```go
mockMessenger.On("Send", "dest", "msg").
    Return(nil).              // Valor de retorno
    Once()                    // Solo se puede llamar 1 vez

mockMessenger.On("Send", "dest", "msg").
    Return(nil).
    Times(3)                  // Debe ser llamado exactamente 3 veces

mockMessenger.On("Send", "dest", "msg").
    Return(nil).
    Maybe()                   // Opcional: puede no ser llamado
```

---

### 7. **Once(), Times(n), Maybe()** - Configurar número de llamadas esperadas

```go
func (c *Call) Once() *Call
func (c *Call) Times(i int) *Call
func (c *Call) Maybe() *Call
```

**Propósito:** Especificar cuántas veces debe ser llamado el método.

#### **Once()** - Exactamente 1 vez

```go
mockMessenger.On("Send", "dest", "msg").Return(nil).Once()

service.Notify("dest", "msg")  // ✅ Primera llamada OK
service.Notify("dest", "msg")  // ❌ Segunda llamada → Falla
```

#### **Times(n)** - Exactamente N veces

```go
mockMessenger.On("Send", "dest", "msg").Return(nil).Times(3)

service.Notify("dest", "msg")  // 1ª llamada
service.Notify("dest", "msg")  // 2ª llamada
service.Notify("dest", "msg")  // 3ª llamada ✅
service.Notify("dest", "msg")  // 4ª llamada ❌ Falla (esperaba solo 3)
```

#### **Maybe()** - Opcional (0 o más veces)

```go
mockMessenger.On("Send", "dest", "msg").Return(nil).Maybe()

// Si NO se llama, no falla
mockMessenger.AssertExpectations(t)  // ✅ Pasa aunque no se llamó
```

---

### 8. **Run(func)** - Ejecutar lógica personalizada

```go
func (c *Call) Run(fn func(args Arguments)) *Call
```

**Propósito:** Ejecutar código personalizado cuando se llama el método mock.

**Cuándo usar:** Para simular efectos secundarios o lógica compleja.

```go
var capturedMessage string

mockMessenger.On("Send", mock.Anything, mock.Anything).
    Run(func(args mock.Arguments) {
        // Capturar el mensaje enviado
        capturedMessage = args.String(1)
        fmt.Println("Enviando:", capturedMessage)
    }).
    Return(nil)

service.Notify("user@test.com", "Hello")

// Ahora capturedMessage contiene "Hello"
assert.Equal(t, "Hello", capturedMessage)
```

**Ejemplo avanzado: Simular delay**

```go
mockMessenger.On("Send", mock.Anything, mock.Anything).
    Run(func(args mock.Arguments) {
        time.Sleep(100 * time.Millisecond)  // Simular latencia
    }).
    Return(nil)
```

---

### 9. **WaitUntil(chan)** - Esperar sincronización (avanzado)

```go
func (c *Call) WaitUntil(w <-chan time.Time) *Call
```

**Propósito:** Sincronizar el mock con otros eventos asíncronos.

**Cuándo usar:** Tests de concurrencia o asincronía.

```go
waitChan := make(chan time.Time)

mockMessenger.On("Send", "dest", "msg").
    WaitUntil(waitChan).
    Return(nil)

// El mock esperará hasta que se envíe algo al canal
go func() {
    time.Sleep(1 * time.Second)
    waitChan <- time.Now()  // Desbloquea el mock
}()

service.Notify("dest", "msg")  // Espera hasta que waitChan recibe un valor
```

---

### 10. **After(duration)** - Simular delay

```go
func (c *Call) After(d time.Duration) *Call
```

**Propósito:** Simular latencia antes de devolver el valor.

```go
mockMessenger.On("Send", "dest", "msg").
    Return(nil).
    After(500 * time.Millisecond)  // Espera 500ms antes de devolver

start := time.Now()
service.Notify("dest", "msg")
elapsed := time.Since(start)

assert.True(t, elapsed >= 500*time.Millisecond)  // ✅ Pasaron al menos 500ms
```

---

## 💻 Ejemplos Prácticos del Proyecto {#ejemplos-practicos}

### Ejemplo 1: Usar `AssertExpectations` (Test de éxito)

**Archivo:** `service_test.go`

```go
func TestNotify_ValidInput_Success(t *testing.T) {
    // ARRANGE
    mockMessenger := new(MockMessenger)
    mockMessenger.On("Send", "user@example.com", "Hello World").Return(nil)
    //            ^^^ Expectativa configurada
    service := NewNotificationService(mockMessenger)
    
    // ACT
    err := service.Notify("user@example.com", "Hello World")
    
    // ASSERT
    assert.NoError(t, err)
    mockMessenger.AssertExpectations(t)  // ✅ Verifica que .On() se cumplió
}
```

**¿Qué verifica `AssertExpectations` aquí?**
- ✅ Que `Send` fue llamado
- ✅ Con argumentos `("user@example.com", "Hello World")`
- ✅ Al menos 1 vez (por defecto)

---

### Ejemplo 2: Usar `AssertNotCalled` (Validación falla)

**Archivo:** `service_test.go`

```go
func TestNotify_EmptyDestination_ReturnsError(t *testing.T) {
    // ARRANGE
    mockMessenger := new(MockMessenger)
    service := NewNotificationService(mockMessenger)
    
    // ACT: Destino vacío → validación falla ANTES de llamar al mock
    err := service.Notify("", "mensaje válido")
    
    // ASSERT
    assert.Error(t, err)
    assert.Equal(t, ErrEmptyDestination, err)
    mockMessenger.AssertNotCalled(t, "Send")  // ✅ Verifica que Send NO fue llamado
}
```

**¿Por qué usar `AssertNotCalled` aquí?**
- La validación falla en `service.Notify()` **antes** de llamar a `messenger.Send()`
- Queremos asegurar que el mock **nunca** fue invocado
- Si se llamó, significa que la validación no funcionó ❌

---

### Ejemplo 3: Usar `AssertNumberOfCalls` (Múltiples llamadas)

**Archivo:** `service_test.go`

```go
func TestNotify_WithAnything_AcceptsAnyArgument(t *testing.T) {
    // ARRANGE
    mockMessenger := new(MockMessenger)
    mockMessenger.On("Send", mock.Anything, mock.Anything).Return(nil)
    //                       ^^^^^^^^^^^^^^  ^^^^^^^^^^^^^^
    //                       Acepta CUALQUIER valor
    service := NewNotificationService(mockMessenger)
    
    // ACT: Llamar 2 veces con valores diferentes
    err1 := service.Notify("test1@example.com", "Message 1")
    err2 := service.Notify("test2@example.com", "Message 2")
    
    // ASSERT
    assert.NoError(t, err1)
    assert.NoError(t, err2)
    mockMessenger.AssertNumberOfCalls(t, "Send", 2)  // ✅ Fue llamado exactamente 2 veces
}
```

---

### Ejemplo 4: Usar `Once()` (Restringir número de llamadas)

```go
func TestNotify_OnlyOnce(t *testing.T) {
    // ARRANGE
    mockMessenger := new(MockMessenger)
    mockMessenger.On("Send", "dest", "msg").Return(nil).Once()
    //                                                   ^^^^^^
    //                                                   Solo 1 llamada permitida
    service := NewNotificationService(mockMessenger)
    
    // ACT: Primera llamada
    err1 := service.Notify("dest", "msg")
    assert.NoError(t, err1)  // ✅ OK
    
    // Segunda llamada → Falla
    err2 := service.Notify("dest", "msg")  // ❌ Panic o fallo
}
```

---

### Ejemplo 5: Usar `Run()` para capturar argumentos

```go
func TestNotify_CaptureArguments(t *testing.T) {
    // ARRANGE
    mockMessenger := new(MockMessenger)
    var capturedDest, capturedMsg string
    
    mockMessenger.On("Send", mock.Anything, mock.Anything).
        Run(func(args mock.Arguments) {
            // Capturar argumentos para verificación posterior
            capturedDest = args.String(0)
            capturedMsg = args.String(1)
        }).
        Return(nil)
    
    service := NewNotificationService(mockMessenger)
    
    // ACT
    service.Notify("user@test.com", "Important message")
    
    // ASSERT: Verificar que se capturaron los valores correctos
    assert.Equal(t, "user@test.com", capturedDest)
    assert.Equal(t, "Important message", capturedMsg)
}
```

---

### Ejemplo 6: Usar `Times()` para verificar reintentos

```go
func TestNotify_RetryLogic(t *testing.T) {
    // ARRANGE
    mockMessenger := new(MockMessenger)
    
    // Configurar: debe ser llamado exactamente 3 veces (3 reintentos)
    mockMessenger.On("Send", "dest", "msg").
        Return(errors.New("temporary error")).
        Times(3)
    
    // ACT: Supongamos que el servicio reintenta 3 veces
    for i := 0; i < 3; i++ {
        mockMessenger.Send("dest", "msg")
    }
    
    // ASSERT
    mockMessenger.AssertExpectations(t)  // ✅ Verifica que fue llamado 3 veces
}
```

---

## ☕ Comparación con Java/Mockito {#comparacion-java}

| Funcionalidad | Go (Testify) | Java (Mockito) |
|--------------|--------------|----------------|
| **Configurar mock** | `mock.On("Method", args).Return(value)` | `when(mock.method(args)).thenReturn(value)` |
| **Verificar todas las expectativas** | `mock.AssertExpectations(t)` | `verify(mock).method(args)` (manual por cada método) |
| **Verificar llamada específica** | `mock.AssertCalled(t, "Method", args)` | `verify(mock).method(args)` |
| **Verificar NO llamado** | `mock.AssertNotCalled(t, "Method")` | `verify(mock, never()).method()` |
| **Verificar número de llamadas** | `mock.AssertNumberOfCalls(t, "Method", 3)` | `verify(mock, times(3)).method()` |
| **Solo 1 llamada** | `.On("Method").Return(x).Once()` | `verify(mock, times(1)).method()` |
| **N llamadas** | `.On("Method").Return(x).Times(n)` | `verify(mock, times(n)).method()` |
| **Opcional (0+ veces)** | `.On("Method").Return(x).Maybe()` | `verify(mock, atLeast(0)).method()` |
| **Cualquier argumento** | `mock.Anything` | `any()` / `anyString()` |
| **Ejecutar lógica personalizada** | `.Run(func(args) {...})` | `doAnswer(invocation -> {...})` |
| **Simular delay** | `.After(duration)` | `doAnswer(invocation -> { Thread.sleep(...); })` |

### Ejemplo completo: Go vs Java

**Go (Testify):**
```go
func TestSendEmail(t *testing.T) {
    // Arrange
    mock := new(MockMessenger)
    mock.On("Send", "dest", "msg").Return(nil).Once()
    
    // Act
    err := mock.Send("dest", "msg")
    
    // Assert
    assert.NoError(t, err)
    mock.AssertExpectations(t)
    mock.AssertCalled(t, "Send", "dest", "msg")
    mock.AssertNumberOfCalls(t, "Send", 1)
}
```

**Java (Mockito + JUnit):**
```java
@Test
void testSendEmail() {
    // Arrange
    Messenger mock = Mockito.mock(Messenger.class);
    when(mock.send("dest", "msg")).thenReturn(null);
    
    // Act
    Error err = mock.send("dest", "msg");
    
    // Assert
    assertNull(err);
    verify(mock, times(1)).send("dest", "msg");
}
```

---

## 🎯 Casos de Uso Comunes {#casos-uso}

### Caso 1: Test básico con verificación simple

```go
func TestBasic(t *testing.T) {
    mock := new(MockMessenger)
    mock.On("Send", "dest", "msg").Return(nil)
    
    // Usar mock
    mock.Send("dest", "msg")
    
    // Verificar
    mock.AssertExpectations(t)  // ✅ Más simple
}
```

---

### Caso 2: Verificar que un método NO debe ser llamado

```go
func TestNotCalled(t *testing.T) {
    mock := new(MockMessenger)
    
    // NO configuramos expectativas
    // NO llamamos al mock
    
    // Verificar que NO fue llamado
    mock.AssertNotCalled(t, "Send")  // ✅
}
```

---

### Caso 3: Verificar llamadas múltiples con diferentes argumentos

```go
func TestMultipleCalls(t *testing.T) {
    mock := new(MockMessenger)
    mock.On("Send", "user1@test.com", "Msg1").Return(nil)
    mock.On("Send", "user2@test.com", "Msg2").Return(nil)
    
    mock.Send("user1@test.com", "Msg1")
    mock.Send("user2@test.com", "Msg2")
    
    // Verificar ambas llamadas
    mock.AssertCalled(t, "Send", "user1@test.com", "Msg1")
    mock.AssertCalled(t, "Send", "user2@test.com", "Msg2")
    mock.AssertNumberOfCalls(t, "Send", 2)
}
```

---

### Caso 4: Verificar orden de llamadas (con `InOrder`)

**Nota:** Testify no tiene soporte nativo para verificar orden, pero puedes hacerlo manualmente:

```go
func TestCallOrder(t *testing.T) {
    mock := new(MockMessenger)
    callOrder := []string{}
    
    mock.On("Send", mock.Anything, mock.Anything).
        Run(func(args mock.Arguments) {
            callOrder = append(callOrder, args.String(0))
        }).
        Return(nil)
    
    mock.Send("first@test.com", "msg")
    mock.Send("second@test.com", "msg")
    
    // Verificar orden
    assert.Equal(t, []string{"first@test.com", "second@test.com"}, callOrder)
}
```

**Mockito (Java) tiene `.inOrder()` nativo:**
```java
InOrder inOrder = Mockito.inOrder(mock);
inOrder.verify(mock).send("first@test.com", "msg");
inOrder.verify(mock).send("second@test.com", "msg");
```

---

### Caso 5: Verificar que un método es llamado con cualquier argumento

```go
func TestAnyArgument(t *testing.T) {
    mock := new(MockMessenger)
    mock.On("Send", mock.Anything, mock.Anything).Return(nil)
    
    mock.Send("any@email.com", "any message")
    
    // Verificar con mock.Anything
    mock.AssertCalled(t, "Send", mock.Anything, mock.Anything)  // ✅
    
    // O verificar con valores específicos
    mock.AssertCalled(t, "Send", "any@email.com", "any message")  // ✅
}
```

---

## ✅ Mejores Prácticas {#mejores-practicas}

### 1. **Siempre verificar las expectativas**

```go
// ✅ BIEN
func TestGood(t *testing.T) {
    mock := new(MockMessenger)
    mock.On("Send", "dest", "msg").Return(nil)
    
    mock.Send("dest", "msg")
    
    mock.AssertExpectations(t)  // ✅ Siempre verificar
}

// ❌ MAL
func TestBad(t *testing.T) {
    mock := new(MockMessenger)
    mock.On("Send", "dest", "msg").Return(nil)
    
    mock.Send("dest", "msg")
    
    // ❌ Olvidaste verificar las expectativas
}
```

---

### 2. **Usa `AssertExpectations` para verificación general**

```go
// ✅ BIEN: Usa AssertExpectations para verificar TODAS las expectativas
mock.AssertExpectations(t)

// ⚠️ Redundante: No necesitas ambos si solo hay una expectativa
mock.AssertExpectations(t)
mock.AssertCalled(t, "Send", "dest", "msg")  // Redundante
```

---

### 3. **Usa `AssertNotCalled` para verificar que algo NO pasó**

```go
// ✅ BIEN: Verificar que la validación previene la llamada
func TestValidationPreventsCall(t *testing.T) {
    mock := new(MockMessenger)
    service := NewNotificationService(mock)
    
    service.Notify("", "msg")  // Destino vacío
    
    mock.AssertNotCalled(t, "Send")  // ✅ Send NO debe ser llamado
}
```

---

### 4. **Usa `Once()` o `Times(n)` cuando el número importa**

```go
// ✅ BIEN: Restricción explícita
mock.On("Send", "dest", "msg").Return(nil).Once()

// ❌ MAL: Sin restricción, puede ser llamado infinitas veces
mock.On("Send", "dest", "msg").Return(nil)
```

---

### 5. **Usa `mock.Anything` para argumentos que no importan**

```go
// ✅ BIEN: Cualquier mensaje es válido
mock.On("Send", "dest", mock.Anything).Return(nil)

// ❌ MAL: Demasiado específico si el mensaje no importa
mock.On("Send", "dest", "mensaje exacto").Return(nil)
```

---

### 6. **Usa `Run()` solo cuando necesites lógica personalizada**

```go
// ✅ BIEN: Capturar argumentos para verificación compleja
mock.On("Send", mock.Anything, mock.Anything).
    Run(func(args mock.Arguments) {
        // Lógica personalizada
    }).
    Return(nil)

// ❌ MAL: Run() innecesario si solo devuelves un valor
mock.On("Send", "dest", "msg").
    Run(func(args mock.Arguments) {
        // No haces nada útil aquí
    }).
    Return(nil)  // Podrías haber usado solo Return()
```

---

## 📊 Tabla de Referencia Rápida {#tabla-referencia}

| Método | Propósito | Cuándo Usar |
|--------|-----------|-------------|
| `AssertExpectations(t)` | Verifica TODAS las expectativas configuradas con `.On()` | Al final de cada test |
| `AssertCalled(t, "Method", args)` | Verifica que un método fue llamado con argumentos específicos | Verificación explícita de una llamada |
| `AssertNotCalled(t, "Method")` | Verifica que un método NO fue llamado | Validaciones que previenen llamadas |
| `AssertNumberOfCalls(t, "Method", n)` | Verifica que un método fue llamado exactamente N veces | Cuando el número de llamadas importa |
| `.Once()` | Restricción: solo 1 llamada permitida | Métodos que deben llamarse una sola vez |
| `.Times(n)` | Restricción: exactamente N llamadas | Reintentos o bucles |
| `.Maybe()` | Llamada opcional (0 o más veces) | Métodos que pueden o no ejecutarse |
| `.Run(func)` | Ejecuta lógica personalizada al llamar el método | Capturar argumentos o efectos secundarios |
| `.After(duration)` | Simula latencia antes de devolver | Tests de timeout o performance |
| `mock.Anything` | Acepta cualquier valor para un argumento | Cuando el valor no importa |

---

## 🎓 Ejercicios de Práctica

### Ejercicio 1: Verificar llamada básica

```go
func TestExercise1(t *testing.T) {
    mock := new(MockMessenger)
    mock.On("Send", "test@example.com", "Hello").Return(nil)
    
    mock.Send("test@example.com", "Hello")
    
    // TODO: Verificar que la expectativa se cumplió
    // Respuesta: mock.AssertExpectations(t)
}
```

---

### Ejercicio 2: Verificar que NO fue llamado

```go
func TestExercise2(t *testing.T) {
    mock := new(MockMessenger)
    
    // No llamamos al mock
    
    // TODO: Verificar que Send NO fue llamado
    // Respuesta: mock.AssertNotCalled(t, "Send")
}
```

---

### Ejercicio 3: Verificar número de llamadas

```go
func TestExercise3(t *testing.T) {
    mock := new(MockMessenger)
    mock.On("Send", mock.Anything, mock.Anything).Return(nil)
    
    mock.Send("user1@test.com", "Msg 1")
    mock.Send("user2@test.com", "Msg 2")
    mock.Send("user3@test.com", "Msg 3")
    
    // TODO: Verificar que fue llamado 3 veces
    // Respuesta: mock.AssertNumberOfCalls(t, "Send", 3)
}
```

---

### Ejercicio 4: Capturar argumentos con Run()

```go
func TestExercise4(t *testing.T) {
    mock := new(MockMessenger)
    var capturedMsg string
    
    // TODO: Configurar el mock para capturar el mensaje
    // Respuesta:
    // mock.On("Send", mock.Anything, mock.Anything).
    //     Run(func(args mock.Arguments) {
    //         capturedMsg = args.String(1)
    //     }).
    //     Return(nil)
    
    mock.Send("dest", "Important")
    
    assert.Equal(t, "Important", capturedMsg)
}
```

---

## 📚 Resumen Final

### ¿Cuándo usar cada método?

```
┌─────────────────────────────────────────────────────────────┐
│                  DECISIÓN: ¿Qué verificar?                  │
└────────────────────────┬────────────────────────────────────┘
                         │
         ┌───────────────┼───────────────┐
         │               │               │
         ▼               ▼               ▼
    ¿Todas las    ¿Llamada      ¿NO fue llamado?
    expectativas?  específica?
         │               │               │
         ▼               ▼               ▼
  AssertExpectations  AssertCalled   AssertNotCalled
         │               │               │
         │               │               │
    ¿Configuraste    ¿Sabes los      ¿Validación
    .On()?           argumentos       previene
         │           exactos?         llamada?
         │               │               │
         ▼               ▼               ▼
    ✅ Sí          ✅ Sí            ✅ Sí
                                    
         
    ¿Importa el número de llamadas?
                  │
         ┌────────┴────────┐
         ▼                 ▼
    ¿Exactamente    ¿Al menos
    N veces?        1 vez?
         │              │
         ▼              ▼
  AssertNumberOfCalls  AssertExpectations
   .Times(n)           (por defecto)
```

---

## 🎯 Conclusión

### Los 3 métodos más usados:

1. **`AssertExpectations(t)`** - Usa esto en el 90% de los casos
2. **`AssertNotCalled(t, "Method")`** - Para verificar que algo NO pasó
3. **`AssertNumberOfCalls(t, "Method", n)`** - Cuando el número importa

### Regla de oro:

> **Siempre verifica las expectativas al final de cada test**

```go
// ✅ SIEMPRE haz esto
mock.AssertExpectations(t)
```

---

**Archivo:** `doc/14-metodos-verificacion-mocks.md`  
**Proyecto:** go-clean-api  
**Fecha:** 2025-12-16  
**Versión:** 1.0

