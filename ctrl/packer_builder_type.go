package ctrl

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/reedina/dryBuilder_services/model"
)

//CreatePackerBuilderType (POST)
func CreatePackerBuilderType(w http.ResponseWriter, r *http.Request) {
	var builderType model.PackerBuilderType

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&builderType); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	//Does the Instance Type Resource Exist ?
	if model.DoesPackerBuilderTypeResourceExist(&builderType) == true {
		respondWithError(w, http.StatusConflict, "Resource already exists")
		return
	}

	//Resource does not exist, go ahead and create resource
	if err := model.CreatePackerBuilderType(&builderType); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, builderType)
}

//GetPackerBuilderTypes  (GET)
func GetPackerBuilderTypes(w http.ResponseWriter, r *http.Request) {

	instanceTypes, err := model.GetPackerBuilderTypes()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, instanceTypes)
}

//GetPackerBuilderType (GET)
func GetPackerBuilderType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Packer Builder Type ID")
		return
	}

	instanceType := model.PackerBuilderType{ID: id}
	if err := model.GetPackerBuilderType(&instanceType); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Packer Builder Type not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, instanceType)
}

//GetPackerBuilderTypeByName (GET)
func GetPackerBuilderTypeByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	typeName := vars["name"]

	instanceType := model.PackerBuilderType{}
	instanceType.Type = typeName

	if err := model.GetPackerBuilderTypeByName(&instanceType); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Packer Builder Type not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, instanceType)
}

//UpdatePackerBuilderType (PUT)
func UpdatePackerBuilderType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Packer Builder Type ID")
		return
	}

	var instanceType model.PackerBuilderType

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&instanceType); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	instanceType.ID = id

	// Does Instance Type Resource Exist ?
	if model.DoesPackerBuilderTypeIDExist(instanceType.ID) != true {
		respondWithError(w, http.StatusBadRequest, "Packer Builder Type ID does not exist")
		return
	}
	// Does Instance Type Name exists for another ID
	if model.DoesPackerBuilderTypeExistForAnotherID(instanceType.Type, instanceType.ID) == true {
		respondWithError(w, http.StatusBadRequest, "Packer Builder Type Exists for another Builder Type Type ID")
		return
	}
	if err := model.UpdatePackerBuilderType(&instanceType); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, instanceType)
}

//DeletePackerBuilderType (DELETE)
func DeletePackerBuilderType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Packer Builder Type ID")
		return
	}
	instanceType := model.PackerBuilderType{ID: id}

	if err := model.DeletePackerBuilderType(&instanceType); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
