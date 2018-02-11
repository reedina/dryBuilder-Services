package model

import (
	"database/sql"
)

//InstanceType (TYPE)
type InstanceType struct {
	ID       int    `json:"id"`
	Category string `json:"category"`
	Type     string `json:"type"`
	Vcpu     int    `json:"vcpu"`
	MemoryGb int    `json:"memory_gb"`
}

//InstanceTypes (TYPE)
type InstanceTypes struct {
	InstanceTypes []*InstanceType `json:"instance_types"`
}

//DoesInstanceTypeResourceExist (POST)
func DoesInstanceTypeResourceExist(instanceType *InstanceType) bool {

	err := db.QueryRow("SELECT id, type FROM aws_instance_types WHERE type=?", instanceType.Type).
		Scan(&instanceType.ID, &instanceType.Type)

	if err == sql.ErrNoRows {
		return false
	}

	return true
}

//DoesInstanceTypeIDExist (POST)
func DoesInstanceTypeIDExist(ID int) bool {

	var id int
	err := db.QueryRow("SELECT id FROM aws_instance_types WHERE id=?", ID).Scan(&id)

	if err == sql.ErrNoRows {
		return false
	}

	return true
}

//DoesInstanceTypeExistForAnotherID (PUT)
func DoesInstanceTypeExistForAnotherID(instanceType string, id int) bool {

	var dbID int
	err := db.QueryRow("SELECT id FROM aws_instance_types WHERE type=?", instanceType).Scan(&dbID)

	if err == sql.ErrNoRows {
		return false
	}

	if dbID != id {
		return true
	}

	return false
}

//CreateInstanceType (POST)
func CreateInstanceType(instanceType *InstanceType) error {

	res, err := db.Exec("INSERT INTO aws_instance_types(category, type, vcpu, memory_gb) VALUES(?,?,?,?)",
		instanceType.Category, instanceType.Type, instanceType.Vcpu, instanceType.MemoryGb)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	instanceType.ID = int(id)

	return nil
}

//GetInstanceTypes (GET)
func GetInstanceTypes() ([]InstanceType, error) {
	rows, err := db.Query("SELECT id, category, type, vcpu, memory_gb FROM aws_instance_types")

	if err != nil {
		return nil, err
	}

	instanceTypes := []InstanceType{}

	for rows.Next() {
		defer rows.Close()

		var r InstanceType
		if err := rows.Scan(&r.ID, &r.Category, &r.Type, &r.Vcpu, &r.MemoryGb); err != nil {
			return nil, err
		}
		instanceTypes = append(instanceTypes, r)
	}

	return instanceTypes, nil
}

//GetInstanceType (GET)
func GetInstanceType(instanceType *InstanceType) error {
	return db.QueryRow("SELECT category, type, vcpu, memory_gb FROM aws_instance_types WHERE id=?", instanceType.ID).
		Scan(&instanceType.Category, &instanceType.Type, &instanceType.Vcpu, &instanceType.MemoryGb)
}

//GetInstanceTypeByName (GET)
func GetInstanceTypeByName(instanceType *InstanceType) error {
	return db.QueryRow("SELECT id, category, type, vcpu, memory_gb from aws_instance_types where type=?",
		instanceType.Type).Scan(&instanceType.ID, &instanceType.Category, &instanceType.Type, &instanceType.Vcpu, &instanceType.MemoryGb)
}

//UpdateInstanceType (PUT)
func UpdateInstanceType(instanceType *InstanceType) error {
	_, err :=
		db.Exec("UPDATE aws_instance_types SET category=?, type=?, vcpu=?, memory_gb=? WHERE id=?",
			instanceType.Category, instanceType.Type, instanceType.Vcpu, instanceType.MemoryGb, instanceType.ID)

	return err
}

//DeleteInstanceType (DELETE)
func DeleteInstanceType(instanceType *InstanceType) error {
	_, err := db.Exec("DELETE FROM aws_instance_types WHERE id=?", instanceType.ID)

	return err
}
