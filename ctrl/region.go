package ctrl

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/reedina/dryBuilder_services/model"
)

//CreateRegion (POST)
func CreateRegion(w http.ResponseWriter, r *http.Request) {
	var region model.Region

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&region); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	//Does the Region Resource Exist ?
	if model.DoesRegionResourceExist(&region) == true {
		respondWithError(w, http.StatusConflict, "Resource already exists")
		return
	}

	//Resource does not exist, go ahead and create resource
	if err := model.CreateRegion(&region); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, region)
}

//GetRegions  (GET)
func GetRegions(w http.ResponseWriter, r *http.Request) {

	regions, err := model.GetRegions()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, regions)
}

//GetRegion (GET)
func GetRegion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid AWS Region ID")
		return
	}

	region := model.Region{ID: id}
	if err := model.GetRegion(&region); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Region not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, region)
}

//GetRegionByName (GET)
func GetRegionByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	regionName := vars["name"]

	region := model.Region{}
	region.Region = regionName

	if err := model.GetRegionByName(&region); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Region not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, region)
}

//UpdateRegion (PUT)
func UpdateRegion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Region ID")
		return
	}

	var region model.Region

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&region); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	region.ID = id

	// Does Region Resource Exist ?
	if model.DoesRegionIDExist(region.ID) != true {
		respondWithError(w, http.StatusBadRequest, "Region ID does not exist")
		return
	}
	// Does Region Name exists for another ID
	if model.DoesRegionExistForAnotherID(region.Region, region.ID) == true {
		respondWithError(w, http.StatusBadRequest, "Region Exists for another Region ID")
		return
	}
	if err := model.UpdateRegion(&region); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, region)
}

//DeleteRegion (DELETE)
func DeleteRegion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Region ID")
		return
	}
	region := model.Region{ID: id}

	if err := model.DeleteRegion(&region); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
