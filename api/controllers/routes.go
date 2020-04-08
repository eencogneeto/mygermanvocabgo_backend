package controllers

import "github.com/eencogneeto/backend/api/middlewares"

func (s *Server) initializeRoutes() {
	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Posts routes
	s.Router.HandleFunc("/nouns", middlewares.SetMiddlewareJSON(s.CreateNoun)).Methods("POST")
	s.Router.HandleFunc("/nouns", middlewares.SetMiddlewareJSON(s.GetNouns)).Methods("GET")
	s.Router.HandleFunc("/nouns/{id}", middlewares.SetMiddlewareJSON(s.GetNoun)).Methods("GET")
	s.Router.HandleFunc("/nouns/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateNoun))).Methods("PUT")
	s.Router.HandleFunc("/nouns/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteNoun)).Methods("DELETE")
}
