package services

import "github.com/Diegovalen47/mutants/models"

type DnaStatsResponseDto struct {
	MutantsCount int     `json:"count_mutant_dna"`
	HumansCount  int     `json:"count_human_dna"`
	Ratio        float64 `json:"ratio"`
}

func GetDnaStats(dnas []models.Dna) DnaStatsResponseDto {
	mutantsCount := 0
	humansCount := 0
	ratio := 0.0

	for _, dna := range dnas {
		if dna.IsMutant {
			mutantsCount++
		} else {
			humansCount++
		}
	}

	if humansCount == 0 && mutantsCount == 0 {
		ratio = 0
	} else if humansCount == 0 && mutantsCount != 0 {
		ratio = 1
	} else {
		ratio = float64(mutantsCount) / float64(humansCount)
	}

	return DnaStatsResponseDto{
		MutantsCount: mutantsCount,
		HumansCount:  humansCount,
		Ratio:        ratio,
	}
}
