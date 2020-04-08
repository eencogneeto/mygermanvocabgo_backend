package controllers

import (
	"net/http"

	"github.com/eencogneeto/backend/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This API")
}
