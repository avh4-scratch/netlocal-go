package netlocal

import (
	"fmt"
	"net/http"
)

type Server struct {
	ResponseBodies map[string]string
	ResponseCodes  map[string]int
}

func Start() *Server {
	return &Server{ResponseBodies: make(map[string]string), ResponseCodes: make(map[string]int)}
}

func (s *Server) StubGet(port int, path string, responseCode int, responseBody string) {
	s.ResponseBodies[path] = responseBody
	s.ResponseCodes[path] = responseCode
	var portString = fmt.Sprintf(":%d", port)
	go http.ListenAndServe(portString, s)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	responseBody := s.ResponseBodies[path]
	responseCode := s.ResponseCodes[path]
	w.WriteHeader(responseCode)
	fmt.Fprintf(w, responseBody)
}

func Clear() {}
