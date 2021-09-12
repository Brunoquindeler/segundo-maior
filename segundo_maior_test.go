package main

import "testing"

func TestSegundoMaior(t *testing.T) {

	cenarios := []struct {
		descricao string
		dado      []int
		esperado  int
	}{
		{
			descricao: "Teste 1",
			dado:      []int{2, 5, 1, 6, 3, 8, 4},
			esperado:  6,
		},
		{
			descricao: "Teste 2",
			dado:      []int{8, 1, 6, 3, 2, 4},
			esperado:  6,
		},
	}

	for _, cenario := range cenarios {
		obtido := SegundoMaior(cenario.dado)

		retornaResultado(t, cenario.dado, cenario.esperado, obtido)
	}
}

func retornaResultado(t *testing.T, dado []int, esperado, obtido int) {
	t.Helper()

	if esperado != obtido {
		t.Errorf("\nDado: %v \nEsperado: %d \nObtido: %d", dado, esperado, obtido)
	}
}
