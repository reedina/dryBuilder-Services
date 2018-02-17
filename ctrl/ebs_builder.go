package ctrl

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/reedina/dryBuilder_services/model"
)

//CreateEbsBuilder (POST)
func CreateEbsBuilder(w http.ResponseWriter, r *http.Request) {
	var ebs model.EbsBuilder

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ebs); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	//Does the Instance Type Resource Exist ?
	if model.DoesEbsBuilderResourceExist(&ebs) == true {
		respondWithError(w, http.StatusConflict, "Resource already exists")
		return
	}

	//Resource does not exist, go ahead and create resource
	if err := model.CreateEbsBuilder(&ebs); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, ebs)
}

//GetEbsBuilders  (GET)
func GetEbsBuilders(w http.ResponseWriter, r *http.Request) {

	ebsBuilders, err := model.GetEbsBuilders()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, ebsBuilders)
}

//GetEbsBuilder (GET)
func GetEbsBuilder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid AWS EbsBuilder ID")
		return
	}

	ebs := model.EbsBuilder{ID: id}
	if err := model.GetEbsBuilder(&ebs); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "EbsBuilder not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, ebs)
}

//GetEbsBuilderByName (GET)
func GetEbsBuilderByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ebsName := vars["name"]

	ebs := model.EbsBuilder{}
	ebs.BuilderName = ebsName

	if err := model.GetEbsBuilderByName(&ebs); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "EbsBuilder not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, ebs)
}

//UpdateEbsBuilder (PUT)
func UpdateEbsBuilder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid EbsBuilder ID")
		return
	}

	var ebs model.EbsBuilder

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ebs); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	ebs.ID = id

	// Does EbsBuilder Resource Exist ?
	if model.DoesEbsBuilderIDExist(ebs.ID) != true {
		respondWithError(w, http.StatusBadRequest, "EbsBuilder ID does not exist")
		return
	}
	// Does EbsBuilder Account Name exists for another ID
	if model.DoesEbsBuilderExistForAnotherID(ebs.BuilderName, ebs.ID) == true {
		respondWithError(w, http.StatusBadRequest, "EbsBuilder Builder Name Exists for another EbsBuilder ID")
		return
	}
	if err := model.UpdateEbsBuilder(&ebs); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, ebs)
}

//DeleteEbsBuilder (DELETE)
func DeleteEbsBuilder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid EbsBuilder ID")
		return
	}
	ebs := model.EbsBuilder{ID: id}

	if err := model.DeleteEbsBuilder(&ebs); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
