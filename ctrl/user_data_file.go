package ctrl

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/reedina/dryBuilder_services/model"
)

//CreateUserDataFile (POST)
func CreateUserDataFile(w http.ResponseWriter, r *http.Request) {
	var userDataFile model.UserDataFile

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userDataFile); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	//Does the UserDataFile Resource Exist ?
	if model.DoesUserDataFileResourceExist(&userDataFile) == true {
		respondWithError(w, http.StatusConflict, "Resource already exists")
		return
	}

	//Resource does not exist, go ahead and create resource
	if err := model.CreateUserDataFile(&userDataFile); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, userDataFile)
}

//GetUserDataFiles  (GET)
func GetUserDataFiles(w http.ResponseWriter, r *http.Request) {

	userDataFiles, err := model.GetUserDataFiles()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, userDataFiles)
}

//GetUserDataFile (GET)
func GetUserDataFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid UserDataFile ID")
		return
	}

	userDataFile := model.UserDataFile{ID: id}
	if err := model.GetUserDataFile(&userDataFile); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "UserDataFile not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, userDataFile)
}

//GetUserDataFileByName (GET)
func GetUserDataFileByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userDataFileName := vars["name"]

	userDataFile := model.UserDataFile{}
	userDataFile.Name = userDataFileName

	if err := model.GetUserDataFileByName(&userDataFile); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "UserDataFile not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, userDataFile)
}

//UpdateUserDataFile (PUT)
func UpdateUserDataFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid UserDataFile ID")
		return
	}

	var userDataFile model.UserDataFile

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userDataFile); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	userDataFile.ID = id

	// Does UserDataFile Resource Exist ?
	if model.DoesUserDataFileIDExist(userDataFile.ID) != true {
		respondWithError(w, http.StatusBadRequest, "UserDataFile ID does not exist")
		return
	}
	// Does UserDataFile Name exists for another ID
	if model.DoesUserDataFileExistForAnotherID(userDataFile.Name, userDataFile.ID) == true {
		respondWithError(w, http.StatusBadRequest, "UserDataFile Exists for another UserDataFile ID")
		return
	}
	if err := model.UpdateUserDataFile(&userDataFile); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, userDataFile)
}

//DeleteUserDataFile (DELETE)
func DeleteUserDataFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid UserDataFile ID")
		return
	}
	userDataFile := model.UserDataFile{ID: id}

	if err := model.DeleteUserDataFile(&userDataFile); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
