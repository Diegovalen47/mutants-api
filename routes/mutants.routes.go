package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Diegovalen47/mutants/db"
	"github.com/Diegovalen47/mutants/models"
	"github.com/Diegovalen47/mutants/services"
)

type postDnaBodyDto struct {
	Dna []string
}

func PostDnaHandler(w http.ResponseWriter, r *http.Request) {

	var dnaBody postDnaBodyDto

	json.NewDecoder(r.Body).Decode(&dnaBody)

	sequence := strings.Join(dnaBody.Dna, "")
	size := len(dnaBody.Dna)

	// Validar si la secuencia ya existe en BD y obtener el registro
	var existingDna models.Dna
	db.DB.Where("sequence = ?", sequence).First(&existingDna)

	if existingDna.ID != 0 {
		if existingDna.IsMutant {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
		return
	}

	// Si no existe en BD, validar si es mutante
	isMutant, errDna := services.IsDnaMutant(dnaBody.Dna)

	if errDna != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errDna.Error())
		return
	}

	newDna := models.Dna{Sequence: sequence, Size: size, IsMutant: isMutant}

	// Crear el registro en BD
	createdDna := db.DB.Create(&newDna)

	err := createdDna.Error

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	if newDna.IsMutant {
		w.WriteHeader(http.StatusOK)
		return
	} else {
		w.WriteHeader(http.StatusForbidden)
		return
	}
}

func GetDnaStatsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get DNA Stats"))
}
