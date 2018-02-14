package model

import (
	"database/sql"
)

//PackerBuilderType (TYPE)
type PackerBuilderType struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

//PackerBuilderTypes (TYPE)
type PackerBuilderTypes struct {
	PackerBuilderTypes []*PackerBuilderType `json:"packer_builder_types"`
}

//DoesPackerBuilderTypeResourceExist (POST)
func DoesPackerBuilderTypeResourceExist(builderType *PackerBuilderType) bool {

	err := db.QueryRow("SELECT id, type FROM packer_builder_types WHERE type=?", builderType.Type).
		Scan(&builderType.ID, &builderType.Type)

	if err == sql.ErrNoRows {
		return false
	}

	return true
}

//DoesPackerBuilderTypeIDExist (POST)
func DoesPackerBuilderTypeIDExist(ID int) bool {

	var id int
	err := db.QueryRow("SELECT id FROM packer_builder_types WHERE id=?", ID).Scan(&id)

	if err == sql.ErrNoRows {
		return false
	}

	return true
}

//DoesPackerBuilderTypeExistForAnotherID (PUT)
func DoesPackerBuilderTypeExistForAnotherID(builderType string, id int) bool {

	var dbID int
	err := db.QueryRow("SELECT id FROM packer_builder_types WHERE type=?", builderType).Scan(&dbID)

	if err == sql.ErrNoRows {
		return false
	}

	if dbID != id {
		return true
	}

	return false
}

//CreatePackerBuilderType (POST)
func CreatePackerBuilderType(builderType *PackerBuilderType) error {

	res, err := db.Exec("INSERT INTO packer_builder_types(name, description, type) VALUES(?,?,?)",
		builderType.Name, builderType.Description, builderType.Type)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	builderType.ID = int(id)

	return nil
}

//GetPackerBuilderTypes (GET)
func GetPackerBuilderTypes() ([]PackerBuilderType, error) {
	rows, err := db.Query("SELECT id, name, description, type FROM packer_builder_types")

	if err != nil {
		return nil, err
	}

	builderTypes := []PackerBuilderType{}

	for rows.Next() {
		defer rows.Close()

		var r PackerBuilderType
		if err := rows.Scan(&r.ID, &r.Name, &r.Description, &r.Type); err != nil {
			return nil, err
		}
		builderTypes = append(builderTypes, r)
	}

	return builderTypes, nil
}

//GetPackerBuilderType (GET)
func GetPackerBuilderType(builderType *PackerBuilderType) error {
	return db.QueryRow("SELECT name, description, type FROM packer_builder_types WHERE id=?", builderType.ID).
		Scan(&builderType.Name, &builderType.Description, &builderType.Type)
}

//GetPackerBuilderTypeByName (GET)
func GetPackerBuilderTypeByName(builderType *PackerBuilderType) error {
	return db.QueryRow("SELECT id, name, description, type from packer_builder_types where type=?",
		builderType.Type).Scan(&builderType.ID, &builderType.Name, &builderType.Description, &builderType.Type)
}

//UpdatePackerBuilderType (PUT)
func UpdatePackerBuilderType(builderType *PackerBuilderType) error {
	_, err :=
		db.Exec("UPDATE packer_builder_types SET name=?, description=?, type=? WHERE id=?",
			builderType.Name, builderType.Description, builderType.Type, builderType.ID)

	return err
}

//DeletePackerBuilderType (DELETE)
func DeletePackerBuilderType(builderType *PackerBuilderType) error {
	_, err := db.Exec("DELETE FROM packer_builder_types WHERE id=?", builderType.ID)

	return err
}
