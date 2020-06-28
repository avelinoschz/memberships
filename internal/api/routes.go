package api

func (s *Server) routes() {

	s.router.HandleFunc("/alive", s.HandleAlive()).Methods("GET")
}
