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
	return args.Error(0)
}
