package model

import (
	"database/sql"
)

//Region  (TYPE)
type Region struct {
	ID         int    `json:"id"`
	RegionName string `json:"name"`
	Region     string `json:"region"`
	EndPoint   string `json:"endpoint"`
}

//Regions (TYPE)
type Regions struct {
	Regions []*Region `json:"regions"`
}

//DoesRegionResourceExist (POST)
func DoesRegionResourceExist(region *Region) bool {

	err := db.QueryRow("SELECT id, region FROM aws_regions WHERE region=?", region.Region).
		Scan(&region.ID, &region.Region)

	if err == sql.ErrNoRows {
		return false
	}

	return true
}

//DoesRegionIDExist (POST)
func DoesRegionIDExist(ID int) bool {

	var id int
	err := db.QueryRow("SELECT id FROM aws_regions WHERE id=?", ID).Scan(&id)

	if err == sql.ErrNoRows {
		return false
	}

	return true
}

//DoesRegionExistForAnotherID (PUT)
func DoesRegionExistForAnotherID(region string, id int) bool {

	var dbID int
	err := db.QueryRow("SELECT id FROM aws_regions WHERE region=?", region).Scan(&dbID)

	if err == sql.ErrNoRows {
		return false
	}

	if dbID != id {
		return true
	}

	return false
}

//CreateRegion (POST)
func CreateRegion(region *Region) error {

	res, err := db.Exec("INSERT INTO aws_regions(region, region_name, endpoint) VALUES(?,?,?)",
		region.Region, region.RegionName, region.EndPoint)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	region.ID = int(id)

	return nil
}

//GetRegions (GET)
func GetRegions() ([]Region, error) {
	rows, err := db.Query("SELECT id, region, region_name, endpoint FROM aws_regions")

	if err != nil {
		return nil, err
	}

	regions := []Region{}

	for rows.Next() {
		defer rows.Close()

		var r Region
		if err := rows.Scan(&r.ID, &r.Region, &r.RegionName, &r.EndPoint); err != nil {
			return nil, err
		}
		regions = append(regions, r)
	}

	return regions, nil
}

//GetRegion (GET)
func GetRegion(region *Region) error {
	return db.QueryRow("SELECT region,region_name, endpoint FROM aws_regions WHERE id=?", region.ID).
		Scan(&region.Region, &region.RegionName, &region.EndPoint)
}

//GetRegionByName (GET)
func GetRegionByName(region *Region) error {
	return db.QueryRow("SELECT id, region, region_name, endpoint from aws_regions where region=?",
		region.Region).Scan(&region.ID, &region.Region, &region.RegionName, &region.EndPoint)
}

//UpdateRegion (PUT)
func UpdateRegion(region *Region) error {
	_, err :=
		db.Exec("UPDATE aws_regions SET region=?, region_name=?, endpoint=? WHERE id=?",
			region.Region, region.RegionName, region.EndPoint, region.ID)

	return err
}

//DeleteRegion (DELETE)
func DeleteRegion(region *Region) error {
	_, err := db.Exec("DELETE FROM aws_regions WHERE id=?", region.ID)

	return err
}
