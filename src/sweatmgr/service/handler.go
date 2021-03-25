package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	consts "github.com/eduardocfalcao/hands-on-go-and-mongodb/src/sweatmgr/consts"
	logger "github.com/eduardocfalcao/hands-on-go-and-mongodb/src/sweatmgr/logger"
	models "github.com/eduardocfalcao/hands-on-go-and-mongodb/src/sweatmgr/models"
	"github.com/gorilla/mux"
)

// PingResponse - PingResponse type
type PingResponse struct {
	Message string `json:"message"`
}

func withUserContext(req *http.Request) *http.Request {
	ctx := context.WithValue(req.Context(), consts.UserIDCtxKey, req.Header.Get("UserID"))
	return req.WithContext(ctx)
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	response := PingResponse{Message: "pong"}

	status := http.StatusOK
	respBytes, err := json.Marshal(response)
	if err != nil {
		status = http.StatusInternalServerError
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(respBytes)
}

func createSweatHandler(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	status := http.StatusOK
	req = withUserContext(req)

	var s models.Sweat

	err := decoder.Decode(&s)
	if err != nil {
		status = http.StatusInternalServerError
	} else {
		fmt.Println(s)
		err = s.Create(req.Context())
		if err != nil {
			status = http.StatusInternalServerError
		}
	}

	rw.Header().Add("Content-Type", "applciation/json")
	rw.WriteHeader(status)
}

func getSweatSamplesHandler(rw http.ResponseWriter, req *http.Request) {
	status := http.StatusOK
	sweats, err := models.ListAllSweat(req.Context())

	if err != nil {
		logger.Get().Errorf("Error fetching data", err)
		status = http.StatusInternalServerError
	}

	respBytes, err := json.Marshal(sweats)
	if err != nil {
		logger.Get().Errorf("Error marshaling data", err)
		status = http.StatusInternalServerError
	}

	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(respBytes)
}

func getSweatByIDHandler(rw http.ResponseWriter, req *http.Request) {
	status := http.StatusOK
	params := mux.Vars(req)
	id := params["id"]

	sweat, err := models.GetSweatByID(req.Context(), id)

	if err != nil {
		logger.Get().Errorf("Error fetching sweat by the ID: ", id, ". Error: ", err)
		status = http.StatusNotFound
	} else {
		respBytes, err := json.Marshal(sweat)
		if err != nil {
			logger.Get().Errorf("Error mashaling the sweat. Error: ", err)
		} else {
			rw.Header().Add("Content-Type", "application/json")
			rw.Write(respBytes)
		}
	}

	rw.WriteHeader(status)
}

func getSweatByUserIDHandler(rw http.ResponseWriter, req *http.Request) {

}
