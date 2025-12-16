package notification

import "github.com/stretchr/testify/mock"

type MockMessenger struct {
	mock.Mock
}

/*
mock.Mock es un struct que viene Testify

Nos permite
-  Hacer llamadas (called)
-  Configurar retornos
-  Verificar expectativas con On y con el hacer Afirmaciones (asserts)
*/

func (mockStruct *MockMessenger) Send(destination string, message string) error {
	// argumentos
	args := mockStruct.Called(destination, message)
	/*
		args := mockStruct.Called(destination, message)
		En este momento:
		Testify registra: "Se llamo Send con {'user@example.com, "Hola desde Testing"}"
		Testify busca en las expectativas configuradas en con .On()
		- Encuentra .On("Send", "user@testing.com", "Hola desde Testing").Return(nil)
		- args ahora contiene un: mock.Arguments{nil}
	*/

	/*
		mockStruct.Called() ?
		1. Registra la llamada al metodo "Send" con los argumentos recibidos

		2. Busca expectativas configuradas con el .On()

		3. Devuelve un objeto mock.Arguments con los valores configurados con .Return()

	*/

	/*
	 Ejemplo de lo que hace "mockStruct.Called"
	  args = mock.Arguments{nil}
	*/
	return args.Error(0)
	/*
		- args.Error(0)  accede a la posicion 0 del slice
		  args[0] = nil
			Convierte a nil a tipo error
			Devuelve nil (nil significa que no hay error)
	*/
}
