package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type ApiServer struct {
	svc Service
}

func NewApiServer(svc Service) *ApiServer {
	return &ApiServer{svc}
}

func (api *ApiServer) StartServer(listenAddr string) error {
	http.HandleFunc("/", api.handleGetCatFact)
	return http.ListenAndServe(listenAddr, nil)
}

func (api *ApiServer) handleGetCatFact(w http.ResponseWriter, r *http.Request) {
	resp, err := api.svc.GetCatFact(context.Background())
	if err != nil {
		JsonWriter(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		return
	}

	JsonWriter(w, http.StatusOK, resp)
}

func JsonWriter(w http.ResponseWriter, status int, value any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(value)
	if err != nil {
		return err
	}
	return nil
}
