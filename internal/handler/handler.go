package handler

import (
	"diplom/internal"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Handlers interface {
	Register(router *mux.Router)
}

type handler struct {
	Service
}

func NewHandler(service *Service) Handlers {
	return &handler{
		*service,
	}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc("/", Handler)
}

func Handler(w http.ResponseWriter, r *http.Request) {

	resultJson, err := json.MarshalIndent(internal.ResultS, " ", " ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resultJson)
}
