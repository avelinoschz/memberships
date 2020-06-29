package api

func (s *Server) routes() {

	r := s.router.PathPrefix("/memberships/api/v1").Subrouter()

	r.HandleFunc("/alive", s.HandleAlive()).Methods("GET")
	r.HandleFunc("/members", s.HandleMemberCreate()).Methods("POST")
}
