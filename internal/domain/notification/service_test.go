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
	// El on genera expectativa, cuando se llame a send con estos 2 arguments
	mockMessenger.On("Send", "user@testing.com", "Hola desde Testing").Return(nil)
	/*nil significa que no hay error
	mock.On(..).Return(nil) configuramos que el mock no devuelva un error
	*/

	service := NewNotificationService(mockMessenger)
	// Act

	err := service.Notify("user@testing.com", "Hola desde Testing")
	// Asserts
	assert.NoError(t, err)

	mockMessenger.AssertExpectations(t)
	/*AssertExpectations
	  verifica qeu se llamo a "Send" con {"user@testing.com","Hola desde Testing" }
	*/
}

func TestNotify_EmptyDestination_ReturnsError(t *testing.T) {
	mockMessenger := new(MockMessenger)
	service := NewNotificationService(mockMessenger)
	err := service.Notify("", "mensaje válido")

	/*
		err: valor actual retornado por service.Notify "ErrEmptyDestination"
		t: Objeto de test
	*/
	assert.Error(t, err) /* En esta linea recien nuestro t conoce el error*/

	assert.Equal(t, ErrEmptyDestination, err)
	/*
		ErrEmptyDestination : valor esperado
		err: valor que realmente recibo el error generado la realidad
		testing.T
	*/

	mockMessenger.AssertNotCalled(t, "Send")

}
