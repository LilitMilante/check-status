package api

import (
	"check-status/service"
	"encoding/json"
	"net/http"
)

type Server struct {
	r  *http.ServeMux
	ss *service.StatusService
}

func NewServer(ss *service.StatusService) *Server {
	r := http.NewServeMux()

	s := Server{
		r:  r,
		ss: ss,
	}

	return &s
}

func (s *Server) Start(port string) error {
	s.r.HandleFunc("/status", s.CheckStatusHandler)

	return http.ListenAndServe(":"+port, s.r)
}

func (s *Server) CheckStatusHandler(w http.ResponseWriter, r *http.Request) {
	var addresses struct {
		Adr []string `json:"adr"`
	}

	err := json.NewDecoder(r.Body).Decode(&addresses)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	status := s.ss.CheckStatus(addresses.Adr)

	err = json.NewEncoder(w).Encode(status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
