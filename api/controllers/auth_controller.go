package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ZootHii/blog-go-backend/api/auth"
	"github.com/ZootHii/blog-go-backend/api/models"
	"github.com/ZootHii/blog-go-backend/api/responses"
	"github.com/ZootHii/blog-go-backend/api/utils/customerrors"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.CreateToken(user.Email, user.Password)
	if err != nil {
		customError := customerrors.CustomError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, customError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) CreateToken(email, password string) (string, error) {

	user := models.User{}

	err := server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.ID)
}
