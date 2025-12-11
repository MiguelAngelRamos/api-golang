package notification

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Test[Metodo]_[Escenario]_[ResultadoEsperado]
*/

func TestNotify_ValidInput_Success(t *testing.T) {
	// Arrange (Preparar)
	mockMessenger := new(MockMessenger)

	// Configurar la expectativa
	// 	mockMessenger.On("metodo", arg1, arg2) retorno
	mockMessenger.On("Send", "user@testing.com", "Hola desde Testing").Return(nil)

	service := NewNotificationService(mockMessenger)
	// Act

	err := service.Notify("user@testing.com", "Hola desde Testing")
	// Asserts
	assert.NoError(t, err)

	mockMessenger.AssertExpectations(t)
}
