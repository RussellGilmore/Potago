package controllers

import "github.com/RussellGilmore/potago/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Potato routes
	s.Router.HandleFunc("/potato", middlewares.SetMiddlewareJSON(s.CreatePotato)).Methods("POST")
	s.Router.HandleFunc("/potato", middlewares.SetMiddlewareJSON(s.GetPotatoes)).Methods("GET")
	s.Router.HandleFunc("/potato/{id}", middlewares.SetMiddlewareJSON(s.GetPotato)).Methods("GET")
	s.Router.HandleFunc("/potato/{id}", middlewares.SetMiddlewareJSON(s.UpdatePotato)).Methods("PUT")
	s.Router.HandleFunc("/potato/{id}", middlewares.SetMiddlewareJSON(s.DeletePotato)).Methods("DELETE")
}
