package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eduardocfalcao/hands-on-go-and-mongodb/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateSweat(t *testing.T) {
	var jsonStr = []byte(`{ "glucose": 1.12, "sodium": 0.98, "chloride": 0.003 }`)

	req, err := http.NewRequest("POST", "/entry", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createSweatHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}

func TestSweatSamples(t *testing.T) {
	req, err := http.NewRequest("GET", "/sweat_samples", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getSweatSamplesHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var sweats []models.Sweat

	err = json.Unmarshal([]byte(rr.Body.String()), &sweats)
	if err != nil {
		t.Fatal(err)
	}

	if len(sweats) < 1 {
		t.Errorf("No sweat samples. Expected at least 1")
	}

}

func TestGetSweatByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/sweat/6025c8968ca2ab64a682f041", nil)

	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{
		"id": "6025c8968ca2ab64a682f041",
	})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getSweatByIDHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var sweat models.Sweat

	err = json.Unmarshal([]byte(rr.Body.String()), &sweat)
	if err != nil {
		t.Fatal(err)
	}

	expectedID, err := primitive.ObjectIDFromHex("6025c8968ca2ab64a682f041")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expectedID, sweat.ID)
}
