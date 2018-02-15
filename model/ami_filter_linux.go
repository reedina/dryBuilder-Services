package model

import (
	"database/sql"
)

//AmiFilterLinux (TYPE)
type AmiFilterLinux struct {
	ID                 int    `json:"id"`
	BuilderTypesID     int    `json:"builder_types_id"`
	FriendlyName       string `json:"friendly_name"`
	Description        string `json:"description"`
	SSHUsername        string `json:"ssh_username"`
	VirtualizationType string `json:"virtualization_type"`
	Name               string `json:"name"`
	RootDeviceType     string `json:"root_device_type"`
	MostRecent         string `json:"most_recent"`
	Owners             string `json:"owners"`
}

//AmiFilterLinuxes (TYPE)
type AmiFilterLinuxes struct {
	AmiFilterLinuxes []*AmiFilterLinux `json:"ami_filter_linuxes"`
}

//DoesAmiFilterLinuxResourceExist (POST)
func DoesAmiFilterLinuxResourceExist(filterLinux *AmiFilterLinux) bool {

	err := db.QueryRow("SELECT id, friendly_name FROM ami_filter_linux WHERE friendly_name=?", filterLinux.FriendlyName).
		Scan(&filterLinux.ID, &filterLinux.FriendlyName)

	if err == sql.ErrNoRows {
		return false
	}

	return true
}

//DoesAmiFilterLinuxIDExist (POST)
func DoesAmiFilterLinuxIDExist(ID int) bool {

	var id int
	err := db.QueryRow("SELECT id FROM ami_filter_linux WHERE id=?", ID).Scan(&id)

	if err == sql.ErrNoRows {
		return false
	}

	return true
}

//DoesAmiFilterLinuxExistForAnotherID (PUT)
func DoesAmiFilterLinuxExistForAnotherID(friendlName string, id int) bool {

	var dbID int
	err := db.QueryRow("SELECT id FROM ami_filter_linux WHERE friendly_name=?", friendlName).Scan(&dbID)

	if err == sql.ErrNoRows {
		return false
	}

	if dbID != id {
		return true
	}

	return false
}

//CreateAmiFilterLinux (POST)
func CreateAmiFilterLinux(filterLinux *AmiFilterLinux) error {

	res, err := db.Exec("INSERT INTO ami_filter_linux(builder_types_id, friendly_name, description, ssh_username,"+
		"virtualization_type, name, root_device_type, most_recent, owners) VALUES(?,?,?,?,?,?,?,?,?)",
		filterLinux.BuilderTypesID, filterLinux.FriendlyName, filterLinux.Description, filterLinux.SSHUsername,
		filterLinux.VirtualizationType, filterLinux.Name, filterLinux.RootDeviceType, filterLinux.MostRecent, filterLinux.Owners)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	filterLinux.ID = int(id)

	return nil
}

//GetAmiFilterLinuxes (GET)
func GetAmiFilterLinuxes() ([]AmiFilterLinux, error) {
	rows, err := db.Query("SELECT id, builder_types_id, friendly_name, description, ssh_username, virtualization_type" +
		",name, root_device_type, most_recent, owners FROM ami_filter_linux")

	if err != nil {
		return nil, err
	}

	filterLinuxes := []AmiFilterLinux{}

	for rows.Next() {
		defer rows.Close()

		var r AmiFilterLinux
		if err := rows.Scan(&r.ID, &r.BuilderTypesID, &r.FriendlyName, &r.Description, &r.SSHUsername,
			&r.VirtualizationType, &r.Name, &r.RootDeviceType, &r.MostRecent, &r.Owners); err != nil {
			return nil, err
		}
		filterLinuxes = append(filterLinuxes, r)
	}

	return filterLinuxes, nil
}

//GetAmiFilterLinux (GET)
func GetAmiFilterLinux(filterLinux *AmiFilterLinux) error {
	return db.QueryRow("SELECT builder_types_id, friendly_name, description, ssh_username, virtualization_type"+
		",name, root_device_type, most_recent, owners FROM ami_filter_linux WHERE id=?", filterLinux.ID).
		Scan(&filterLinux.BuilderTypesID, &filterLinux.FriendlyName, &filterLinux.Description, &filterLinux.SSHUsername,
			&filterLinux.VirtualizationType, &filterLinux.Name, &filterLinux.RootDeviceType, &filterLinux.MostRecent,
			&filterLinux.Owners)

}

//GetAmiFilterLinuxByName (GET)
func GetAmiFilterLinuxByName(filterLinux *AmiFilterLinux) error {
	return db.QueryRow("SELECT id, builder_types_id, friendly_name, description, ssh_username, virtualization_type"+
		",name, root_device_type, most_recent, owners from ami_filter_linux where friendly_name=?",
		filterLinux.FriendlyName).Scan(&filterLinux.ID, &filterLinux.BuilderTypesID, &filterLinux.FriendlyName,
		&filterLinux.Description, &filterLinux.SSHUsername, &filterLinux.VirtualizationType, &filterLinux.Name, &filterLinux.RootDeviceType,
		&filterLinux.MostRecent, &filterLinux.Owners)
}

//UpdateAmiFilterLinux (PUT)
func UpdateAmiFilterLinux(filterLinux *AmiFilterLinux) error {

	_, err := db.Exec("UPDATE ami_filter_linux SET builder_types_id=?, friendly_name=?, description=?, "+
		"ssh_username=?, virtualization_type=?, name=?, root_device_type=?, most_recent=?, owners=? WHERE id=?",
		filterLinux.BuilderTypesID, filterLinux.FriendlyName, filterLinux.Description, filterLinux.SSHUsername,
		filterLinux.VirtualizationType, filterLinux.Name, filterLinux.RootDeviceType, filterLinux.MostRecent,
		filterLinux.Owners, filterLinux.ID)

	return err
}

//DeleteAmiFilterLinux (DELETE)
func DeleteAmiFilterLinux(filterLinux *AmiFilterLinux) error {
	_, err := db.Exec("DELETE FROM ami_filter_linux WHERE id=?", filterLinux.ID)

	return err
}
