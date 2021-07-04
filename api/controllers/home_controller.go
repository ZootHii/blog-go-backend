package controllers

import (
	"net/http"

	"github.com/ZootHii/blog-go-backend/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "BlogApp GO API")

}
