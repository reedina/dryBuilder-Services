package model

import (
	"database/sql"
)

//Auth (TYPE)
type Auth struct {
	ID          int    `json:"id"`
	AccountName string `json:"account_name"`
	AccessType  string `json:"access_type"`
	AccessKeyID string `json:"access_key_id"`
	SecretKey   string `json:"secret_key"`
}

//DoesAuthResourceExist (POST)
func DoesAuthResourceExist(auth *Auth) bool {

	err := db.QueryRow("SELECT id, account_name FROM aws_auth WHERE account_name=?", auth.AccountName).
		Scan(&auth.ID, &auth.AccountName)

	if err == sql.ErrNoRows {
		return false
	}

	return true
}

//DoesAuthIDExist (POST)
func DoesAuthIDExist(ID int) bool {

	var id int
	err := db.QueryRow("SELECT id FROM aws_auth WHERE id=?", ID).Scan(&id)

	if err == sql.ErrNoRows {
		return false
	}

	return true
}

//DoesAuthExistForAnotherID (PUT)
func DoesAuthExistForAnotherID(accountName string, id int) bool {

	var dbID int
	err := db.QueryRow("SELECT id FROM aws_auth WHERE account_name=?", accountName).Scan(&dbID)

	if err == sql.ErrNoRows {
		return false
	}

	if dbID != id {
		return true
	}

	return false
}

//CreateAuth (POST)
func CreateAuth(auth *Auth) error {

	res, err := db.Exec("INSERT INTO aws_auth(account_name, access_type, access_key_id, secret_key) VALUES(?,?,?,?)",
		auth.AccountName, auth.AccessType, auth.AccessKeyID, auth.SecretKey)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	auth.ID = int(id)

	return nil
}

//GetAuths (GET)
func GetAuths() ([]Auth, error) {
	rows, err := db.Query("SELECT id, account_name, access_type, access_key_id, secret_key FROM aws_auth")

	if err != nil {
		return nil, err
	}

	auths := []Auth{}

	for rows.Next() {
		defer rows.Close()

		var r Auth
		if err := rows.Scan(&r.ID, &r.AccountName, &r.AccessType, &r.AccessKeyID, &r.SecretKey); err != nil {
			return nil, err
		}
		auths = append(auths, r)
	}

	return auths, nil
}

//GetAuth (GET)
func GetAuth(auth *Auth) error {
	return db.QueryRow("SELECT account_name, access_type, access_key_id, secret_key FROM aws_auth WHERE id=?", auth.ID).
		Scan(&auth.AccountName, &auth.AccessType, &auth.AccessKeyID, &auth.SecretKey)
}

//GetAuthByName (GET)
func GetAuthByName(auth *Auth) error {
	return db.QueryRow("SELECT id, account_name, access_type, access_key_id, secret_key from aws_auth where account_name=?",
		auth.AccountName).Scan(&auth.ID, &auth.AccountName, &auth.AccessType, &auth.AccessKeyID, &auth.SecretKey)
}

//UpdateAuth (PUT)
func UpdateAuth(auth *Auth) error {
	_, err :=
		db.Exec("UPDATE aws_auth SET account_name=?, access_type=?, access_key_id=?, secret_key=? WHERE id=?",
			auth.AccountName, auth.AccessType, auth.AccessKeyID, auth.SecretKey, auth.ID)

	return err
}

//DeleteAuth (DELETE)
func DeleteAuth(auth *Auth) error {
	_, err := db.Exec("DELETE FROM aws_auth WHERE id=?", auth.ID)

	return err
}
