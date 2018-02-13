package ctrl

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/reedina/dryBuilder_services/model"
)

//CreateAuth (POST)
func CreateAuth(w http.ResponseWriter, r *http.Request) {
	var auth model.Auth

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&auth); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	//Does the Instance Type Resource Exist ?
	if model.DoesAuthResourceExist(&auth) == true {
		respondWithError(w, http.StatusConflict, "Resource already exists")
		return
	}

	//Resource does not exist, go ahead and create resource
	if err := model.CreateAuth(&auth); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, auth)
}

//GetAuths  (GET)
func GetAuths(w http.ResponseWriter, r *http.Request) {

	auths, err := model.GetAuths()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, auths)
}

//GetAuth (GET)
func GetAuth(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid AWS Auth ID")
		return
	}

	auth := model.Auth{ID: id}
	if err := model.GetAuth(&auth); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Auth not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, auth)
}

//GetAuthByName (GET)
func GetAuthByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	authName := vars["name"]

	auth := model.Auth{}
	auth.AccountName = authName

	if err := model.GetAuthByName(&auth); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Auth not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, auth)
}

//UpdateAuth (PUT)
func UpdateAuth(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Auth ID")
		return
	}

	var auth model.Auth

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&auth); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	auth.ID = id

	// Does Auth Resource Exist ?
	if model.DoesAuthIDExist(auth.ID) != true {
		respondWithError(w, http.StatusBadRequest, "Auth ID does not exist")
		return
	}
	// Does Auth Account Name exists for another ID
	if model.DoesAuthExistForAnotherID(auth.AccountName, auth.ID) == true {
		respondWithError(w, http.StatusBadRequest, "Auth Account Name Exists for another Auth ID")
		return
	}
	if err := model.UpdateAuth(&auth); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, auth)
}

//DeleteAuth (DELETE)
func DeleteAuth(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Auth ID")
		return
	}
	auth := model.Auth{ID: id}

	if err := model.DeleteAuth(&auth); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
