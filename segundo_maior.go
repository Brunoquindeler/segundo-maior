package main

func SegundoMaior(numeros []int) int {
	primeiro := 0
	segundo := 0

	for _, numero := range numeros {
		if numero > primeiro {
			segundo = primeiro
			primeiro = numero
		}
		if numero < primeiro && numero > segundo {
			segundo = numero
		}
	}

	return segundo
}
