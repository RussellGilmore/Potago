package controllers

import (
	"net/http"

	"github.com/RussellGilmore/potago/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Demo Home Controller")
}
