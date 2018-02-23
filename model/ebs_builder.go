package model

import (
	"database/sql"
)

// EbsBuilder (TYPE)
type EbsBuilder struct {
	ID                 int    `json:"id"`
	BuilderName        string `json:"builder_name"`
	AmiName            string `json:"ami_name"`
	AwsAuthID          int    `json:"aws_auth_id"`
	AwsRegionsID       int    `json:"aws_regions_id"`
	AwsInstanceTypesID int    `json:"aws_instance_types_id"`
	AwsAmiFilterID     int    `json:"aws_src_ami_filter_linux_id"`
	SSHUsername        string `json:"ssh_username"`
	AmiDescription     string `json:"ami_description"`
}

//DoesEbsBuilderResourceExist (POST)
func DoesEbsBuilderResourceExist(ebs *EbsBuilder) bool {

	err := db.QueryRow("SELECT id, builder_name FROM ebs_builders WHERE builder_name=?", ebs.BuilderName).
		Scan(&ebs.ID, &ebs.BuilderName)

	if err == sql.ErrNoRows {
		return false
	}

	return true
}

//DoesEbsBuilderIDExist (POST)
func DoesEbsBuilderIDExist(ID int) bool {

	var id int
	err := db.QueryRow("SELECT id FROM ebs_builders WHERE id=?", ID).Scan(&id)

	if err == sql.ErrNoRows {
		return false
	}

	return true
}

//DoesEbsBuilderExistForAnotherID (PUT)
func DoesEbsBuilderExistForAnotherID(builderName string, id int) bool {

	var dbID int
	err := db.QueryRow("SELECT id FROM ebs_builders WHERE builder_name=?", builderName).Scan(&dbID)

	if err == sql.ErrNoRows {
		return false
	}

	if dbID != id {
		return true
	}

	return false
}

//CreateEbsBuilder (POST)
func CreateEbsBuilder(ebs *EbsBuilder) error {

	res, err := db.Exec("INSERT INTO ebs_builders(builder_name, ami_name, aws_auth_id, aws_regions_id, aws_instance_types_id,"+
		"aws_src_ami_filter_linux_id, ssh_username, ami_description) VALUES(?,?,?,?,?,?,?,?)",
		ebs.BuilderName, ebs.AmiName, ebs.AwsAuthID, ebs.AwsRegionsID, ebs.AwsInstanceTypesID, ebs.AwsAmiFilterID, ebs.SSHUsername, ebs.AmiDescription)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	ebs.ID = int(id)

	return nil
}

//GetEbsBuilders (GET)
func GetEbsBuilders() ([]EbsBuilder, error) {
	rows, err := db.Query("SELECT id, builder_name, ami_name, aws_auth_id, aws_regions_id, aws_instance_types_id," +
		" aws_src_ami_filter_linux_id, ssh_username, ami_description FROM ebs_builders")

	if err != nil {
		return nil, err
	}

	ebsBuilders := []EbsBuilder{}

	for rows.Next() {
		defer rows.Close()

		var r EbsBuilder
		if err := rows.Scan(&r.ID, &r.BuilderName, &r.AmiName, &r.AwsAuthID, &r.AwsRegionsID, &r.AwsInstanceTypesID,
			&r.AwsAmiFilterID, &r.SSHUsername, &r.AmiDescription); err != nil {
			return nil, err
		}
		ebsBuilders = append(ebsBuilders, r)
	}

	return ebsBuilders, nil
}

//GetEbsBuilder (GET)
func GetEbsBuilder(ebs *EbsBuilder) error {
	return db.QueryRow("SELECT id, builder_name, ami_name, aws_auth_id, aws_regions_id, aws_instance_types_id, "+
		"aws_src_ami_filter_linux_id,ssh_username, ami_description FROM ebs_builders WHERE id=?", ebs.ID).
		Scan(&ebs.ID, &ebs.BuilderName, &ebs.AmiName, &ebs.AwsAuthID, &ebs.AwsRegionsID, &ebs.AwsInstanceTypesID, &ebs.AwsAmiFilterID,
			&ebs.SSHUsername, &ebs.AmiDescription)
}

//GetEbsBuilderByName (GET)
func GetEbsBuilderByName(ebs *EbsBuilder) error {
	return db.QueryRow("SELECT id, builder_name, ami_name, aws_auth_id, aws_regions_id, aws_instance_types_id, "+
		"aws_src_ami_filter_linux_id, ssh_username, ami_description from ebs_builders where account_name=?",
		ebs.BuilderName).Scan(&ebs.ID, &ebs.BuilderName, &ebs.AmiName, &ebs.AwsAuthID, &ebs.AwsRegionsID, &ebs.AwsInstanceTypesID,
		&ebs.AwsAmiFilterID, &ebs.SSHUsername, &ebs.AmiDescription)
}

//UpdateEbsBuilder (PUT)
func UpdateEbsBuilder(ebs *EbsBuilder) error {
	_, err :=
		db.Exec("UPDATE ebs_builders SET builder_name=?, ami_name=?, aws_auth_id=?, aws_regions_id=?,"+
			" aws_instance_types_id=?, aws_src_ami_filter_linux_id=?, ssh_username=?, ami_description=? WHERE id=?",
			ebs.BuilderName, ebs.AmiName, ebs.AwsAuthID, ebs.AwsRegionsID, ebs.AwsInstanceTypesID, ebs.AwsAmiFilterID,
			ebs.SSHUsername, ebs.AmiDescription, ebs.ID)

	return err
}

//DeleteEbsBuilder (DELETE)
func DeleteEbsBuilder(ebs *EbsBuilder) error {
	_, err := db.Exec("DELETE FROM ebs_builders WHERE id=?", ebs.ID)

	return err
}
