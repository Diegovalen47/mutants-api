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

	// Iterar por cada celda de la matriz
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// Verificar horizontalmente
			if j+3 < n && dna[i][j] == dna[i][j+1] && dna[i][j+1] == dna[i][j+2] && dna[i][j+2] == dna[i][j+3] {
				sequencesOfFourCount++
			}

			// Verificar verticalmente
			if i+3 < n && dna[i][j] == dna[i+1][j] && dna[i+1][j] == dna[i+2][j] && dna[i+2][j] == dna[i+3][j] {
				sequencesOfFourCount++
			}

			// Verificar diagonal principal
			if i+3 < n && j+3 < n && dna[i][j] == dna[i+1][j+1] && dna[i+1][j+1] == dna[i+2][j+2] && dna[i+2][j+2] == dna[i+3][j+3] {
				sequencesOfFourCount++
			}

			// Verificar diagonal secundaria
			if i+3 < n && j-3 >= 0 && dna[i][j] == dna[i+1][j-1] && dna[i+1][j-1] == dna[i+2][j-2] && dna[i+2][j-2] == dna[i+3][j-3] {
				sequencesOfFourCount++
			}

			if sequencesOfFourCount > 1 {
				return true
			}
		}
	}

	return false
}
