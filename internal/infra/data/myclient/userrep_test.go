package myclient

import (
	"testing"
)

//Teste da função de listar todos os usuarios
func Test_ListALLUsers(t *testing.T) {
	list, err := NewUserRepository().ListALLUser()
	if list == nil && err != nil {
		t.Errorf("Erro no Retorno")
	}
}
