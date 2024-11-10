package routes

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/Diegovalen47/mutants/db"
)

func TestPostDnaHandlerIfIsMutant(t *testing.T) {

	db.DBConnection()

	req, err := http.NewRequest("POST", "localhost:8080/mutant",
		bytes.NewBuffer(
			[]byte(`{"dna": ["ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"]}`),
		),
	)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostDnaHandler)

	handler.ServeHTTP(rr, req)

	status := rr.Code
	want := http.StatusOK
	if status != want {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, want)
	}
}

func TestPostDnaHandlerIfIsHuman(t *testing.T) {
	req, err := http.NewRequest("POST", "localhost:8080/mutant",
		bytes.NewBuffer(
			[]byte(`{"dna": ["ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"]}`),
		),
	)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostDnaHandler)

	handler.ServeHTTP(rr, req)

	status := rr.Code
	want := http.StatusForbidden
	if status != want {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, want)
	}
}

func TestGetDnaStatsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080/stats", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetDnaStatsHandler)

	handler.ServeHTTP(rr, req)

	status := rr.Code
	want := http.StatusOK

	if status != want {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, want)
	}

	expectedPattern := `{"count_mutant_dna":\d+,"count_human_dna":\d+,"ratio":\d+(\.\d+)?}`
	matched, err := regexp.MatchString(expectedPattern, rr.Body.String())

	if err != nil {
		t.Fatalf("error matching response body: %v", err)
	}
	if !matched {
		t.Errorf("handler returned unexpected body format: got %v", rr.Body.String())
	}
}
