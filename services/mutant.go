package services

import "errors"

func IsDnaMutant(dna []string) (bool, error) {
	n := len(dna)

	if n == 0 {
		return false, errors.New("dna is empty")
	}

	for _, row := range dna {
		if len(row) != n {
			return false, errors.New("dna is not square")
		}
	}

	return isMutant(dna), nil
}

func isMutant(dna []string) bool {

	n := len(dna)

	sequencesOfFourCount := 0

	// Si la matriz es menor a 4x4, no puede haber más de 1 secuencia de 4
	if n < 4 {
		return false
	}

	// Iterar por cada celda de la matriz
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {

			currentChar := dna[i][j]

			// Verificar horizontalmente
			if j+3 < n && currentChar == dna[i][j+1] && currentChar == dna[i][j+2] && currentChar == dna[i][j+3] {
				sequencesOfFourCount++
			}

			// Verificar verticalmente
			if i+3 < n && currentChar == dna[i+1][j] && currentChar == dna[i+2][j] && currentChar == dna[i+3][j] {
				sequencesOfFourCount++
			}

			// Verificar diagonal principal
			if i+3 < n && j+3 < n && currentChar == dna[i+1][j+1] && currentChar == dna[i+2][j+2] && currentChar == dna[i+3][j+3] {
				sequencesOfFourCount++
			}

			// Verificar diagonal secundaria
			if i+3 < n && j-3 >= 0 && currentChar == dna[i+1][j-1] && currentChar == dna[i+2][j-2] && currentChar == dna[i+3][j-3] {
				sequencesOfFourCount++
			}

			if sequencesOfFourCount > 1 {
				return true
			}
		}
	}

	return false
}
