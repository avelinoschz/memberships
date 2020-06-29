package api

func (s *Server) routes() {

	r := s.router.PathPrefix("/memberships/api/v1").Subrouter()

	r.HandleFunc("/alive", s.handleAlive()).Methods("GET")
	r.HandleFunc("/login", s.handleAuthLogin()).Methods("POST")

	r.HandleFunc("/members", s.handleMembersCreate()).Methods("POST")
	r.HandleFunc("/members/{id}/payments", s.handlePaymentsGet()).Methods("GET")
	r.HandleFunc("/members/{id}/payments", s.handlePaymentsCreate()).Methods("POST")
}
