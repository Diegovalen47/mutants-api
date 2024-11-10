package services

import (
	"testing"

	"github.com/Diegovalen47/mutants/models"
)

func TestGetDnaStatsWhenEmpty(t *testing.T) {
	dnas := []models.Dna{}

	want := DnaStatsResponseDto{
		MutantsCount: 0,
		HumansCount:  0,
		Ratio:        0,
	}

	result := GetDnaStats(dnas)

	if want != result {
		t.Fatalf("want %v, got %v", want, result)
	}
}

func TestGetDnaStatsWhenHumans(t *testing.T) {
	dnas := []models.Dna{
		{IsMutant: false},
		{IsMutant: false},
	}

	want := DnaStatsResponseDto{
		MutantsCount: 0,
		HumansCount:  2,
		Ratio:        0,
	}

	result := GetDnaStats(dnas)

	if want != result {
		t.Fatalf("want %v, got %v", want, result)
	}
}

func TestGetDnaStatsWhenMutants(t *testing.T) {
	dnas := []models.Dna{
		{IsMutant: true},
		{IsMutant: true},
	}

	want := DnaStatsResponseDto{
		MutantsCount: 2,
		HumansCount:  0,
		Ratio:        1,
	}

	result := GetDnaStats(dnas)

	if want != result {
		t.Fatalf("want %v, got %v", want, result)
	}
}

func TestGetDnaStatsWhenEqualHumansAndMutants(t *testing.T) {
	dnas := []models.Dna{
		{IsMutant: false},
		{IsMutant: true},
	}

	want := DnaStatsResponseDto{
		MutantsCount: 1,
		HumansCount:  1,
		Ratio:        1,
	}

	result := GetDnaStats(dnas)

	if want != result {
		t.Fatalf("want %v, got %v", want, result)
	}
}

func TestGetDnaStatsWhenMultipleHumansAndMutants(t *testing.T) {
	dnas := []models.Dna{
		{IsMutant: false},
		{IsMutant: true},
		{IsMutant: false},
		{IsMutant: false},
		{IsMutant: false},
		{IsMutant: false},
		{IsMutant: false},
		{IsMutant: false},
	}

	want := DnaStatsResponseDto{
		MutantsCount: 1,
		HumansCount:  7,
		Ratio:        0.14285714285714285,
	}

	result := GetDnaStats(dnas)

	if want != result {
		t.Fatalf("want %v, got %v", want, result)
	}
}
