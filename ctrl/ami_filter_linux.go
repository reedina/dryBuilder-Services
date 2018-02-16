package ctrl

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/reedina/dryBuilder_services/model"
)

//CreateAmiFilterLinux (POST)
func CreateAmiFilterLinux(w http.ResponseWriter, r *http.Request) {
	var filterLinux model.AmiFilterLinux

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&filterLinux); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	//Does the Instance Type Resource Exist ?
	if model.DoesAmiFilterLinuxResourceExist(&filterLinux) == true {
		respondWithError(w, http.StatusConflict, "Resource already exists")
		return
	}

	//Does the Packer Builder ID Exist ?
	if model.DoesPackerBuilderTypeIDExist(filterLinux.BuilderTypesID) == false {
		respondWithError(w, http.StatusBadRequest, "Packer Builder Types does not exist")
		return
	}

	//Resource does not exist, go ahead and create resource
	if err := model.CreateAmiFilterLinux(&filterLinux); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, filterLinux)
}

//GetAmiFilterLinuxes  (GET)
func GetAmiFilterLinuxes(w http.ResponseWriter, r *http.Request) {

	filterLinuxes, err := model.GetAmiFilterLinuxes()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, filterLinuxes)
}

//GetAmiFilterLinux (GET)
func GetAmiFilterLinux(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid AMI Filter Linux ID")
		return
	}

	filterLinux := model.AmiFilterLinux{ID: id}
	if err := model.GetAmiFilterLinux(&filterLinux); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "AMI Filter Linux not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, filterLinux)
}

//GetAmiFilterLinuxByName (GET)
func GetAmiFilterLinuxByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	filterLinux := model.AmiFilterLinux{}
	filterLinux.Name = name

	if err := model.GetAmiFilterLinuxByName(&filterLinux); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "AMI Filter Linux not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, filterLinux)
}

//UpdateAmiFilterLinux (PUT)
func UpdateAmiFilterLinux(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid AMI Filter Linux ID")
		return
	}

	var filterLinux model.AmiFilterLinux

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&filterLinux); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	filterLinux.ID = id

	// Does Instance Type Resource Exist ?
	if model.DoesAmiFilterLinuxIDExist(filterLinux.ID) != true {
		respondWithError(w, http.StatusBadRequest, "AMI Filter Linux ID does not exist")
		return
	}

	// Does Instance Type Name exists for another ID
	if model.DoesAmiFilterLinuxExistForAnotherID(filterLinux.Name, filterLinux.ID) == true {
		respondWithError(w, http.StatusConflict, "AMI Filter Linux Exists for another AMI Filter Linux ID")
		return
	}

	//Does the Packer Builder ID Exist ?
	if model.DoesPackerBuilderTypeIDExist(filterLinux.BuilderTypesID) == false {
		respondWithError(w, http.StatusBadRequest, "Packer Builder Types does not exist")
		return
	}

	if err := model.UpdateAmiFilterLinux(&filterLinux); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, filterLinux)
}

//DeleteAmiFilterLinux (DELETE)
func DeleteAmiFilterLinux(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid AMI Filter Linux ID")
		return
	}
	filterLinux := model.AmiFilterLinux{ID: id}

	if err := model.DeleteAmiFilterLinux(&filterLinux); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
