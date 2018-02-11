package ctrl

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/reedina/dryBuilder_services/model"
)

//CreateInstanceType (POST)
func CreateInstanceType(w http.ResponseWriter, r *http.Request) {
	var instandeType model.InstanceType

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&instandeType); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	//Does the Instance Type Resource Exist ?
	if model.DoesInstanceTypeResourceExist(&instandeType) == true {
		respondWithError(w, http.StatusConflict, "Resource already exists")
		return
	}

	//Resource does not exist, go ahead and create resource
	if err := model.CreateInstanceType(&instandeType); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, instandeType)
}

//GetInstanceTypes  (GET)
func GetInstanceTypes(w http.ResponseWriter, r *http.Request) {

	instanceTypes, err := model.GetInstanceTypes()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, instanceTypes)
}

//GetInstanceType (GET)
func GetInstanceType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid AWS Instance Type ID")
		return
	}

	instanceType := model.InstanceType{ID: id}
	if err := model.GetInstanceType(&instanceType); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Instance Type not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, instanceType)
}

//GetInstanceTypeByName (GET)
func GetInstanceTypeByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	typeName := vars["name"]

	instanceType := model.InstanceType{}
	instanceType.Type = typeName

	if err := model.GetInstanceTypeByName(&instanceType); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Instance Type not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, instanceType)
}

//UpdateInstanceType (PUT)
func UpdateInstanceType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Instance Type ID")
		return
	}

	var instanceType model.InstanceType

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&instanceType); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	instanceType.ID = id

	// Does Instance Type Resource Exist ?
	if model.DoesInstanceTypeIDExist(instanceType.ID) != true {
		respondWithError(w, http.StatusBadRequest, "Instance Type ID does not exist")
		return
	}
	// Does Instance Type Name exists for another ID
	if model.DoesInstanceTypeExistForAnotherID(instanceType.Type, instanceType.ID) == true {
		respondWithError(w, http.StatusBadRequest, "Instance Type Exists for another Instance Type ID")
		return
	}
	if err := model.UpdateInstanceType(&instanceType); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, instanceType)
}

//DeleteInstanceType (DELETE)
func DeleteInstanceType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Instance Type ID")
		return
	}
	instanceType := model.InstanceType{ID: id}

	if err := model.DeleteInstanceType(&instanceType); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
