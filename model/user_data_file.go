package model

import (
	"database/sql"
)

//UserDataFile  (TYPE)
type UserDataFile struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	SourceCode  string `json:"source_code"`
	Tags        string `json:"tags"`
}

//UserDataFiles (TYPE)
type UserDataFiles struct {
	UserDataFiles []*UserDataFile `json:"user_data_files"`
}

//DoesUserDataFileResourceExist (POST)
func DoesUserDataFileResourceExist(userDataFile *UserDataFile) bool {

	err := db.QueryRow("SELECT id, name FROM user_data_files WHERE name=?", userDataFile.Name).
		Scan(&userDataFile.ID, &userDataFile.Name)

	if err == sql.ErrNoRows {
		return false
	}

	return true
}

//DoesUserDataFileIDExist (POST)
func DoesUserDataFileIDExist(ID int) bool {

	var id int
	err := db.QueryRow("SELECT id FROM aws_userDataFiles WHERE id=?", ID).Scan(&id)

	if err == sql.ErrNoRows {
		return false
	}

	return true
}

//DoesUserDataFileExistForAnotherID (PUT)
func DoesUserDataFileExistForAnotherID(name string, id int) bool {

	var dbID int
	err := db.QueryRow("SELECT id FROM user_data_files WHERE name=?", name).Scan(&dbID)

	if err == sql.ErrNoRows {
		return false
	}

	if dbID != id {
		return true
	}

	return false
}

//CreateUserDataFile (POST)
func CreateUserDataFile(userDataFile *UserDataFile) error {

	res, err := db.Exec("INSERT INTO user_data_files(name, description, source_code, tags) VALUES(?,?,?,?)",
		userDataFile.Name, userDataFile.Description, userDataFile.SourceCode, userDataFile.Tags)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	userDataFile.ID = int(id)

	return nil
}

//GetUserDataFiles (GET)
func GetUserDataFiles() ([]UserDataFile, error) {
	rows, err := db.Query("SELECT id, name, description, source_code, tags FROM user_data_files")

	if err != nil {
		return nil, err
	}

	userDataFiles := []UserDataFile{}

	for rows.Next() {
		defer rows.Close()

		var r UserDataFile
		if err := rows.Scan(&r.ID, &r.Name, &r.Description, &r.SourceCode, &r.Tags); err != nil {
			return nil, err
		}
		userDataFiles = append(userDataFiles, r)
	}

	return userDataFiles, nil
}

//GetUserDataFile (GET)
func GetUserDataFile(userDataFile *UserDataFile) error {
	return db.QueryRow("SELECT name, description, source_code, tags FROM user_data_files WHERE id=?", userDataFile.ID).
		Scan(&userDataFile.Name, &userDataFile.Description, &userDataFile.SourceCode, &userDataFile.Tags)
}

//GetUserDataFileByName (GET)
func GetUserDataFileByName(userDataFile *UserDataFile) error {
	return db.QueryRow("SELECT id, name, description, source_code, tags from user_data_files where name=?",
		userDataFile.Name).Scan(&userDataFile.ID, &userDataFile.Name, &userDataFile.Description, &userDataFile.SourceCode, &userDataFile.Tags)
}

//UpdateUserDataFile (PUT)
func UpdateUserDataFile(userDataFile *UserDataFile) error {
	_, err :=
		db.Exec("UPDATE user_data_files SET name=?, description=?, source_code=?, tags=? WHERE id=?",
			userDataFile.Name, userDataFile.Description, userDataFile.SourceCode, userDataFile.Tags, userDataFile.ID)

	return err
}

//DeleteUserDataFile (DELETE)
func DeleteUserDataFile(userDataFile *UserDataFile) error {
	_, err := db.Exec("DELETE FROM user_data_files WHERE id=?", userDataFile.ID)

	return err
}
