package services

import (
	"testing"
)

func TestIsMutant(t *testing.T) {
	dnas := [][]string{
		{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"},
		{"ACCTGACTT", "AACCTTGGC", "AACCTTGGC", "AACTCACTT", "AACCTTGGT", "AACCTTGGT", "AACTTACTT", "AACCTTGGT", "AACCTTGGC"},
	}

	want := true

	for _, dna := range dnas {
		result, err := IsDnaMutant(dna)

		if want != result || err != nil {
			t.Fatalf("want %t, got %t, error %v", want, result, err)
		} else {
			t.Logf("want %t, got %t, error %v", want, result, err)
		}
	}
}

func TestIsNotMutant(t *testing.T) {
	dnas := [][]string{
		{"ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"},
		{"ACCTGACTT", "GTCCTTGGC", "AACCTTGGC", "GTGTAACTT", "CACCTAGGT", "TACCTTGGC", "AGCTGACTT", "AATCTTGGT", "AACCTTGGC"},
	}

	want := false

	for _, dna := range dnas {
		result, err := IsDnaMutant(dna)

		if want != result || err != nil {
			t.Fatalf("want %t, got %t, error %v", want, result, err)
		} else {
			t.Logf("want %t, got %t, error %v", want, result, err)
		}
	}
}
