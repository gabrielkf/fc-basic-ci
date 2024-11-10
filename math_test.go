package main

import "testing"

func TestSoma(t *testing.T) {

	total := soma(15, 15)

	if total != 30 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", total, 30)
	}
}

func TestMult(t *testing.T) {

	actual := mult(3, 4)
	expected := 10

	if actual != expected {
		t.Errorf("Expected %d, received %d", expected, actual)
	}
}
