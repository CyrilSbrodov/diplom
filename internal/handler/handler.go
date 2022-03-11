package handler

import (
	"diplom/internal"
	"encoding/json"
	"errors"
	"fmt"
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

	resultJson, err := json.MarshalIndent(internal.Result, " ", " ")
	if err != nil {
		errors.New(fmt.Sprintf("не удалось перекодировать данные. ошибка: %v", err))
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(resultJson)
}
