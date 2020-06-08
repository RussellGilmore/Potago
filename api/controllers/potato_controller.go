package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/RussellGilmore/potago/api/models"
	"github.com/RussellGilmore/potago/api/responses"
	"github.com/RussellGilmore/potago/api/utils/formaterror"
	"github.com/gorilla/mux"
)

func (server *Server) CreatePotato(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	potato := models.Potato{}
	err = json.Unmarshal(body, &potato)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	potato.Prepare()
	potatoCreated, err := potato.SavePotato(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, potatoCreated.ID))
	responses.JSON(w, http.StatusCreated, potatoCreated)
}

func (server *Server) GetPotatoes(w http.ResponseWriter, r *http.Request) {

	potato := models.Potato{}

	potatoes, err := potato.FindAllPotatoes(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, potatoes)
}

func (server *Server) GetPotato(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	potato := models.Potato{}

	potatoReceived, err := potato.FindPotatoByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, potatoReceived)
}

func (server *Server) UpdatePotato(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	potato := models.Potato{}
	err = server.DB.Debug().Model(models.Potato{}).Where("id = ?", pid).Take(&potato).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Potato not found"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	potatoUpdate := models.Potato{}
	err = json.Unmarshal(body, &potatoUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	potatoUpdate.Prepare()

	potatoUpdate.ID = potato.ID

	potatoUpdated, err := potatoUpdate.UpdateAPotato(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, potatoUpdated)
}

func (server *Server) DeletePotato(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Check if the potato exist
	potato := models.Potato{}
	err = server.DB.Debug().Model(models.Potato{}).Where("id = ?", pid).Take(&potato).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	_, err = potato.DeleteAPotato(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusNoContent, "")
}
