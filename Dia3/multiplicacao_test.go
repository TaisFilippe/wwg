package multiplicacao

import "testing"

func TestMultiplique(t *testing.T) {

	// Chamar a função que multiplica
	obtive := Multiplique(7, 5)

	// Dizer o que esperamos que aconteça
	espero := 35

	// Comparar o que recebemos com o que esperavamos receber
	if obtive != espero {
		t.Errorf("espero '%d'; mas obtive '%d'", espero, obtive)

		//Sinalise que o teste falahou
	}
}
